# GoJSONParserAPI

GoJSONParserAPI is a lightweight, efficient API server written in Go, designed to demonstrate the parsing of complex JSON payloads. It provides a RESTful endpoint that accepts JSON data, showcasing advanced parsing techniques including handling nested objects and arrays.

## Features

- **JSON Parsing**: Robust handling of various JSON structures with support for nested objects and arrays.
- **RESTful Endpoint**: A dedicated POST endpoint to receive and process JSON data.
- **Custom Data Handling**: Template for processing and responding with parsed data.
- **Error Handling**: Comprehensive error responses for invalid requests or data.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

### Installing

To set up the project locally, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/GoJSONParserAPI.git
    ```
2. Navigate to the project directory:
    ```bash
    cd GoJSONParserAPI
    ```
3. Run the server:
    ```bash
    go run main.go
    ```

The server should now be running on `http://localhost:8080`.

### Usage

To use the API, send a POST request to `/person` with a JSON payload. For example:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","age":30,"city":"New York","friends":[{"name":"Jane Doe","age":25},{"name":"Richard Roe","age":28}]}' http://localhost:8080/person
Contributing
Contributions are welcome! Please feel free to submit a pull request.

License
This project is licensed under the MIT License - see the LICENSE file for details.
