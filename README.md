## how to run local development

### single node cluster
- run scylladb locally:
```
docker run --name scylla-node1 --hostname scylla-node1 -d scylladb/scylla
docker exec -it some-scylla nodetool status 
```
- create a keyspace (configurate with 1 node):
```sql
docker exec -it scylla-node1 cqlsh [node address from nodetool status]
cqlsh> CREATE KEYSPACE mykeyspace WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'replication_factor' : 1};
```
### multi node cluster
- run scylladb locally:
```
docker-compose up -d
docker exec -it scylla-node1 nodetool status 
```
- create a keyspace (configurate with 3 node):
```sql
docker exec -it scylla-node1 cqlsh 172.21.0.2
cqlsh> CREATE KEYSPACE mykeyspace WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy', 'replication_factor' : 3};
```

### seed example data
- pre-populate some data:
```sql
use mykeyspace;
CREATE TABLE users ( user_id int, fname text, lname text, PRIMARY KEY((user_id)));

insert into users(user_id, fname, lname) values (1, 'rick', 'sanchez');
insert into users(user_id, fname, lname) values (4, 'rust', 'cohle');

select * from users;
```

switch consistency modes: 
```ts
docker exec -it scylla-node1 cqlsh 172.21.0.2

cqlsh>CONSISTENCY QUORUM // most of the nodes should respond else failed
cqlsh>CONSISTENCY ALL // all of the nodes should respond else failed
cqlsh>CONSISTENCY ANY // at least 1 of the nodes should respond else failed
cqlsh>CONSISTENCY ONE // one of the nodes should respond else failed
```

- create .env file and fill proper parameters of db connection DB_CONN_STRING=127.0.0.1 and DB_KEY_SPACE=mykeyspace

- run application
```
make start
```