#!/bin/bash
# Author : David Fauzi
# Desc   : Instruction to deploy go apps to GKE 

# If starting shell, make sure to set project and default zone (opt)
gcloud config set project <PROJECT_ID>
gcloud config set compute/zone <ZONE>
# list of zones, location, and available machine types:
# https://cloud.google.com/compute/docs/regions-zones 

# Flow :
# Create Dockerfile -> Build Image -> Create Cluster -> Create kube config -> Apply 


# 1. Clone repo to shell VM
# 1a. Cloning from google source repo example:
gcloud auth list
gsutil cat gs://cloud-training/gsp318/marking/setup_marking_v2.sh | bash
gcloud source repos clone valkyrie-app
cd <folder

# 1b. Cloning from github
git clone <github-repo-link>
cd <folder>

# 2. Dockerfile 
# Make sure Dockerfile for Go is already set up 
# docs: https://docs.docker.com/language/golang/build-images/
# cd to main.go folder
touch Dockerfile
cat > Dockerfile << EOF
# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
EOF


# 3. Build Docker Image
# make sure to cd to the same root as Dockerfile
# 3a. Build with Docker
docker build -t <IMAGE>:<TAG> .
# test run docker image 
docker run -p 8080:8080 valkyrie-prod:v0.0.3 &
# push to container registry 
# https://cloud.google.com/container-registry/docs/pushing-and-pulling
docker tag valkyrie-prod:v0.0.3 gcr.io/$GOOGLE_CLOUD_PROJECT/valkyrie-prod:v0.0.3
docker push gcr.io/$GOOGLE_CLOUD_PROJECT/valkyrie-prod:v0.0.3

# 3b. Build with Google Build -> Push to Cloud registry
gcloud services enable cloudbuild.googleapis.com
gcloud services enable container.googleapis.com
# cd to folder 
gcloud builds submit --tag gcr.io/${GOOGLE_CLOUD_PROJECT}/monolith:1.0.0 .

# 4. Create GKE Cluster 
gcloud container clusters create fancy-cluster --num-nodes 3
# if done, check all node 
gcloud compute instances list

# 4a. Create deployment using kubectl create deployment (not recommended, but it's fast)
kubectl create deployment monolith --image=gcr.io/${GOOGLE_CLOUD_PROJECT}/monolith:1.0.0
# Expose port 
kubectl expose deployment monolith --type=LoadBalancer --port 80 --target-port 8080
# Update deployment to current image 
kubectl set image deployment/monolith monolith=gcr.io/${GOOGLE_CLOUD_PROJECT}/monolith:2.0.0

# 4b. Create kube config 
gcloud container clusters get-credentials <cluster_name> --zone <ZONE>
kubectl create -f k8s/deployment.yaml

# to edit config
kubectl edit deployment valkyrie-dev

# To redeploy, rebuild image, push to container, then edit again 

