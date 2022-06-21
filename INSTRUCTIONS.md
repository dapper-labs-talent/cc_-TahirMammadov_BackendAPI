## Instructions 

You will need to install Docker and Go in order to run this app.

To install docker -> https://docs.docker.com/get-docker/ 
To install golang -> https://go.dev/doc/install


1. Change directory `cd take_home_assignment/`
2. Before starting the app, we need to create a DB instance and migrate. In order to do that please run `docker-compose up -d`
3. Once we run migration successfully, we can run `make run` to start the app
4. (Bonus) In order to run unit tests, please go to tests files (e.g `cd ./tests`) and run `go test -v -run TestApp`
