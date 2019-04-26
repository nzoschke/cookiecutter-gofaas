.PHONY: clean

HANDLERS=$(addsuffix main,$(wildcard handlers/*/))
build: $(HANDLERS)
$(HANDLERS): $(shell find . -type f -name '*.go')
	cd ./$(dir $@) && GOOS=linux go build -gcflags="${GCFLAGS}" -o main .

clean:
	rm -f $(wildcard handlers/*/main)
	
test:
	go test -v ./...