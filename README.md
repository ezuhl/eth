Eth user app
-


This is a demo app that:

- Creates a user and returns a jwt
- Creates api key
- Uses that api token to get average head block rate per minute from beacon node


Api
--
- POST {host}/user/create
  - body
  - {
    "username":"uname",
    "password":"testme"
    }


- GET {host}/user/key
  - Header: Authorization: Bearer {jwt token}

- GET {host}/chainhead/avgheight/{api_key}
  - Header: Authorization: Bearer {jwt token}


Startup:
- Check out code to $GOPATH/src/github.com/ezuhl
- Get .env file for settings and place in /env folder
- go run ./cmd/postgres/main.go [start postgres]
- go run ./cmd/eth/main.go [start app]
- use Postman to make requests


Db Migration:
- A database migration will run when you run the app.  Make sure nothing, other than EthPostgress instance, is license on the Postgres port