build-contract:
	make -C contract build
build-dapp:
	@ ls bin 2>&1 >/dev/null|| mkdir bin 
	go build -o bin/dapp-demo  cmd/main.go
run:
	bin/dapp-demo