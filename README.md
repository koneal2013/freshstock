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
2. Run `make run`
    - Alternatively, you can run the API in a docker container with `make docker-run`.
3. The API server will be listening on port `8080`.
4. Run `make docker-down` to stop and remove the docker container.

Make sure you have set your GOPATH and PATH environments correctly.

## Makefile Usage

- Compile the project: `make compile`.
- Clean previous builds `make clean`.
- Run tests `make test`.
- Build docker image: `make docker-build`.
- Run docker container: `make docker-run`.

For more details, please refer to the Makefile.

## Assumptions

While developing this program, several assumptions were made, particularly in the areas of input validation and error handling:

1. Input Validation: It was assumed that all incoming data would be in the form of JSON and that the client would provide all the necessary fields (name, code, and unit price) when adding a new produce item. The program expects the produce code to be a string of alphanumeric characters and hyphens, the name to be an alphanumeric string, and the unit price to be a number with up to two decimal places. If these conditions are not met, the program will return an error. It was also assumed that the client would provide a valid produce code when fetching or deleting a produce item, and a valid query string when searching the produce inventory.
2. Error Handling: The program assumes that errors may occur during the processing of requests, such as when a client tries to add a produce item with a code that already exists in the inventory, or when a client tries to fetch or delete a produce item that does not exist. In such cases, the program will return an appropriate HTTP status code and a JSON response with an error message. It was also assumed that the program may receive invalid JSON or malformed requests, in which case it will return a 400 Bad Request status code and an error message.
3. Concurrency: The program assumes that multiple clients may try to read from or write to the produce inventory concurrently. To handle this, the program uses a read-write mutex to synchronize access to the inventory. This allows multiple concurrent reads for high throughput, while ensuring that writes (additions and deletions) are atomic and consistent.
4. In-Memory Storage: The program assumes that the produce inventory can be stored in memory for the lifetime of the program. This means that all produce data will be lost when the program is terminated.
5. RESTful API: The program assumes that the client is familiar with the principles of REST and will use the appropriate HTTP methods (GET, POST, DELETE) to interact with the API. The program uses standard HTTP status codes to indicate the success or failure of a request, and it assumes that the client will check these status codes and handle them appropriately.

