## Sqlc
https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html
- sqlc generate


## oapi
https://github.com/oapi-codegen/oapi-codegen
# for the binary install
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

oapi-codegen --package=main --generate types,client specapi.yaml > clientapi.go  

oapi-codegen --package=main  --generate types,server specapi.yaml > generatedserver.go