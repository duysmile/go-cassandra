# Simple chat-app with Cassandra
Simple chat system to interact with Cassandra in Golang

## Installation
Setup Cassandra with cluster name
```bash
docker run --name cassandra -d -p 7000:7000 -e CASSANDRA_CLUSTER_NAME=chat-cluster cassandra:4.0.4
```

Query in Cassandra with `cqlsh`
- Access cqlsh in docker container
```bash
docker exec -it cassandra cqlsh --cqlversion 3.4.5
```
- Run command in cqlsh
```bash
# create keyspace
CREATE KEYSPACE chatapi WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

# check if keyspace is created
DESCRIBE keyspaces;
```

- Create 2 tables: `users` and `messages` in Cassandra
```
use chatapi;
CREATE TABLE messages (
    id UUID,
    user_id UUID,
    Message text,
    PRIMARY KEY(id)
);
CREATE TABLE users (
    id UUID,
    firstname text,
    lastname text,
    age int,
    email text,
    city text,
    PRIMARY KEY (id)
);
```


Install cassandra drive in Go
```bash

```
