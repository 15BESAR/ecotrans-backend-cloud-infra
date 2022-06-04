# Introduction

ecoTrans is a mobile application that allows users to travel to reach their destination with the most environmentally friendly method. It provides user with list of transportation route alternative along with the carbon emission that's reduced compared to private transportation. It also gives user air quality prediction and rewards user with points that can be exchanged for various voucher. The app also gives user voucher recommendation based on user preferences, voucher availibility, and purchase pattern.
</br>

## Endpoints

Currently hosted on GKE with endpoints:
`34.101.186.107`
</br>
Try accessing API with http://34.101.186.107

## Backend and Cloud Tech Stack

Backend created with Go with gin framework, containerization with docker and cloud build, ingress load balancing with nginx, and deployed on Google Kubernetes Engine (GKE)

### Notes (For Android Dev)

For current implementation, there's API that's already implementated but not included in the docs because it's intended for debugging/testing only. You also don't have to use all the API listed (eg: you only use read and delete but we'll provide whole CRUD API anyway).

### Notes (For ML Dev)

Current implementation doesn't utilize ML since the models haven't been deployed yet, please confirm the deployment strategy ASAP to ease the integration process between ML backend and endpoint in the future

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
bash deploy-gke-scripts/deploy-auto.sh
```

# API Documentation

The Rest API is described below..

## Root and Versioning

### 1. Root

#### Request

`GET /`
`Accept: application/json`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "msg": "Ecotrans GO Backend API"
}
```

### 2. Version

#### Request

`GET /version`
`Accept: application/json`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "lastUpdate": "May 21",
  "version": "0.50"
}
```

## User Authentication

### 3. Register Account

#### Request

`POST /register`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "username": "davidfauzi",
  "password": "kekekesiu",
  "email": "davidfauzi@gmail.com",
  "firstName": "david",
  "lastName": "fauzi",
  "birthDate": "2001-01-01T00:00:00Z"
}
```

#### :white_check_mark: Success Response

`HTTP 201 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "status": "Account has been created"
}
```

### 4. User Login

#### Request

`POST /login`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "username": "davidfauzi",
  "password": "kekekesiu"
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "success",
  "loginResult": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJFY29UcmFucyIsImV4cCI6MTY1Mzc1MTExNSwidXNlcm5hbWUiOiJkYXZpZGZhdXppIn0.neR6lbT79PG1rs98feVwvhWftU_YcfDWSpzkh3bGJYw",
    "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db"
  }
}
```

### 5. Refresh Token

### Request

`Post /refresh`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6Ik..."
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "success",
  "token": "eyJhbGciOiJIUzI1NiIsx2az5cCI6Ik..."
}
```

## Users

### 6. Get All Users

#### Request

`GET /users`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "success",
  "users": [
    {
      "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
      "email": "davidfauzi@gmail.com",
      "username": "davidfauzi",
      "password": "$2a$10$Iy63y2dLid.vVad6XMZcseRpXB.3J4Jn5FftMkE92fOdQngk6k9gm",
      "firstName": "david",
      "lastName": "fauzi",
      "birthDate": "2001-01-01T00:00:00Z",
      "age": 21,
      "gender": "",
      "job": "Student",
      "points": 40000,
      "voucherInterest": "foodAndBeverage,tranportation,phoneCredit",
      "domicile": "Bandung",
      "education": "Bachelor",
      "marriageStatus": false,
      "income": 2500000,
      "vehicle": "car",
      "Journeys": null
    }
  ]
}
```

### 7. Get User By ID

#### Request

`GET /user/:userId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "user": {
    "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
    "email": "davidfauzi@gmail.com",
    "username": "davidfauzi",
    "password": "$2a$10$Iy63y2dLid.vVad6XMZcseRpXB.3J4Jn5FftMkE92fOdQngk6k9gm",
    "firstName": "david",
    "lastName": "fauzi",
    "birthDate": "2001-01-01T00:00:00Z",
    "age": 21,
    "gender": "",
    "job": "Student",
    "points": 40000,
    "voucherInterest": "foodAndBeverage,tranportation,phoneCredit",
    "domicile": "Bandung",
    "education": "Bachelor",
    "marriageStatus": false,
    "income": 2500000,
    "vehicle": "car",
    "Journeys": null
  }
}
```

### 8. Update User Profile By ID

#### Request

`PUT /user/:userId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "job": "Student",
  "voucherInterest": "foodAndBeverage,tranportation,phoneCredit",
  "domicile": "Bandung",
  "education": "Bachelor",
  "marriageStatus": false,
  "income": 2500000,
  "vehicle": "car"
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
    "error": false,
		"msg":   "User get",
        "user" : {
    "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
    "email": "davidfauzi@gmail.com",
    "username": "davidfauzi",
    "password": "$2a$10$Iy63y2dLid.vVad6XMZcseRpXB.3J4Jn5FftMkE92fOdQngk6k9gm",
    .....
        }
}
```

### 9. Delete User Profile By ID

#### Request

`DELETE /user/:userId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "user deleted"
}
```

## Maps and Journey

### 10. AutoComplete Gmaps API

#### Request

`POST /autocomplete`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "input": "jalan su"
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "predictions": [
    // The same as Gmaps Autocomplete API, read the documentation
  ]
}
```

