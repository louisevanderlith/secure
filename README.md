# secure
Mango API: Secure

Secure handles user registration and login

## Run with Docker
* $ docker build -t avosa/secure:dev .
* $ docker rm SecureDEV
* $ docker run -d -e KEYPATH=/certs/ -e PUBLICKEY=fullchain.pem -e PRIVATEKEY=privkey.pem -p 8086:8086 -v $(pwd)/db/:db/ --network mango_net --name SecureAPI avosa/secure:dev 
* $ docker logs secureDEV

### Logins
* admin@mango.avo : Admin4v0