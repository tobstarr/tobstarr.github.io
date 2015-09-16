# IAM region specific policies

	{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Action": [
					"ec2:DescribeInstances"
				],
			"Condition": {
				"StringEquals": {
					"ec2:Region":"eu-central-1"
				}
			},
				"Effect": "Allow",
				"Resource": "*"
			}
		]
	}
