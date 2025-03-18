# Strata

## Docker Setup
The root of the project contains a Dockerfile and compose.yml file. 
This allows a quick installation using the following command within this project's root source code directory.
<br>
To run the following commands, you must have [Docker-Compose](https://docs.docker.com/compose/install/) installed. Alternative, you can use the [Go Setup](#go-setup) procedure. 

```sh
docker compose up
```
The container can be shut down and removed, while keeping the built image, using the following command:

```sh
docker compose down
```

## Go Setup
The following set of commands can be used to install required packages, then build the 'strata' executable in the project root directory. These commands require your terminal's current working directory is the project's root directory.  
```sh
# Install packages:
go install ./cmd

# Build project
go build -o strata ./cmd

# Result is placed in ./strata
``` 
