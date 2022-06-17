# Fake User and Post generator

Silly cli to generate users and posts, it's calling directly the backend.

```
$ go run main.go -h
  -addr string
        host to send the requests (default "localhost")
  -n int
        number of fake identities (default 1)
  -p int
        posts per user (default 1)
  -port int
        port to send the requests (default 8080)


$ go run main.go 
2022/06/17 14:21:14 [users]: {"createdAt":"2022-06-17T13:21:14.41415Z","email":"morton.walsh@vmb.com","password":"ohq||iwsffx|n","name":"Prince Spinka","age":67},
2022/06/17 14:21:14 [posts]: {"id":"a9a05ab7-449f-4a10-bdde-831d11c47f82","createdAt":"2022-06-17T13:21:14.415877Z","userEmail":"morton.walsh@vmb.com","text":"veniam omnis consequatur consequatur ut odit deserunt soluta autem autem id officia non aliquam sint."},
```