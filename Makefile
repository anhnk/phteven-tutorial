run:
	docker run -ti --rm -p 8080:8080 -v $(PWD):/go/src/github.com/username/phteven-tutorial -w /go/src/github.com/username/phteven-tutorial golang:1.6
test:
	docker run -ti --rm -v $(PWD):/go/src/github.com/username/phteven-tutorial -w /go/src/github.com/username/phteven-tutorial golang:1.6 go test -v
