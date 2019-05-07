# Mavencode-assignment
What is this? It's a coding assignment I did as part of an interview process. It's a simple golang microservice that stores and queries 'person' objects via RabbitMQ and Redis. I've reproduced the assignment in a later section of this document, [here](#Assignment).

# Running the service
I've packaged the service up using docker-compose. To run it, you need to have docker-compose and golang installed. With those prerequisites, do the following:
* Clone this repository
* From the root directory of the repository, run `docker-compose build`, then `docker-compose up`. You should have four containers running on your system: one Redis, one RabbitMQ, one with the webserver, and one with the redisclient.
* To see what's going on, point your browser at http://localhost:8080. That should get you a 'Welcome' message.
* If you want to watch what's going on in redis, you can install redis-cli and run it with the monitor command `redis-cli monitor` - I've set the docker container up so that redis should be exposed to the host.
* Without data, there's not much to look at. In order to get data in the system, you can do the following:
    * From the project root, cd to test
    * execute the file `make_test_data.sh`. This file will build the test data generator locally and execute it. Executing it will create 1000 random person records having random timestamps within the past hour. It does this by generating records via the [go-randomdata](https://github.com/Pallinder/go-randomdata) library and sending them via PUT requests to the endpoint at http://localhost:8080/person. 
        * note: because of the limited number of options for firstname and lastname in the go-randomdata library, generating 1000 names at random will likely have many duplicate names. This means that fewer than 1000 distinct records will be added to the database. This is due to the choice of firstname-lastname for keys. (I called this out as an issue in the Design Decisions section of this document.)
    * With the test data in place, you can visit the following urls to see the data: http://localhost:8080/persons/60sec, http://localhost:8080/persons/5min, http://localhost:8080/persons/1hr
    * Additionally, you can visit http://localhost:8080/persons to retrieve all of the records (due to the mandated TTL, you should only see records for the last hour). Also, http://localhost/person/firstname-lastname can be used to query an individual person record (assuming it's still in the cache). Neither of these were in the assignment, but they weren't too hard to add, and are nice to have.
## General implementation / code tour
The main functional parts are: the web server, the redis client, the redis server and the RabbitMQ server. The Redis and RabbitMQ servers are just standard docker containers running their respective services. These are managed via their settings in the docker-compose.yml file. The web server and redisclient are both written in golang. There was a fair bit of shared code between them, so I put that in a separate library, under the /shared directory. The non-shared bits are in the /webserver and /redisclient directories. Configuration is set up with reasonable defaults (based on running everything on one host), and overrides are passed in via environment variables, with help from the he viper library. This is used to set things up properly in the docker-compose environment.

## Limitations / to-do list
As this was intended to be a quick demonstration exercise, there are lots of limitations. Here are a few:
* Error handling is not very robust. In particular, I'm not returning good statuses / messages on failed calls. This would be the next thing to work on.
* There is minimal validation of data.
* There is very little security around the services - in a production environment, we'd need to do more to secure rabbit, redis, and the webserver.
* I need to add unit/integration tests to the code. 
* The system could use some refactoring - some of the naming is less clear than it could be.

## Assignment

### Create a MicroService in Golang With:
1. Restful endpoint to post a JSON => {firstname, lastname, address, gender, timestamp}
2. Restful endpoint to show you a list of {firstname, lastname, address, gender, timestamp} with query to filter for the last 60s, 5mins or 1hr

*array must be ordered by timestamp

### Process
1. When the POST request is done, the JSON must be pushed to a Message Queue { Use any queue of your choice in a Docker container}
2. Implement a Consumer to read the Queue and combine the firstname and lastname into fullname and push it to a Redis (Also a Redis Docker)
3. Query the Redis Queue and return the list of users with query to filter for the last 60s, 5mins or 1h

*array must be ordered by timestamp, The Redis Queue should have a TTL for the dataset inserted of not more than 1hr


Implement this in Golang and Push it to Github for review
   
# Implementation

## Design decisions:
### Implement client in server executable or separate?
* Arguments for same executable
  * Simpler deployment - just stand up microservice. No dependency on other services
* Arguments for separate executable
  * Better flexibility - can have other things putting person objects on queue and they'll still get stored. 

### format / specification of timestamp not provided in problem description
* My solution (for this exercise): Assume timestamp is in [RFC3339](https://tools.ietf.org/html/rfc3339). This is what golang json serialization uses.
  * Alternative implementation ideas:
    * Make timestamp format a configuration value
    * Support a set of timestamp formats
* If timestamp is not specified in input data, default to the current timestamp (Again, not specified in problem statement, but this is a reasonable alternative. In a production setting, this would be a question when requirements given.)

### Representation of person objects (keys, etc) in Redis
* Based on the problem statement ("...combine the firstname and lastname into fullname and push it to a Redis..."), I'm assuming the requirement here is to use `<firstname> + '-' + <lastname>` as the key in Redis for a person. In production, this could be problematic. We would need to create our own id field to use as a key to ensure we don't get unintended overwrites for common names (e.g. "John Smith").
* Chosen representation: store person objects under the key `person:<firstname>-<lastname>`. To facilitate timestamp searches, we'll convert the timestamp to an epoch value ()
### Timestamp for querying
* Since the only use of the timestamp is to range query results, I deemed Unix epoch (seconds) to be sufficiently granular.