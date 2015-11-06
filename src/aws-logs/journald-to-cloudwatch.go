package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/dynport/dgtk/cli"
	"github.com/dynport/wunderscale/awscfg"
	"github.com/dynport/wunderscale/images"
	"github.com/dynport/wunderscale/instances"
	"github.com/dynport/wunderscale/ping"
	"github.com/dynport/wunderscale/wait"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	router := cli.NewRouter()
	router.RegisterFunc("logs", indexLogs, "Index Logs")
	router.RegisterFunc("run-instances", runInstances, "Run Instances")
	switch err := router.RunWithArgs(); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		l.Fatal(err)
	}
}

func runInstances() error {
	l := log.New(os.Stderr, "", 0)
	_ = l
	c, err := cfg()
	if err != nil {
		return err
	}
	list := []*ec2.Image{}
	// also check out {"name":"ubuntu-15.04-Snappy-core-stable-15.04-20151023.1408","id":"ami-bd6d53ca"}
	ubuntu := []string{"14.04", "15.10"}
	for _, u := range ubuntu {
		i, err := images.Ubuntu(c, u)
		if err != nil {
			return err
		}
		list = append(list, i)
	}
	i, err := images.Coreos(c)
	if err != nil {
		return err
	}
	list = append(list, i)
	wg := &sync.WaitGroup{}
	for _, i := range list {
		wg.Add(1)
		go func(i *ec2.Image) {
			defer wg.Done()
			if err := checkImage(i, c); err != nil {
				l.Printf("error=%s", err)
			}
		}(i)
	}
	wg.Wait()
	return nil
}

func checkImage(image *ec2.Image, cfg *aws.Config) error {
	l := log.New(os.Stderr, "image="+p2s(image.Name)+" ", 0)
	i := &ec2.RunInstancesInput{
		ImageId:      image.ImageId,
		KeyName:      aws.String("tschwab"),
		MaxCount:     aws.Int64(1),
		MinCount:     aws.Int64(1),
		InstanceType: aws.String("t2.small"),
		Placement:    &ec2.Placement{AvailabilityZone: aws.String("eu-central-1a")},
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{
				SubnetId:                 aws.String("subnet-8d39d8e4"),        // not sure what we
				Groups:                   []*string{aws.String("sg-23fd604a")}, // actually need here
				DeviceIndex:              aws.Int64(0),
				AssociatePublicIpAddress: aws.Bool(true),
			},
		},
	}
	l.Printf("state=starting")
	res, err := ec2.New(session.New(cfg)).RunInstances(i)
	if err != nil {
		return err
	}
	if v, ex := len(res.Instances), 1; ex != v {
		return fmt.Errorf("expected %d instance, found %d", ex, v)
	} else if res.Instances[0].InstanceId == nil {
		return fmt.Errorf("instance id is nil")
	}
	instanceID := *res.Instances[0].InstanceId

	start := time.Now()
	instance, err := instances.WaitForState(cfg, instanceID, "running", 5*time.Second, 10*time.Minute)
	if err != nil {
		return err
	}
	l.Printf("state=running total_time=%.06f", time.Since(start).Seconds())

	err = wait.For(1*time.Second, 10*time.Minute, func() (bool, error) {
		return ping.SSH(*instance.PublicIpAddress, 100*time.Millisecond), nil
	})
	if err != nil {
		return err
	}
	l.Printf("state=ssh total_time=%.06f", time.Since(start).Seconds())
	_, err = ec2.New(session.New(cfg)).TerminateInstances(&ec2.TerminateInstancesInput{InstanceIds: []*string{aws.String(instanceID)}})
	return err
}

func p2s(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func cfg() (*aws.Config, error) {
	c, err := awscfg.NewFromPath(os.ExpandEnv("$HOME/.phrase/config.json"))
	if err != nil {
		return nil, err
	}
	return c.WithRegion("eu-central-1"), nil
}

func indexLogs() error {
	l := log.New(os.Stderr, "", 0)
	c, err := cfg()
	if err != nil {
		return err
	}
	cl := cloudwatchlogs.New(session.New(c))
	out := json.NewEncoder(os.Stdout)
	if res, err := cl.DescribeLogGroups(nil); err != nil {
		return err
	} else {
		for _, r := range res.LogGroups {
			out.Encode(r)
		}
	}
	name := "default"
	var sequenceNumber *string
	if res, err := cl.DescribeLogStreams(&cloudwatchlogs.DescribeLogStreamsInput{LogGroupName: aws.String(name)}); err != nil {
		return err
	} else {
		for _, r := range res.LogStreams {
			if r.LogStreamName != nil && *r.LogStreamName == "x220" {
				sequenceNumber = r.UploadSequenceToken
			}
		}
		out.Encode(res)
	}
	_ = sequenceNumber

	for {
		events := []*cloudwatchlogs.InputLogEvent{}
		for i := 0; i < 500; i++ {
			now := time.Now()
			events = append(events, &cloudwatchlogs.InputLogEvent{
				Message: aws.String(fmt.Sprintf("message %s", now.Format(time.RFC3339Nano))), Timestamp: aws.Int64(timeToMillis(now)),
			})
		}

		start := time.Now()
		if res, err := cl.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{SequenceToken: sequenceNumber, LogEvents: events, LogGroupName: aws.String("default"), LogStreamName: aws.String("x220")}); err != nil {
			return err
		} else {
			l.Printf("put metrics in %.06f", time.Since(start).Seconds())
			sequenceNumber = res.NextSequenceToken
		}
		fmt.Print("waiting: ")
		if line, _, err := bufio.NewReader(os.Stdin).ReadLine(); err == nil {
			l.Printf("read: %s", line)
		}
	}
}

func timeToMillis(in time.Time) int64 {
	return int64(float64(in.UnixNano()) / float64(time.Millisecond))
}