### 11. Get All Alternative Route

#### Request

`GET /routes`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "origin": "Jalan Tubagus Depan No 76",
  "destination": "Borma Dago",
  "preference": "walking"
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "geocode_waypoints": [],
  "routes": [] // The same as Gmaps Direction API, with addition of carbon in each route object
}
```

### 12. Get All Journeys

#### Request

`GET /journeys`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "journeys": [
    {
      "journeyId": "0a881751-3bb0-4ec0-9652-82241a8ba5f6",
      "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
      "startTime": "2018-12-10T13:49:51Z",
      "finishTime": "2018-12-10T16:49:51Z",
      "origin": "ChIJl02Bz3GMaS4RCgefgFZdKtI",
      "destination": "ChIJY9TrwiH0aS4RrvGqlZvI_Mw",
      "distanceTravelled": 10.43,
      "emissionSaved": 4.45,
      "reward": 102
    }
  ]
}
```

### 13. Finish Journey

#### Request

`POST /journey`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "origin": "ChIJl02Bz3GMaS4RCgefgFZdKtI", // gmaps place_id
  "destination": "ChIJY9TrwiH0aS4RrvGqlZvI_Mw",
  "startTime": "2018-12-10T13:49:51.141Z",
  "finishTime": "2018-12-10T16:49:51.141Z",
  "distanceTravelled": 10.43, // in km
  "carbonSaved": 4.45, // in g co2
  "rewards": 102 // in point
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "journeyId": "c06404e8-6d87-46a1-9049-13ec12d887ce",
  "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
  "startTime": "2018-12-10T13:49:51.141Z",
  "finishTime": "2018-12-10T16:49:51.141Z",
  "origin": "ChIJl02Bz3GMaS4RCgefgFZdKtI",
  "destination": "ChIJY9TrwiH0aS4RrvGqlZvI_Mw",
  "distanceTravelled": 10.43,
  "emissionSaved": 4.45,
  "reward": 102
}
```

### 14. Get Journey By ID

#### Request

`GET /journey/:journeyId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "journey": {
    "journeyId": "c06404e8-6d87-46a1-9049-13ec12d887ce",
    "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
    "startTime": "2018-12-10T13:49:51.141Z",
    ...
  }
}
```

## Voucher

### 15. Get All Vouchers

#### Request

`GET /vouchers?company=tokopedia`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
There's optional query string company=<partner_name> to filter the data

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "vouchers": [
    {
      "voucherId": "f0e21e1d-6ce6-450b-a115-e7c50c006d3b",
      "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
      "partnerName": "Tokopedia",
      "voucherName": "Free Ongkir 10 Ribu",
      "voucherDesc": "Gratis Ongkir sebesar 10 ribu untuk pembelian barang melalui aplikasi tokopedia",
      "category": "ecommerce",
      "imageUrl": "https://storage.googleapis.com/voucher-images-2909/jco.jpg",
      "stock": 0,
      "price": 1000
    }
  ]
}
```

### 16. Get Voucher By ID

#### Request

`GET /voucher/:voucherId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "voucher": {
    "voucherId": "f0e21e1d-6ce6-450b-a115-e7c50c006d3b",
    "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
    "partnerName": "Tokopedia",
    "voucherName": "Free Ongkir 10 Ribu",
    ...
  }
}
```

### 17. Add Voucher

#### Request

