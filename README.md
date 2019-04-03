# TeamGen

[![Build Status](https://travis-ci.com/anthonydevelops/team-gen.svg?branch=master)](https://travis-ci.com/anthonydevelops/team-gen/)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtkp9993/SimpleCRUDApp)](https://goreportcard.com/report/github.com/anthonydevelops/team-gen)
[![GitHub license](https://img.shields.io/badge/license-Apache%20license%202.0-blue.svg)](https://github.com/anthonydevelops/team-gen/blob/master/LICENSE.md)

TeamGen is a full-stack web application for finding the perfect team to be a part of. Intended for a class setting, the user is able to design a profile,choose exactly what they look for in teammates, and find or create their own team based off their given preferences. 

# Installation

```bash
# Clone the repo
git clone https://github.com/anthonydevelops/team-gen
cd /team-gen

# Create config/config.toml file with login credentials based off config/example_config.toml
touch config/config.toml

# Run Postgres, set environment if needed with -e (POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD)
# Also, create a local folder to have persistance of docker db image when developing
docker run -p 5433:5432 -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=test --name pg --rm postgres

# Run server to test connection & begin development
go run main.go
```

# Documentation

Detailed documentation can be found in the [docs](https://github.com/anthonydevelops/team-gen/docs) folder.

# Architecture

TeamGen is built utilizing [ReactJS](https://reactjs.org/) as the component library for the front-end, [Go](https://golang.org/) for our backend server language to connect to [PostgreSQL](https://www.postgresql.org/), and lastly [Docker](https://www.docker.com/) to build in a microservice-oriented way to improve many aspects of the application. 

_TODO Create architecture layout_
