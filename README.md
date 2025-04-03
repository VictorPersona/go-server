

# Go Server

This is a simple HTTP server written in Go. It serves static files and provides basic functionality for handling HTTP requests.

## Features

- Serves static files from the `static` directory.
- Handles a `/hello` endpoint that responds with a greeting.
- Processes form submissions via the `/form` endpoint.

## Project Structure

```
/usr/local/go/src/go-server/
├── go-server
├── main.go
├── static/
│   ├── form.html
│   └── index.html
└── README.md
```

### File Descriptions

- **main.go**: The main Go file that contains the server logic.
- **static/index.html**: A simple static homepage.
- **static/form.html**: A form for submitting user data.
- **README.md**: Documentation for the project.

## Endpoints

### `/`
- **Method**: GET
- **Response**: Serves static files from the `static` directory.

### `/hello`
- **Method**: GET
- **Response**: Returns a simple "Hello" message.
- **Error Handling**: Returns a 404 error if the path is incorrect or the method is not `GET`.

### `/form`
- **Method**: POST
- **Response**: Processes form data and responds with the submitted name and address.
- **Error Handling**: Returns an error if the form cannot be parsed.

## How to Run

1. Ensure you have Go installed on your system.
2. Navigate to the project directory:
    ```bash
    cd /usr/local/go/src/go-server
    ```
3. Run the server:
    ```bash
    go run main.go
    ```
4. Open your browser and navigate to `http://localhost:8080`.

## Example Usage

- Visit `http://localhost:8080` to view the static homepage.
- Visit `http://localhost:8080/hello` to see the greeting message.
- Submit the form at `http://localhost:8080/static/form.html` to test form handling.

## License

This project is licensed under the MIT License.

