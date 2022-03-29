# Building a Production Ready REST API in Go


## Prerequisites

In order to be able to process some commands on a host machine you would need to have installed the below packages:
* GoLang
* Docker

## Preparation
* Run `docker-compose up --build -d` to start the database and application

## Some useful command

### Run below command for e2e test
````
go test ./... -tags=e2e -v
````