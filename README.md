# secure
Mango API: Secure

Secure handles user registration and login

## Run with Docker
* $ go build
* $ gulp
* $ docker build -t avosa/secure:latest .
* $ docker rm SecureDEV
* $ docker run -d -e RUNMODE=DEV -p 8086:8086 --network mango_net --name SecureDEV avosa/secure:latest 
* $ docker logs secureDEV

### Logins
* admin@mango.avo : Admin4v0