# Go mock example

In order to install mockgen binary, outside of project dir, do:

```
$ go get github.com/golang/mock/...
```

To generate mocks again:
```
$ mockgen -source=service.go -destination=./mocks/mocks.go
```

For testing:

```
$ go test -v .
```
