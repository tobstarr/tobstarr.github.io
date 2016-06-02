# README

## URLs

	https://storage.googleapis.com/kubernetes-release/release/v1.2.0-beta.0/bin/linux/amd64/kubectl

## Secure API Communication

	https://172.31.6.3:6443

## Query with labels

	ks get pods -l "k8s-app=heapster" -o json

## Access service via API

	https://<master-ip>/api/v1/proxy/namespaces/default/services/jenkins/

## IPs for a service

	service=nginx
	curl -s -H "Authorization: Bearer $(cat /etc/secrets/token)" -k https://${KUBERNETES_SERVICE_HOST}/api/v1/namespaces/default/endpoints/${service} | jq '.subsets[] | .addresses[] | .ip' -c -r | xargs | tr ' ' ','

## Add Context

	kubectl config set-cluster local --server http://127.0.0.1:808
	kubectl config set-context local --cluster local
