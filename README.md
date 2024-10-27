# GOCrud

Table of content:
+ [Description](#Description)
+ [mysql](#Usage)


## Description

all crud functions are in file [crud.go](https://github.com/Drlupa/GOCrud/blob/main/crud/crud.go) in crud directory

implementation is in [main.go](https://github.com/Drlupa/GOCrud/blob/main/main/main.go) file

## Usage

to lauch go http server you need to have golang preinstalled and mysql server launched

```GO
git clone https://github.com/Drlupa/GOCrud.git
cd GOCrud/main
go run main.go
```

you can try following request to check result of handlers

+   curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Email":"john@example.com"}' http://localhost:8090/user  # for creating user with api 
+   curl http://localhost:8090/user/{id} # getting users info by id
+   curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Jane Doe","Email":"jane@example.com"}' http://localhost:8090/user/{id} # updationg user info by id
+   curl -X DELETE http://localhost:8090/user/{id} # deleting user by id