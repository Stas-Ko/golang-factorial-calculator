# Golang Factorial Calculator

This repository contains a simple Golang application for calculating factorials concurrently using goroutines. The API exposes an endpoint for calculating factorials based on input JSON data. The project utilizes the httprouter library for handling HTTP requests.

## Task

1. Create a public repository on GitHub or any other public code repository website.
2. Create a REST endpoint called calculate available at port 8989 so it can be accessed at http://localhost:8989/calculate.
3. Use https://github.com/julienschmidt/httprouter for creating a server.
4. The calculate endpoint should take JSON with the following structure: `{"a": int, "b": int}` and calculate factorial of a and b using goroutines (https://en.wikipedia.org/wiki/Factorial).
5. The calculate endpoint will return JSON with the a! and b!.
6. Create middleware to check if a and b exist in JSON, and they are non-negative integers. In case of failure, return JSON: `{"error": "Incorrect input"}` with an error status code 400 Bad Request.

## Running the Server

1. Open the command prompt (cmd) in the project directory.
2. Run the command:
   ```bash
   go run main.go

This will start the server. You will see the message "Server is running on port 8989..." in the console, indicating that the server is successfully running.

## Sending HTTP Requests

Now, in the command prompt (cmd), use the curl command to send HTTP requests. For example, to send a POST request to your server:

curl -X POST -H "Content-Type: application/json" -d "{\"a\": 3, \"b\": 5}" http://localhost:8989/calculate

This will return a JSON response, for instance:
{"a_factorial": "6!", "b_factorial": "120!", "a_factorial_value": 6, "b_factorial_value": 120}

If you input negative values, the server will respond with:
{"error": "Incorrect input"}

Feel free to explore and test the application using different inputs.


## Building from Source

If you want to build the application from source, follow these steps:

1. Ensure you have Go installed on your machine.

2. Clone the repository to your local machine:
   ```bash
   git clone https://github.com/Stas-Ko/golang-factorial-calculator.git


   
