package awsgo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var cfg = &ec2.RunInstancesInput{
	IamInstanceProfile: &ec2.IamInstanceProfileSpecification{Name: aws.String("profileName")},
	InstanceType:       aws.String("t2.medium"),
	MaxCount:           aws.Int64(1),
	MinCount:           aws.Int64(1),
	ImageId:            aws.String("image-id"),
	KeyName:            aws.String("tschwab"),
	Placement:          &ec2.Placement{AvailabilityZone: aws.String("eu-west-1")},
	NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
		{
			SubnetId:                 aws.String("subnet-1"),           // not sure what we
			Groups:                   []*string{aws.String("sg-1234")}, // actually need here
			DeviceIndex:              aws.Int64(0),
			AssociatePublicIpAddress: aws.Bool(true),
		},
	},
	BlockDeviceMappings: []*ec2.BlockDeviceMapping{
		{
			DeviceName: aws.String("/dev/sda1"),
			Ebs: &ec2.EbsBlockDevice{
				VolumeSize: aws.Int64(50),
				VolumeType: aws.String("gp2"),
			},
		},
	},
}
