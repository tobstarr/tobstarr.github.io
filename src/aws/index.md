# AWS

## API Documentation Links

* [EC2](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Operations.html)
* [ELB](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_Operations.html)
* [RDS](http://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Operations.html)

## Upload Server Certificate

	aws iam upload-server-certificate --server-certificate-name example.com --certificate-body=file://example.com.crt --private-key file://example.com.key

## List Ubuntu Images

canonical owner id: 099720109477

	aws ec2 describe-images --owners 099720109477 --filters 'Name=name,Values=ubuntu/images/hvm/ubuntu-*-14.04-amd64-*' | jq -c -r '.Images[] | {name: .Name, id: .ImageId }' | sort -r | less


## List own Images

	aws ec2 describe-images --owner=self --region eu-central-1 | jq -r -c '.Images[] | { id: .ImageId, state: .State, Name: .Name }'

## Documentation

	http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Operations.html

## List Instances

	aws ec2 describe-instances | jq -c -r '.Reservations[] | .Instances[] | { id: .InstanceId, name: .Tags[0].Value, state: .State.Name, ip: .PublicIpAddress }'

## Create Tags

	aws ec2 create-tags --resources i-b280590e --tags Key=Name,Value=gm

## S3 Storage Classes

* STANDARD
* STANDARD_IA

## EC2 Metadata Address

	169.254.169.254

## Security Credentials

  curl http://169.254.169.254/latest/meta-data/iam/security-credentials/

## Instance ID

	curl http://169.254.169.254/latest/meta-data/instance-id

## Upload asset with public read

	aws s3 cp --acl public-read

## List RHEL amis

	aws ec2 describe-images --owner 309956199498 | jq -c -r '.Images[] | { name: .Name }' | sort

## Run instances

	aws-mfa ec2 run-instances --user-data $(echo $ud | base64) --image-id ami-e25e6cff --key-name tschwab --associate-public-ip-address --instance-type t2.small
