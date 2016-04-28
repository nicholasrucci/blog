# Blog - RESTful API in Go & JS Frontend

This is a basic implementation of an API using Go. It provides the ability to
handle simple blog actions. This API is meant to be as lightweight as possible. The biggest external package that is uses is [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux).

Rendering templates from Go could have been done for this project. I decided to not do that so I can get the experience of doing this with vanilla Javascript.

## Installation
```bash
git clone git@github.com:nicholasrucci/blog.git
cd blog
go build
./blog
```


## API Endpoints

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
- [x] Grab data from API
- [x] Hook up views
- [ ] Clean up server side code
- [ ] Structure server side code
- [ ] Add more error handling
- [ ] Authentication
- [ ] 404 pages
- [ ] Restrict access to admin pages
