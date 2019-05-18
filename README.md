# secure
Mango API: Secure

Secure handles user registration and login

## Run with Docker
* $ docker build -t avosa/secure:dev .
* $ docker rm SecureDEV
* $ docker run -d -e RUNMODE=DEV -e KEYPATH=/certs/ -e PUBLICKEY=fullchain.pem -e PRIVATEKEY=privkey.pem -p 8086:8086 --network mango_net --name SecureDEV avosa/secure:dev 
* $ docker logs secureDEV

### Logins
* admin@mango.avo : Admin4v0