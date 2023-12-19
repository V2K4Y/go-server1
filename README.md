# My Go Server

This is a simple Go server that provides two HTTP endpoints for sorting arrays of integers. The server can perform sorting either sequentially or concurrently using goroutines.

## Endpoints

### Process Single
- URL: `/process-single`
- Method: POST
- Description: Sorts each sub-array sequentially.

### Process Concurrent
- URL: `/process-concurrent`
- Method: POST
- Description: Sorts each sub-array concurrently using goroutines and channels.

## Running Locally

To run the Go server locally, follow these steps:

1. Clone the repository:
  ```bash
  git clone https://github.com/V2K4Y/go-server.git
  cd go-server
  ```
2.Run the server:
  ```bash
  go run main.go
  ```
Open your web browser and visit http://localhost:8000.



## Dockerization
### To containerize the Go server using Docker, follow these steps:
1.Build Docker image:
  ```bash
  docker build -t my-go-server .
  ```
2.Run the Docker container:
  ```bash
  docker run -p 8000:8000 my-go-server
  ```



## DockerHub
You can pull the image and run the container as follows:
  ```bash
  docker pull vky24/my-go-server:latest
  docker run -p 8000:8000 vky24/my-go-server:latest
  ```

   