`POST /voucher`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "partnerID": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
  "partnerName": "Tokopedia",
  "voucherName": "Free Ongkir 15 Ribu",
  "voucherDesc": "Gratis Ongkir sebesar 10 ribu untuk pembelian barang melalui aplikasi tokopedia",
  "category": "ecommerce",
  "imageUrl": "https://storage.googleapis.com/voucher-images-2909/jco.jpg",
  "stock": 10,
  "price": 1000
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "voucher": {
    "voucherId": "c07489a6-05c0-4511-a34e-33297d406bd2",
    "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
    "partnerName": "Tokopedia",
    "voucherName": "Free Ongkir 15 Ribu",
    ...
  }
}
```

### 18. Update Voucher

#### Request

`PUT /voucher/:voucherId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
    "partnerID": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
    "partnerName": "Tokopedia",
    "voucherName":"Free Ongkir 15 Ribu",
    ... // Same parameter
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "voucher": {
    "voucherId": "c07489a6-05c0-4511-a34e-33297d406bd2",
    "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217"
    ... //Same object
  }
}
```

### 19. Delete Voucher

#### Request

`DELETE /voucher/:voucherId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "Purchase deleted"
}
```

## Purchase Voucher API

### 20. Purchase Voucher

#### Request

`POST /purchase`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
`Content-type: application/json`

```json
{
  "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
  "voucherId": "f0e21e1d-6ce6-450b-a115-e7c50c006d3b",
  "buyDate": "2018-12-10T13:49:51.141Z",
  "buyQuantity": 1
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "receipt": {
    "purchaseId": "678dd157-e5fe-4b0c-b36d-e0bbe4b2b6f7",
    "voucherId": "f0e21e1d-6ce6-450b-a115-e7c50c006d3b",
    "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
    "buyDate": "2018-12-10T13:49:51.141Z",
    "buyQuantity": 1,
    "voucherStockRemaining": 14,
    "userPointsRemaining": 39000
  }
}
```

### 21. Get All Purchases History

#### Request

`GET /purchases?user=d06d8777-8..`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`
There's optional query string user=<user_id> to filter the data

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "get purchases",
  "purchases": [
    {
      "purchaseId": "678dd157-e5fe-4b0c-b36d-e0bbe4b2b6f7",
      "voucherId": "f0e21e1d-6ce6-450b-a115-e7c50c006d3b",
      "userId": "d06d8777-896e-4a74-8f81-7b530b17f9db",
      "buyDate": "2018-12-10T13:49:51Z",
      "buyQuantity": 1
    }
  ]
}
```

## Partner Authentication API (Intended for Future Ongoing Admin Website)

### 22. Register Account

#### Request

`POST /company/register`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "username": "fauzi",
  "password": "3123121312",
  "email": "fauzi@gmail.com",
  "partnerName": "Tokopedia"
}
```

#### :white_check_mark: Success Response

`HTTP 201 OK`
`Content-type: application/json`

```json
{
  "status": "Account has been created"
}
```

### 23. User Login

#### Request

`POST /company/login`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "username": "fauzi",
  "password": "3123121312"
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "Succsess login",
  "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXV.."
}
```

### 24. Refresh Token

### Request

`Post /company/refresh`
`Accept: application/json`
`Content-Type: application/json`

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6Ik..."
}
```

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ.."
}
```

## Partners

### 25. Get All Users

#### Request

`GET /users`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "partnerts found",
  "partners": [
    {
      "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
      "partnerName": "Tokopedia",
      "email": "fauzi@gmail.com",
      "username": "fauzi",
      "password": "$2a$10$ASL2wvRHY8fIr3v8x2D/WOuLTcb2Nf5hVppZSz0EquAUsu1gJD48C",
      "Vouchers": null
    }
  ]
}
```

### 26. Get User By ID

#### Request

`GET /user/:partnerId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "partner": {
    "partnerId": "7afc5909-2411-4f9b-8c65-1abc40ce9217",
    "partnerName": "Tokopedia",
    "email": "fauzi@gmail.com",
    "username": "fauzi",
    "password": "$2a$10$ASL2wvRHY8fIr3v8x2D/WOuLTcb2Nf5hVppZSz0EquAUsu1gJD48C",
    "Vouchers": null
  }
}
```

### 27. Delete User Profile By ID

#### Request

`DELETE /user/:partnerId`
`Accept: application/json`
`Authorization: Bearer eyJhbGciOiJIUzI1NiIsIn...`

#### :white_check_mark: Success Response

`HTTP 200 OK`
`Content-type: application/json`

```json
{
  "error": false,
  "msg": "user deleted"
}
```

## Future Ongoing API

### ML-Related API to Connect ML Endpoint and Android Such as :

- GET Air Quality, Temperate, and UV prediction for a location
- GET Voucher recommendation for current user

### Change Password

### Forgot Password

## Future Development

Will try to create Website including landing page, partner login/register, and
dashboard to add/update/delete voucher. For current plan, tech stacks used will be NextJs, Typescript, Redux, Tailwind and will be deployed on Google App Engine
