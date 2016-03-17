# Google Cloud

## List Machine Types

	gcloud compute machine-types list

## Create Kubernetes Cluster

	gcloud container clusters create default --enable-cloud-logging --enable-cloud-monitoring -m g1-small --wait

## Get an access token

	gcloud auth print-access-token
