## CRUD API Project with Dockerized PostgreSQL Database

This project is a simple implementation of a CRUD API using the Go programming language and a PostgreSQL database. The database is run inside a Docker container to simplify the development environment and deployment.

# Prerequisites 
Before getting started, make sure you have the following components installed:

1. Docker: https://www.docker.com/
2. Go: https://golang.org/

# Configuration

1. Clone the project repository
    ```
    git clone https://github.com/mariobenissimo/RestApiPost.git
    ```
3. Navigate to the project directory:
   ```
    cd crud
    ```
5. Start the Docker container for the PostgreSQL database:
    ```
    cd deployments
    docker compose up
    ```

# Usage
* Start the Go application:
    ```
    cd cmd
    go mod run .
    ```
* The application will be running at http://localhost:8000. You can use a tool like cURL or Postman to send requests to the API.

# API Endpoints
* GET /movies: Returns a list of all movies in the database.
* GET /movies/{id}: Returns the details of a specific movie based on the ID.
* POST /movies: Creates a new movie in the database.
* PUT /movies/{id}: Updates the details of an existing movie based on the ID.
* DELETE /movies/{id}: Delete a movie from the database based on the ID.
