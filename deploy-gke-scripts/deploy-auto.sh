# Author : David Fauzi 
# Desc : Script to deploy go app in GKE
# Tested on WSL2 with Cloud SDK with project owner roles
# To deploy, change variable if needed, then run script
# bash deploy-gke-scripts/deploy.sh

# Cluster is placed in asia-southeast2-a with currently 2 nodes
export USER_EMAIL=davidfauzi29@gmail.com
export GCP_PROJECT=test-capstone-350108
export COMPUTE_ZONE=asia-southeast2-a
export CLUSTER_NAME=capstone-cluster
export NUM_NODES=2
export DEPLOYMENT_NAME=go-deployment
export REPLICAS=3
export IMAGE_NAME=go-test
export TAG=v0.04 #Update it 
export IS_BUILD_IMAGE=true #true if wanna build container
export IS_APPLY_GKE=true #true if wanna build container

echo $IS_BUILD_IMAGE
# Build Container and submit to container registry
if [ "$IS_BUILD_IMAGE" == "true" ]; then
echo "Building container image and pushing to container registry..."
gcloud builds submit --tag gcr.io/$GCP_PROJECT/$IMAGE_NAME:$TAG
echo -e "Done pushing to Google Container Registry...\n\n"
fi
# check if current cluster node is not the same as num_node
export CURRENT_NODE=$(gcloud container clusters describe $CLUSTER_NAME --zone $COMPUTE_ZONE | grep initialNodeCount:| sed 's/[^0-9]*//g')
echo 'Current node is '$CURRENT_NODE
if [ "$CURRENT_NODE" -ne "$NUM_NODES" ]; then
echo "updating cluster node from $CURRENT_NODE to $NUM_NODES..." 
gcloud container clusters resize $CLUSTER_NAME --num-nodes=$NUM_NODES --zone $COMPUTE_ZONE
fi

# Update deployment and service
if [ "$IS_APPLY_GKE" == "true" ]; then
echo "applying kubernetes \n\n"
python3 -u "./deploy-gke-scripts/edit_yaml.py" 'deploy-gke-scripts/deployment.yaml' $DEPLOYMENT_NAME $REPLICAS \
$GCP_PROJECT $IMAGE_NAME $TAG
kubectl apply -f deploy-gke-scripts/deployment.yaml
kubectl apply -f deploy-gke-scripts/service.yaml
fi
echo "Deploy Done..."
echo "Check API Endpoint using"
echo "kubectl get services"