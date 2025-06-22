# HTTP SERVER TCP

Minimal HTTP server over raw TCP, written in Go to explore low-level networking and HTTP parsing without using `net/http`.

## Overview

This project demonstrates how to implement a basic HTTP/1.1 server from scratch using Go's net and bufio packages — without relying on the net/http standard library.

## Features

- Raw TCP server built with `net` and `bufio`
- Manual parsing of request line, headers, and methods
- Basic routing for `GET`, `POST`, `PUT`, and `DELETE`
- CI pipeline with GitHub Actions
- Fully containerized with Docker
  
## 🐳 Run with Docker

```bash
docker build -t go-http-server .
docker run -p 8080:8080 go-http-server
```
## 📬 Example Requests (with curl)

Test the server locally or in Docker using these commands:

```bash
# GET /
curl http://localhost:8080/

# POST /submit
curl -X POST http://localhost:8080/submit

# PUT /update
curl -X PUT http://localhost:8080/update

# DELETE /delete
curl -X DELETE http://localhost:8080/delete

# Invalid path (should return 404)
curl http://localhost:8080/invalid
