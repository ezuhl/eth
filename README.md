Eth user app
-


This is a demo app that:

- Create user and returns a jwt
- Creates api token
- Uses that api token to get average head block rate per minute from  beacon node


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
- Get .env file for settings
- go run ./cmd/postgres/main.go [start postgres]
- go run ./cmd/eth/main.go [start app]
- use Postman to make requests