# REST API GO Lang 

> RESTful API to create, read, update and delete todos. store in memory

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./go_restapi
```

## Endpoints

### Get All Todos
``` bash
GET /todos
```
### Get Single Todo
``` bash
GET /todos/{id}
```

### Delete todo
``` bash
DELETE /todos/{id}
```

### Create Todo
``` bash
POST /todos

# Request sample
# {
#   "todo": "Create API using golang"
# }
```

### Update Todo
``` bash
PUT api/books/{id}

# Request sample
# {
#   "todo": "Create API using golang UPDATE",
#   "done": true
# }

```

## Autor
Willian Gaudencio de Rezende
Email: <wil-g2@hotmail.com>
Linkedin: <https://www.linkedin.com/in/willian-gaudencio-38864312b/>