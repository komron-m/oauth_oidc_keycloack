## Getting started

```shell
# clone repo
mkdir -p $GOPATH/src/github.com/komron-m/
cd $GOPATH/src/github.com/komron-m/
git clone git@github.com:komron-m/oauth_oidc_keycloack.git
cd oauth_oidc_keycloack

# install project dependencies
go mod tidy

# init app dependencies
docker-compose up

# run backend application
go run ./cmd/game_app/

# open another terminal and start client application
cd $GOPATH/src/github.com/komron-m/oauth_oidc_keycloack/clients/web
yarn install
yarn start
```

## Contributing

Pull requests are welcome. For any changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
