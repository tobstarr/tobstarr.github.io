# AWS

## Upload Server Certificate

	aws iam upload-server-certificate --server-certificate-name example.com --certificate-body=file://example.com.crt --private-key file://example.com.key

## List Ubuntu Images

canonical owner id: 099720109477

	aws ec2 describe-images --owners 099720109477 --filters 'Name=name,Values=ubuntu/images/hvm/ubuntu-trusty-14.04-amd64-*' | jq -c -r '.Images[] | {name: .Name, id: .ImageId }' | sort -r | less


## List own Images

	aws ec2 describe-images --owner=self --region eu-central-1 | jq -r -c '.Images[] | { id: .ImageId, state: .State, Name: .Name }'

## Documentation

	http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Operations.html
