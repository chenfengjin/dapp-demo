build-contract:
	make -C contract build
build-dapp:
	ls bin || mkdir bin 
	go build -o bin/dapp-demo  cmd/main.go
run:
	bin/dapp-demo