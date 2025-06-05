# Movie API Tutorial

## Prerequisites

- Go installed on your machine. You can download it from [here](https://golang.org/dl/).
- Basic understanding of Go programming language.

## Getting Started

1. **Clone the repository to your local machine:**

   ```bash
   git clone https://github.com/vishal130195/go-movies-crud.git

   cd movie-api
   ```

2. **Run the application:**

   ```go
   go run main.go
   ```

   The server will start on `http://localhost:8000`.

   **Note:** Ensure that you have [Go](https://golang.org/dl/) installed on your machine.

## Usage

### 1. Get All Movies

- **Endpoint:** `GET /movies`
- Retrieve information about all movies.

### 2. Get Movie by ID

- **Endpoint:** `GET /movies/{id}`
- Retrieve information about a specific movie using its ID.

### 3. Create a New Movie

- **Endpoint:** `POST /movies`
- Create a new movie entry.

   ```json
   {
     "isbn": "987654321",
     "title": "Inception",
     "director": {
       "firstname": "Christopher",
       "lastname": "Nolan"
     }
   }
   ```

### 4. Update a Movie

- **Endpoint:** `PUT /movies/{id}`
- Update information about a specific movie using its ID.

   ```json
   {
     "isbn": "987654321",
     "title": "Inception",
     "director": {
       "firstname": "Christopher",
       "lastname": "Nolan"
     }
   }
   ```

### 5. Delete a Movie

- **Endpoint:** `DELETE /movies/{id}`
- Delete a specific movie using its ID.

## Examples

### Get All Movies

```bash
curl http://localhost:8000/movies | jq
```

### Get Movie by ID

```bash
curl http://localhost:8000/movies/1 | jq
```

### Create a New Movie

```bash
curl -X POST -H "Content-Type: application/json" -d '{"isbn":"123456789","title":"Interstellar","director":{"firstname":"Christopher","lastname":"Nolan"}}' http://localhost:8000/movies | jq
```

### Update a Movie

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"isbn":"123456789","title":"Interstellar Updated","director":{"firstname":"Christopher","lastname":"Nolan"}}' http://localhost:8000/movies/1 | jq
```

### Delete a Movie

```bash
curl -X DELETE http://localhost:8000/movies/1
```

## Running the Application

To run the application locally, follow these steps:

1. Open a terminal.
2. Navigate to the project's root directory.
3. Run the following command:

   ```bash
   go run main.go
   ```

The server will start, and you can access the API at `http://localhost:8000`.

Feel free to explore and experiment with these endpoints using a tool like `curl` or Postman.

## Conclusion

Congratulations! You have successfully set up and used the Movie API. Feel free to customize and extend the code according to your requirements.