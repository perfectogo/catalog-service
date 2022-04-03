# catalog-service
## book-shop catalog-service for template

### Migration
#### Migrations to create a new migration:
```
migrate create -ext sql -dir migrations -seq create_catalog_tables
```

* create file *"run-migration.sh"* and write:
```
    #!/bin/bash
    # CURRENT_DIR=$(pwd)

    # mkdir $CURRENT_DIR/migrations

    export POSTGRES_USER=xxxx
    export POSTGRES_PASSWORD=xxxx
    export POSTGRES_DATABASE=xxxx

    migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path="./migrations" up
```
### Protocol Buffers for Go with Gadgets
Install the protoc-gen-gofast binary:
```
    go get github.com/gogo/protobuf/protoc-gen-gofast
```
or:
```
  go get github.com/gogo/protobuf
```
Use it to generate faster marshaling and unmarshaling go code for your protocol buffers.
```
    protoc --gofast_out=. myproto.proto
```