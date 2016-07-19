# JQ

tutorial at: https://stedolan.github.io/jq/tutorial/

	cat issues.json | jq '.[].title' -r

	-r: raw output
	-c: compact output

# Iterate and return json

	jq '.DBInstances[] | {id: .DBInstanceIdentifier, status: .DBInstanceStatus}'

# List AWS instances

	aws ec2 describe-instances | jq -c -r '.Reservations[] | .Instances[] | { name: .Tags[] | select(.Key == "Name").Value, id: .InstanceId, launched: .LaunchTime,  }' | sort


# TSV Output

	kubectl get events -o json | jq '.items[] | "\(.lastTimestamp)\t\(.reason)\t\(.message)"' -c -r


# Keys of a hash

	echo '{"a":1,"b":2}' | jq 'keys[] | .' -c -r

# jq for html

	go get -v github.com/ericchiang/pup
