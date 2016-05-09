# Google Cloud

## Default Zone

	europe-west1-b

## List Machine Types

	gcloud compute machine-types list --zone europe-west1-b

## Create a machine

	gcloud compute instances create build --zone europe-west1-b --image ubuntu-15-10 --machine-type g1-small --scopes https://www.googleapis.com/auth/devstorage.read_only


## Create Kubernetes Cluster

	gcloud container clusters create <CLUSTER_NAME> --enable-cloud-logging --enable-cloud-monitoring -m g1-small --wait

## Get k8s login credentials

	gcloud container clusters get-credentials <CLUSTER_NAME>

## Get an access token

	gcloud auth print-access-token
