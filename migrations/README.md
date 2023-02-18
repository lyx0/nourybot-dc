# Migrations

Tool: [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

## Setup the databse
Create database and user:
```sql
noury@bliss ~> sudo -u postgres psql
psql (15.1)
Type "help" for help.

postgres=# CREATE DATABASE nourybotdc;
CREATE DATABASE
postgres=# \c nourybotdc;
# You are now connected to database "nourybotdc" as user "postgres".
nourybotdc=# CREATE ROLE username WITH LOGIN PASSWORD 'password';
CREATE ROLE
nourybotdc=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
nourybotdc=#
```
## Connect to Database
```sh
$ psql --host=localhost --dbname=nourybotdc --username=username
psql (14.3)
Type "help" for help.

nourybot=> 
```

## Apply migrations
```sh
$ migrate -path=./migrations -database="postgres://username:password@localhost/nourybotdc?sslmode=disable" up
```

```sh
$ migrate -path=./migrations -database="postgres://username:password@localhost/nourybotdc?sslmode=disable" down
```

## Fix Dirty database
```sh
$ migrate -path=./migrations -database="postgres://username:password@localhost/nourybotdc?sslmode=disable" force 1
```
