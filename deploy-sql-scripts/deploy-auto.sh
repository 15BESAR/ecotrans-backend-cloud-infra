export SQL_NAME=dev-instance
export REGION=asia-southeast2
export GCE_ZONE=asia-southeast2-a
export ROOT_PASSWORD=15besarterbaik

gcloud sql instances create $SQL_NAME --database-version=MYSQL_5_7 \
    --cpu=1 --memory=3.75GB \
    --zone=$GCE_ZONE --root-password=$ROOT_PASSWORD
