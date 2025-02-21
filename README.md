# Sosmed Todolist API

Simple API for to-do list management posts on social media

## Tech Stack

- **Backend:**

  - Go (Golang)
  - Echo Framework
  - Zerolog
  - Swagger

- **Database:**
  - PostgreSQL

## BasePath

`/api/v1/`

## Documentation

`/api/swagger/`

## How to run

### Clone repository:

```sh
git clone https://github.com/agungramananda/sosmed-todolist
```

1. Make sure Docker is installed on your system.

2. Set up environment variables:
   Use the provided `.env.example`, rename it to `.env`, and add the following variables:

   ```env
   SVC_NAME=sosmed-todolist
   SVC_HOST=0.0.0.0
   SVC_PORT=8080
   SWAGGER_HOST=0.0.0.0
   SWAGGER_PORT=8080
   DB_HOST=postgres
   DB_PORT=5432
   DB_USERNAME=postgres
   DB_PASSWORD=db_pass
   DB_NAME=sosmed_todolist
   SERVER_PORT=8080
   ```

3. Run the application:

   ```sh
   make run
   ```

   Or

   ```sh
   docker compose build && docker compose up
   ```
