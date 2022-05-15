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
#### :white_check_mark: SUCCESS 
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "userId": 132312312,
    "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```
#### :red_circle: FAILED
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
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "userId": 132312312,
    "token": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "username/password is incorrect"
}
```
## GET All Profile Data
### Request
`GET /profiles`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "users":[
        {
            "firstName": "fname",
            "lastName": "lname",
            "ageOfBirth": "10-03-2000",
            "age": 21,
            "sex": "m/f",
            "address": "lorem ipsum...",
            "occupation": "student",
            "point": 10202,
            "totalRedeem" : 100,
            "totalDistance" : 2023.5,
            "totalEmissionReduced" : 500,
            "badge" : 0
        },
        ...
    ]
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed"
}
```

## GET User Profile Data
### Request
`GET /profile/<userId>`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "firstName": "fname",
    "lastName": "lname",
    "ageOfBirth": "10-03-2000",
    "age": 21,
    "sex": "m/f",
    "address": "lorem ipsum...",
    "occupation": "student",
    "point": 10202,
    "totalRedeem" : 100,
    "totalDistance" : 2023.5,
    "totalEmissionReduced" : 500,
    "badge" : 0
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed/user not found"
}
```

## Update User Profile Data
### Request
`PUT /profile/<userId>`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`
```json
{
    "firstName": "fname",
    "lastName": "lname",
    "ageOfBirth": "10-03-2000",
    "age": 21,
    ...
}
```
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "firstName": "fname",
    "lastName": "lname",
    "ageOfBirth": "10-03-2000",
    "age": 21,
    ...
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed/user not found/update data not valid"
}
```

## AutoComplete Gmaps API
### Request
`GET /autocomplete?input=jalan+su`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "predictions" : [
        {"text" : "Jalan sudirman"},
        {"text" : "Jalan sutomo"},
        {"text" : "Jalan sukajadi"},
    ] // from highest confidence level to lowest
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed"
}
```

## GET All Alternative Route
### Request
`GET /routes`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`
```json
{
    "origin" : {
        "latitude": -6.8837833,
        "longitude": 107.6135736
    },
    "destination" : {
        "latitude": -6.9200488,
        "longitude": 107.6237797
    },
    "preferences" : {
        "walk" : true,
        "bicycle" : false,
        "bus" : true,
        "train" : true
    }
}
```
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "routes":[
        {
        // most likely the same as gmaps direction api documentation response,
        // but i'll add carbonEmissionSaved attributes for each route
        },
        ...
    ]
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed/place doesn't exist/route doesnt exist"
}
```

## Finish Journey
### Request
`POST /finish`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`
```json
{
    "origin": "ChIJl02Bz3GMaS4RCgefgFZdKtI", // gmaps place_id
    "destination": "ChIJY9TrwiH0aS4RrvGqlZvI_Mw",
    "startTime": "2018-12-10T13:49:51.141Z",
    "finishTime": "2018-12-10T16:49:51.141Z",
    "distanceTravelled" : 10.43, // in km
    "carbonSaved" : 4.45, // in g co2
    "rewards" : 102 // in point 
}
```
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "msg": "data saved successfully"
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed/data not valid"
}
```

## Get All Vouchers
### Request
`GET /vouchers`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
   "vouchers":[
       {
           "voucherId":"c2njbn4",
           "name":"Free 3 donut",
           "category":"food and beverages",
           "description": "free 3 donut berlaku untuk seluruh cabang Jco",
           "image": "https://storage.googleapis.com/voucher-images-2909/jco.jpg",
           "partner": "Jco",
           "price" : 1000,
           "expire date": "2018-12-10T16:49:51.141Z"
       },
       ... // listed from most to least recommended
   ]
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed"
}
```

## Buy Voucher
### Request
`POST /voucher`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`
```json
{
    "voucherId": "c2njbn4"
}
```
### Response
#### :white_check_mark: SUCCESS
`HTTP 200 OK`
`Content-type: application/json`
```json
{
    "msg": "Purchase Successful !",
    "pointRemaining":2000
}
```
#### :red_circle: FAILED
`HTTP 400 Bad Request`
`Content-type: application/json`
```json
{
    "msg": "Authorization failed/Stock empty"
}
```

## Other Optional API 
### Get Air Quality Forecast
request air quality prediction for destination location 
### Change Password
### Forgot Password
### Add Voucher 
### Update Voucher
### Delete Voucher
### API for Admin Dashboard Website

