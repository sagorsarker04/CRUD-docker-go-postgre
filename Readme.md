# CRUD API Project

This project is a simple API for managing users, built with Go and PostgreSQL. It can be run in two environments: using Docker containers or a local environment. Both options have their own setup instructions, depending on your preference.

## Prerequisites

* **Go** (if running locally)
* **PostgreSQL** (if running locally)
* **Docker and Docker Compose** (if running in containers)

---

## 📦 Running the Project with Docker Containers

1. **Prepare the Environment**
   Ensure your `.env` file is correctly configured:

   ```env
   DB_HOST=db
   DB_USER=postgres
   DB_PASSWORD=123
   DB_NAME=my_db
   DB_PORT=5432
   ```

   * The `DB_HOST` should be set to `db` to match the service name in the `docker-compose.yml` file.

2. **Start the Containers**
   Run the following command to build and start the containers:

   ```bash
   docker compose up --build
   ```

   * This will create two separate containers: one for your Go backend and one for your PostgreSQL database.
   * If you have a local PostgreSQL server running, stop it first to avoid port conflicts:

   ```bash
   sudo systemctl stop postgresql
   ```

3. **Initialize the Database**
   Visit `http://localhost:8080/table` to create the necessary `users` table. This is required because the database in the Docker container will not have the same data as your local database.

---

## 💻 Running the Project Locally

1. **Install Dependencies**
   Make sure you have Go and PostgreSQL installed locally.

2. **Prepare the Environment**
   Update the `.env` file to connect to your local PostgreSQL:

   ```env
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=123
   DB_NAME=my_db
   DB_PORT=5432
   ```

3. **Start the Database**
   If your PostgreSQL server is not running, start it:

   ```bash
   sudo systemctl start postgresql
   ```

4. **Run the Project**

   ```bash
   go run main.go
   ```

---

## 🔄 Switching Between Local and Docker

* **To switch to Docker:**  Change `DB_HOST` to `db` in your `.env` file and stop your local PostgreSQL server.
* **To switch to Local:**  Change `DB_HOST` to `localhost` and start your local PostgreSQL server.

---

## 🛠️ Additional Features

* Auto table creation at `/table`
* RESTful routes for user management (`/users`, `/users/{id}`)
* Docker container isolation for consistent environments

---

## 📋 Notes

* Make sure the PostgreSQL server (local or containerized) is running before starting the project.
* Containerized database data is stored in the Linux subsystem, so it may differ from your local data.
