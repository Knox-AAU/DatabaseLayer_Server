# DatabaseLayer_Server

Go REST API with CRUD operations for Knox database

## Documentation

### Generate new documentation based on code

Run `swagger generate spec -m -o ./swagger.yaml` from the terminal or directly from the `main.go` file.

After generating the yaml file, run `redocly build-docs swagger.yaml` from the terminal, which will generate the updated html docs.

## Accessing AAU KNOX server

Requires to be accessed via AAU's network, or its VPN.

### Access Database layer API

Your port is 8000 and the API on the server is on 8081.

ssh `<studiemail>`@knox-kb01.srv.aau.dk -L 8000:localhost:8081

### Access Virtuoso database

The code (including tests) accesses Virtuoso on port 8890, which is the same port to access it on the server.

ssh `<studiemail>`@knox-kb01.srv.aau.dk -L 8890:localhost:8890
