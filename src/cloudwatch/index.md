# Cloudwatch

## Pricing

* you can do 1000 requests for 0.01 USD
* that is 100.000 requests per 1 USD
* that is 1.000.000 requests per 10 USD

* 1 mio requests means 1388 requests/hour, 23 requests/minute

For 10 USD you can do 1388 Requests/hour

## AWS/DynamoDB
	ProvisionedReadCapacityUnits
	ProvisionedWriteCapacityUnits

## AWS/EBS
	VolumeIdleTime
	VolumeQueueLength
	VolumeReadBytes
	VolumeReadOps
	VolumeTotalReadTime
	VolumeTotalWriteTime
	VolumeWriteBytes
	VolumeWriteOps

## AWS/EC2
	CPUCreditBalance
	CPUCreditUsage
	CPUUtilization
	DiskReadBytes
	DiskReadOps
	DiskWriteBytes
	DiskWriteOps
	NetworkIn
	NetworkOut
	StatusCheckFailed
	StatusCheckFailed_Instance
	StatusCheckFailed_System

### AWS/ElastiCache
	HashBasedCmds
	SetBasedCmds
	BytesUsedForCache
	CacheHits
	CacheMisses
	CPUUtilization
	CurrConnections
	CurrItems
	Evictions
	FreeableMemory
	GetTypeCmds
	KeyBasedCmds
	ListBasedCmds
	NetworkBytesIn
	NetworkBytesOut
	NewConnections
	Reclaimed
	ReplicationLag
	SetTypeCmds
	SortedSetBasedCmds
	StringBasedCmds
	SwapUsage

## AWS/ELB
	BackendConnectionErrors
	HTTPCode_Backend_2XX
	HTTPCode_Backend_3XX
	HTTPCode_Backend_4XX
	HTTPCode_Backend_5XX
	HTTPCode_ELB_5XX
	HealthyHostCount
	Latency
	RequestCount
	SurgeQueueLength
	UnHealthyHostCount

## AWS/RDS
	BinLogDiskUsage
	CPUUtilization
	DatabaseConnections
	DiskQueueDepth
	FreeStorageSpace
	FreeableMemory
	NetworkReceiveThroughput
	NetworkTransmitThroughput
	ReadIOPS
	ReadLatency
	ReadThroughput
	SwapUsage
	WriteIOPS
	WriteLatency
	WriteThroughput

## AWS/S3
	BucketSizeBytes
	NumberOfObjects

## AWS/SNS
	NumberOfMessagesPublished
	NumberOfNotificationsDelivered
	NumberOfNotificationsFailed
	PublishSize

## AWS/SQS
	ApproximateNumberOfMessagesDelayed
	ApproximateNumberOfMessagesNotVisible
	ApproximateNumberOfMessagesVisible
	NumberOfEmptyReceives
	NumberOfMessagesDeleted
	NumberOfMessagesReceived
	NumberOfMessagesSent
	SentMessageSize
