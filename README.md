# CA-Project
Base repository for the CA project. This project is a combination of Docker, Go, and Polymer to provide a web interface for starting, stopping, adding, and removing Docker containers.

## Saturday 
#### 2PM
Kids are sleeping and I have come up with an idea for the project. I am now learning about Docker Remote API so that I can list, start, and stop docker images. 

Step 1, create base environment.
main.go
 - Starts mysql container
 - Starts api server
  * api server creates database & populates it with container information.
 - Starts web server
  * web server queries mysql for container information
  * web server has option to start and stop containers

#### 2:30
I am able to start the mysql container using main.go. Next step is to create a Go app container that can connect to the mysql server. We want to use the -link feature. I will need to learn how to expose the Docker remote api to the container network. Github issue dated from 2013-2014 states its a feature coming in version 0.8. 

#### 3:30
Kids woke up early from a nap, it's been a circus but I have found Alpine Linux for Docker. I have created the base Go app for the API and it's size is only 11MB with the binary already compiled. Next step is connecting to the mysql database container from the api container. https://github.com/gliderlabs/docker-alpine

To create civis-mysql
docker stop civis-mysql && docker rm civis-mysql
docker run --name civis-mysql -e MYSQL_ROOT_PASSWORD=civis -d mysql

To remove, build and create web/api
docker stop web1 && docker rm web1 && docker rmi ca-web
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t ca-web -f Dockerfile .
docker run --name web1 --link civis-mysql:mysql -d -v /var/run/docker.sock:/var/run/docker.sock ca-web

#### 5:30
API container is running, connects to the mysql container, and pulls data out of the database. Tomorrow I will resume and build api routes.

## Sunday
#### 10:45 AM
Resuming work. The plan is to connect the api server to the Docker host's Remote API, get a list of all containers and their states, and use that to populate the mysql database on startup. I think Redis would be a better fit than mysql, but continuing with mysql due to scope and time restraints. 

#### 12:30 PM
After much rustling with mysql I have better database initilization. If the database doesn't exist, create it. If the tables don't exist, create it. We will likely have multiple API servers for load balancing so these checks on start up are necessary. We also want to initialize the table with the current list of containers so it makes sense to keep the initialization in the api app instead of writing a Go app for the mysql container. We'll keep the mysql container dumb and vanilla. 

I have spent more time refreshing myself on interacting with MySql via Go than I would have liked. I miss Google's Datastore/Cloud SQL api which makes interacting with the services trivial. I switched from Go's database/mysql package to "github.com/jmoiron/sqlx" which allows for inserting whole structs, but the queries are still quite verbose and saddening. Another time I will explore for a mature Go ORM again. 

#### 3:30 PM
Had to break for kids. I now have the endpoint `GET /containers` and `GET /images` which will return an array of all containers. Successful test! 

```
[{"Id":"2059dd4b08a98de82c4d6dc7159ac4b5eabaee9c79795e53d14787d532110ebc","Image":"ca-api","Command":"/main","Created":1439752734,"Status":"Up About a minute","Names":["/api1"]},{"Id":"09779df7805ab4faa22fa890c218ce45352c58226c072a4ee6d6705e635cecba","Image":"mysql","Command":"/entrypoint.sh mysqld","Created":1439742552,"Status":"Up 2 hours","Ports":[{"PrivatePort":3306,"Type":"tcp"}],"Names":["/api1/mysql","/civis-mysql"]}]
```

#### 4:00 PM
I'm making a pivot in the project. The previous setup didn't allow for sensible scaling. New setup will have the ca-master app, which runs on the docker host, provide the management screen and api. The containers will be dumber and I'll create "busy work" for them to do to replicate load. The ca-webserver is now up and running and responds to GET /. I'll build a small Polymer app for it to reply with at the end. This is the least important aspect of the project.

#### 5:00 PM
ca-master now has options to stop and start containers when it starts. I plan to add a config file that will be read at startup. Stopping for the day.

## Monday 10:30 AM
Work is resuming. First step is to get the config file working.

