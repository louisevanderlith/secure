# secure
Mango API: Secure

Secure handles user registration and login

## Run with Docker
* $ docker build -t avosa/secure:dev .
* $ docker rm SecureDEV

* $ docker run -d -v $pwd/db/:/db/ --network host --name Secure avosa/secure:dev 
* $ docker logs secureDEV

### Logins
* admin@mango.avo : Admin4v0

### cUrl
* Obtain Token
curl -XPOST -H 'Authorization: Basic bWFuZ28ud3d3OnNlY3JldA==' -H "Content-type: application/json" -d '{"UserToken":"", "Scopes": []}' 'http://localhost:8086/token'
* View Info on Token
curl -XPOST -H 'Authorization: Basic bWFuZ28ud3d3OnNlY3JldA==' -H "Content-type: application/json" -d '{"AccessCode":"5da20cfd3aff..."}' 'http://localhost:8086/info'



curl -XPOST -H 'Authorization: Basic bWFuZ28ud3d3OnNlY3JldA==' -H "Content-type: application/json" -d '{"AccessCode":"71044c66ad6fd38706c86535fee9aa00f40e4d0277f64b6fda0bf37ddc46cf5acc43d911a0a9653987e7934bfb2b467253b7fb34988bf4ce01846328b8b9e98cef3c7cbed4f6a0bc781959caab32192914ffd37508cc84352dc671293426b855e6713a653e862339d0a2eac6c01ff782b65a5c48ffbb0a08243d1977697e7de207122eec0ee6f0e9723e47bbd8eeedf2174d39c53071b08bfd627f0416f8895dfeba2008672bb613fb68014656272f35564660de4b93dde2c87155da7a87f275a0224b402ecdfc2234c5ef5383d4ab0b53288614771d5a6e74fa69764c0a5601f13209e03b879fcac59db0b8176172e6d1a32419b0b38a1d22f03021ddbc9eff256b102b061ea7f97b405589e4b511f0173a2c2b3ae3e6683b08ad828050e6ac54c1d4000a9ce4d98275ac60a2a45f6658598a4ef9b9f7248852058452c8bd61102603b749518d2e132784a5e8162fa4aa5994bc64f0d83c5de83bc9bca507334893e7641bc63669ed6dc2521b4a9c7d8365f6ac16320c003f208c9e7c8c7c6ef5348dd72b3cf12351299406713a8850246fc4af9ba527a85e4c31c1ecf58be58f97b5d626b10ad325e3761331cc96012b2338e0669fb0b6e228a7129a22cfed473d1023e6b4ace56f5c843b82bef42b25a615fc8388aee991ce000268d2f12620c9685f00152b7a9dcd00dfeb428bcd55719fb16df6cd1e2fd5f56aac24aaee"}' 'http://localhost:8086/info'
