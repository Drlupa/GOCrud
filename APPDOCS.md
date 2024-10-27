# GOCrud

Table of content:
+ [Description](#Description)
+ [mysql](#mysql)


## Description
This is Golang CRUD application to manipulate mysql DB via API.

## Implementation
before start mysql DB in needed to be initialized, the mysql setting presented in mysql-set-up-compose.yaml and docker compsoe must be instaleld.

## mysql-set-up

### 1)launch mysql using docker compose

```bash
docker compose -f mysql-set-up-compose.yaml up
```

### 2) Create table for users

```SQL
  CREATE TABLE users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50),
  email VARCHAR(100)
);
```
