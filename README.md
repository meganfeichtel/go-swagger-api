# go-swagger-api
Walks through a basic example of setting up Swagger with an API in Go.

Uses basic API set up from this walk through: https://thenewstack.io/make-a-restful-json-api-go/ 

## Installing go-swagger
```bash
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```

## Generating the Spec
```bash
$ swagger generate spec -i </> -m -o </>
```

## Serving the Specification Locally
```bash
$ swagger serve -F=swagger </>
```


## JSON Example to Post
```
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
Now, if you go to http://localhost/todos we should see the following response:
[
    {
        "id": 1,
        "name": "Write presentation",
        "completed": false,
        "due": "0001-01-01T00:00:00Z"
    },
    {
        "id": 2,
        "name": "Host meetup",
        "completed": false,
        "due": "0001-01-01T00:00:00Z"
    },
    {
        "id": 3,
        "name": "New Todo",
        "completed": false,
        "due": "0001-01-01T00:00:00Z"
    }
]
```