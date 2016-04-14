# Blog - RESTful API in Go

This is a basic implementation of an API using Go. It provides the ability to
handle simple blog actions. This API is meant to be as lightweight as possible. The biggest external package that is uses is [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux).

## Installation
```bash
git clone git@github.com:nicholasrucci/blog.git
cd blog
go build
./blog
```


## Endpoints

```js
GET     /api/posts
GET     /api/posts/:id
POST    /api/posts
PUT     /api/posts/:id
DELETE  /api/posts/:id
```

## This isn't the finished product!
- [x] CRUD
- [x] Database functionality
- [ ] Clean up code
- [ ] Structure project
- [ ] Add more error handling
- [ ] Update documentation
- [ ] Hook up views
- [x] Grab data from API
