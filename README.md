# FreshStock

## Introduction

FreshStock is a RESTful API server written in Go that uses [Gin-gonic's Gin](https://github.com/gin-gonic/gin) web framework. It provides endpoints for managing produce inventory. A simple React UI is included for interacting with the API.

## Prerequisites

- Go `1.21.5`
- Node.js (for the UI)
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

Please refer to the `internal/api/routes.go` file for more details.

## Installation & Usage

### Backend (API)

1. Clone the repository:
2. Run `make run` to start the API server locally.
    - Alternatively, you can run the API in a docker container with `make docker-run`.
3. The API server will be listening on port `8080`.
4. Run `make docker-down` to stop and remove the docker container.

### Frontend (React UI)

1. Navigate to the `ui` directory:
   ```sh
   cd ui
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the development server:
   ```sh
   npm run dev
   ```
4. Open your browser and go to the URL shown in the terminal (usually `http://localhost:5173`).
5. The UI will interact with the API at `http://localhost:8080/api/v1/produce` by default.

## Features of the React UI

- List all produce in inventory
- Search produce by name
- Add new produce (code, name, unit price)
- Delete produce by code
- Error handling and loading indicators

## Notes

- Make sure the backend API is running before using the UI.
- CORS is enabled by default in the backend for development.
- All produce data is stored in memory and will be lost when the server restarts.

## Makefile Usage

- Compile the project: `make compile`.
- Clean previous builds: `make clean`.
- Run tests: `make test`.
- Build docker image: `make docker-build`.
- Run docker container: `make docker-run`.

For more details, please refer to the Makefile.

## Assumptions

- All incoming data is JSON and must include code, name, and unit price.
- The produce code is a string of alphanumeric characters and hyphens, name is alphanumeric, and unit price is a number with up to two decimal places.
- Errors are returned as JSON with an error message and appropriate HTTP status code.
- The API is RESTful and uses standard HTTP methods and status codes.

