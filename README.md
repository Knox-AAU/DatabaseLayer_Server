# DatabaseLayer_Server
Go REST API with CRUD operations for Knox database

## Documentation
### Generate new documentation based on code
Run `swagger generate spec -m -o ./swagger.yaml` from the terminal or directly from the `main.go` file.

After generating the yaml file, run `redocly build-docs swagger.yaml` from the terminal, which will generate the updated html docs.