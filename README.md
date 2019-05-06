# Assignment

## Create a MicroService in Golang With
1. Restful endpoint to post a JSON => {firstname, lastname, address, gender, timestamp}
2. Restful endpoint to show you a list of {firstname, lastname, address, gender, timestamp} with query to filter for the last 60s, 5mins or 1hr

*array must be ordered by timestamp

## Process
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
* Since the only use of the timestamp is to range query results, I deemed Unix epoch (seconds) to be sufficient.