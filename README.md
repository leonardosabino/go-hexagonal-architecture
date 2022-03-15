# GO - Hexagonal Architecture Template

This service is a template for hexagonal architecture in GO.

## Architecture 
The project is built as follows:


## Folders Structure

This repository contains three main folders: `cmd`, `build` and `internal`.

The `internal` folder contains all the go code, modules and tests that compose the
service.

The `build` folder contains the Dockerfiles used for building the containers, 
the docker-compose file with the description of the local development environment 
and kubernetes manifest files that are required for deploying this component.

## Building

### Source code

To build the source code at the root folder, execute:

```bash
make build
```

The binary executable file will be generated with the name `main` at root folder

### Container

At main folder you can run:
```bash
docker-compose -f build/package/docker/localstack/docker-compose-local.yml up -d
```

## Testing

To test the project at the root folder, execute:
```bash
make test
```
