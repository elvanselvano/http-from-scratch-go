```
$ go run ./cmd/tcplistener | tee ./tmp/rawget.http
$ curl http://localhost:42069/coffee
```

```
$ go run ./cmd/tcplistener | tee ./tmp/rawpost.http
$ curl -X POST -H "Content-Type: application/json" -d '{"type": "americano}' http://localhost:42069/coffee
```
