dev:
	CompileDaemon -command="./todo" -include="*.html"

release:
	go build
	./todo

compile:
	CompileDaemon -command="./todo"
