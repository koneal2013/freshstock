# FreshStock
## Introduction

FreshStock is a RESTful API server written in Go that uses the net/http package. It provides endpoints for managing produce inventory.

## Prerequisites

- Go `1.21`
- Docker (make sure docker commands do not need sudo)

## Endpoints

**[GET]**

- `/api/v1/produce` List all produce in the stock.
- `/api/v1/produce/:code` Get specific produce details.
- `/api/v1/produce/?q=<search_term>` Search produce inventory with a specific term in stock. Replace `<search_term>` with your actual search term.

**[POST]**

- `/api/v1/produce` Add new produce to stock.

**[DELETE]**

- `/api/v1/produce/:code` Remove produce from stock.

Please refer to the `api/routes.go` file for more details.

## Installation

1. Clone the repository: 
2. Run `make install`
3. Run `make run`

Make sure you have set your GOPATH and PATH environments correctly.

## Makefile Usage

- Compile the project: `make compile`
- Build docker image: `make docker-build`
- Run docker container: `make docker-run`

For more details, please refer to the Makefile.