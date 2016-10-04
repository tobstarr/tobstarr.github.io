# Prometheus

## cpu of a container with a given name

	rate(container_cpu_user_seconds_total{kubernetes_container_name='rabbitmq'}[1m])

## select deployment

	container_memory_rss{kubernetes_pod_name=~"^phraseapp-worker.*",container_label_io_kubernetes_container_name!="POD",id=~"^/init.*"}


## container record
	{
		id="/system.slice/docker-4d7043794e5eb33f80e382872b6a962a2c7140073c1fa80786ebfddbdfa1b534.scope",
		image="rabbitmq:3.6.1-management",
		instance="192.168.178.58:4194",
		io_kubernetes_container_hash="625eb666",
		io_kubernetes_container_name="rabbitmq",
		io_kubernetes_container_restartCount="0",
		io_kubernetes_container_terminationMessagePath="/dev/termination-log",
		io_kubernetes_pod_name="rabbitmq-3347683304-begc5",
		io_kubernetes_pod_namespace="default",
		io_kubernetes_pod_terminationGracePeriod="30",
		io_kubernetes_pod_uid="6c021bda-5c06-11e6-a2e5-c03fd56491fd",
		job="cadvisor",
		kubernetes_container_name="rabbitmq",
		kubernetes_namespace="default",
		kubernetes_pod_name="rabbitmq-3347683304-begc5",
		name="k8s_rabbitmq.625eb666_rabbitmq-3347683304-begc5_default_6c021bda-5c06-11e6-a2e5-c03fd56491fd_f170ed83"
	}

## Kubernetes

annotate exporters with `prometheus.io/scrape: "true"`
