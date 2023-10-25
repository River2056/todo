dev:
	gin -i run main.go

release:
	go build
	./todo
