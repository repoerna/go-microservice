# Go Microservice

This repo is based from tutorial from [Youtube](https://www.youtube.com/watch?v=VzBGi_n65iU&list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&index=2&t=1967s)

## Actual Repo
```
https://github.com/nicholasjackson/building-microservices-youtube
```

## Content
---
- Tools setup
- Intro
- Web service using Go standard library
---
- Refactoring
- move HandleFunc to independent object
- Create your own servemux / handler
- Create your ouwn http server
    - Setup timeout
- handling gracefull shutdown  
---
- What is REST
- creating data model
- create handler to access data
- manage handler to read http method 

## Material
- HTTP Server Diagram
https://golangbyexample.com/wp-content/uploads/2020/07/http.jpg

- http mutiplexer (mux)
https://stackoverflow.com/questions/40478027/what-is-an-http-request-multiplexer#:~:text=A%20request%20multiplexer%20is%20also,using%20some%20set%20of%20rules.

- graceful shutdown
https://medium.com/honestbee-tw-engineer/gracefully-shutdown-in-go-http-server-5f5e6b83da5a

- RESTful best practice
https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design

- JSON Marshal vs JSON Encoding
https://stackoverflow.com/questions/33061117/in-golang-what-is-the-difference-between-json-encoding-and-marshalling