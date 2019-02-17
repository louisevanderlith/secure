# secure
Mango API: Secure

Secure handles user registration and login

## Run with Docker
*$ go build
*$ docker build -t avosa/secure:dev .
*$ docker rm secureDEV
*$ docker run -d --network host --name secureDEV avosa/secure:dev 
*$ docker logs secureDEV