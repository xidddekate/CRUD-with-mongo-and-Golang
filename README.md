# CRUD-with-mongo-and-Golang
A simple Dockerized Backend application written in golang using MongoDB as a database demonstrates how basic CRUD operations can be performed along with proper unit testing of each API

# Requirements
* Docker for MongoDB server and API server <br /> 
* Golang installed on machine for development purpose <br /> 

# How to run application
Its very simple to run this application. Thanks to docker containers that takes care of everything.<br /> 
* Spinup docker containers using following command<br /> 
    ```docker-compose up --build```
* And Thats it, our servers will be ready in few minutes as per the configuration specified in [config](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/config) folder and [Dockerfile](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/Dockerfile)
* Backend API server will be listening on port 8080 and mongo server on port 27017 as specified in [docker-compose.yml](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/docker-compose.yml) file
* To run ALL unit Tests, after docker containers are up and running, run the following command in new terminal from root repository<br/> ```go test -v ./...```
* To run a unit test for a particular API (i.e service files in [handlers](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/handlers) folder) use following . <br/> &nbsp;&nbsp;&nbsp; eg. for Get users API which is handled by [getUser.go](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/handlers/getUser.go) file the command will be   <br/>  &nbsp;&nbsp;&nbsp;&nbsp; ```go test -v ./... -run ^TestGetUser```
# Code and folders overview
* [config](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/config) : Has all necessary configurations for environments and Mongo database
* [init-mongo.js](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/init-mongo.js) : To initialize the Mongo Collections
* [models](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/models) : Has all the required Schema that is being followed for this project
* [handlers](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/handlers) : Has all the the functions to handle API calls and their corresponding test files for unit testing purpose
* [database](https://github.com/xidddekate/Crud-with-mongo-and-Go/tree/main/database) : Has files to to interact with database 
  * [database.go](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/database/database.go) : To establish a connection with MongoDB
  * [users.go](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/database/users.go) : for performing CRUD operations on database
  * [users_mock.go](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/database/users_mock.go) : for Mocking database interaction functions written in [users.go](https://github.com/xidddekate/Crud-with-mongo-and-Go/blob/main/database/users.go) inorder to enable unit testing.



