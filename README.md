# Introduction
ecoTrans is a mobile application that allows users to travel to reach their destination with the most environmentally friendly method. It provides user with list of transportation route alternative along with the carbon emission that's reduced compared to private transportation. It also gives user air quality prediction and rewards user with points that can be exchanged for various voucher. The app also gives user voucher recommendation based on user preferences, voucher availibility, and purchase pattern.
</br>

## Backend and Cloud Tech Stack
Backend created with Go with gin framework, containerization with docker and cloud build, ingress load balancing with nginx, and deployed on Google Kubernetes Engine (GKE)

## Run Locally
### Using Go Run
dev
```bash
go mod download
go run main.go
```
build and run for windows
```bash
go build -o main.exe
main.exe
```
build and run for linux
```bash
go build -o main
main
```
### Using docker
```bash
docker build --tag test-go:v0.0 .
docker run -d -p 80:8080 test-go:v0.0
```
## Deploy to GKE 
The script will build the container using Cloud Build, scale cluster if needed, update kubernetes deployment config file, and apply changes
</br>
You need sufficient permission to be able to execute the script successfully
```bash
bash deploy-gke-scripts/deploy.sh 
```

# API Documentation
The Rest API is described below..
## Register User
### Request
`POST /register`
`Accept: application/json`
`Content-Type: application/json`
```json
{
    "username": "foo",
    "password": "123",
    "firstName": "fname",
    "lastName": "lname",
    "ageOfBirth": "10-03-2000",
    "age": 21,
    "sex": "m/f",
    "address": "lorem ipsum...",
    "occupation": "student",
}
```
### Response
#### ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+) SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "userId": 132312312,
    "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```
#### ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "username taken/ password not valid/ personal data not valid"
}
```
## User Login
### Request
`POST /login`
`Accept: application/json`
`Content-Type: application/json`
```json
{
    "username": "foo",
    "password": "123"
}
```
### Response
#### ![#c5f015](https://via.placeholder.com/15/c5f015/000000?text=+) SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "userId": 132312312,
    "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```
#### ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "username/password is incorrect"
}
```