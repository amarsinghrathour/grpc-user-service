
# gRPC User Service

A brief description of what this project does and who it's for
## Overview

This project is a gRPC-based User Service written in Go. It provides functionalities for managing user details and includes a search capability. The service uses SQLite with GORM for data storage and interacts via gRPC APIs.

## Prerequisites
Before running this project, ensure you have the following installed:

- Go 1.22+
- Protocol Buffers Compiler (protoc)
- Postman (for API testing)
- Docker (optional, for containerization)

### Setup
1. Install Go:

- Follow the official Go installation instructions for your platform: Golang Installation Guide.
2. Install Protocol Buffers Compiler:

- Download and install protoc compiler from Protocol Buffers Releases.
3. Install Postman:

- Download and install Postman from Postman Website.

## Directory Structure
grpc-user-service/ \
├── cmd/ \
│ └── server/ \
│ └── main.go \
├── database/ \
│ └── database.go \
├── factories/ \
│ └── user_factory.go \
├── models/ \
│ └── user.go \
├── proto/ \
│ └── user.proto \
├── repositories/ \
│ ├── user_repository.go \
│ └── user_repository_impl.go \
├── services/ \
│ └── user_service.go \
├── test/ \
│   └── main_test.go \
├── utilities/ \
│ └── common.go \
├── go.mod \
├── go.sum \
├── README.md \
├── Dockerfile \
└── Makefile


## Build and Run Locally

### 1. Clone the Repository
```shell
git clone https://github.com/amarsinghrathour/grpc-user-service
cd grpc-user-service
```
### 2. Build and Run the Application

#### build
```shell
go build -o server ./cmd/server/main.go

```

#### Run
```shell
./server

```
The gRPC server will start listening on port 50051.

### 3. Docker (Optional)
Build Docker Image

```sh
docker build -t grpc-user-service .
```
 Run Docker Container
```sh
docker run -p 50051:50051 grpc-user-service
```

### Using Postman for Testing
#### Postman Collection
You can use Postman to test the gRPC service. Import the Postman [Collection](https://crimson-comet-608249.postman.co/workspace/My-Workspace~de0bdb3d-86ea-496b-8d2c-3d4dbe3287aa/collection/6688f04cc0896f778dc8f4af?action=share&creator=21727383&active-environment=21727383-9984993a-f775-4c7b-bc3f-b1212d297a8d) to get started with example API calls.

## Running Unit Tests
To run unit tests for this project:
```shell
go test ./...
```
## License
This project is licensed under the GNU General Public License v3.0 - see the LICENSE file for details.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![AGPL License](https://img.shields.io/badge/license-AGPL-blue.svg)](http://www.gnu.org/licenses/agpl-3.0)

