# Project Template

## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) and server configuration structure code by proto file
make buf
# Generate OpenAPI v3 documentation for the API
make openapi
# Generate ent code
make ent
# Generate wire code
make wire
# Build program
make build
# Build docker image
make docker
# Generate all files
make all
```

