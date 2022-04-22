# Port

## How to run this application

Before run your aplication you need to running the postgres database with *docker-compose.yml* using 

```
docker-compose up
```

After that, you can process a json file with ports information passing the file path with the make command 

```
make run file=[json_path]
```

Below an example

```
make run file=ports.json
```

## How to run the tests

For running the tests, use the command below

```
make test
```

## About project structure

- *cmd* -> this folder contains all aplication entrypoints
- *domain* -> this folder contains all about  port domain
- *entity* -> this folder contains the port entity
- *infrastructure* -> this folder contains all about infrastructure layer like database and cli
- *interfaces* -> this folder contains the interfaces that communicates with external world
- *migrations* -> here contains all migration scripts that running on docker-compose up command
- *service* -> this folder contains all business logic
- *docker-compose.yml* -> contains the postgres services
- *Makefile* -> commands to run and tests the applications
