# Todo CRUD REST API 📝

This project is a simple RESTful CRUD (Create, Read, Update, Delete) API for managing todos. It provides basic operations to interact with a todo list, allowing users to create, read, update, and delete todo items.

I built this project to get started with Go.

## API Endpoints 🚀

1. **Get all todos**
   - Endpoint: `GET /todo`
   - Description: Retrieve a list of all todos.

2. **Create a new todo**
   - Endpoint: `POST /todo`
   - Payload:
     - Name (string): The name of the todo.
     - Description (string): A description of the todo.
     - Completed (boolean): Indicates whether the todo is completed.
     - ID (number): Unique identifier for the todo.
   - Description: Create a new todo with the provided information.

3. **Update a todo by ID**
   - Endpoint: `PUT /todo/{id}`
   - Payload:
     - Name (string): The updated name of the todo.
     - Description (string): The updated description of the todo.
     - Completed (boolean): The updated completion status of the todo.
   - Description: Update an existing todo identified by its ID with the provided information.

4. **Delete a todo by ID**
   - Endpoint: `DELETE /todo/{id}`
   - Description: Delete a todo item based on its unique ID.

## Project URL 🌐

Visit the live deployment of this API: [Todo CRUD API](https://go-todo-crud-production.up.railway.app/todo)

## Getting Started 🛠️

To run this project locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/mecies/go-todo-crud.git
   ```

2. Install dependencies:

   ```bash
   cd go-todo-crud
   go install
   ```

3. Create a .env file and populate the following vars (requires a Postgres DB)
    ```bash
    DB_USER=
    DB_PASS=
    DB_HOST=
    DB_NAME=
    DB_PORT=
    PORT=
    ```

4. Run the project:

   ```bash
   go run main.go
   ```
   

   The API will be accessible at `http://localhost:{PORT}`.

## License 📜

This project is licensed under the do-whatever-you-want-license.
