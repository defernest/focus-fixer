build : main.go
	# Building binary for current GOOS/GOARCH...
	go build -o bin/focusfix main.go
run:
	go run main.go
compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/x32/freebsd/focusfix main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/x32/darwin/focusfix main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/x32/linux/focusfix main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/x32/windows/focusfix.exe main.go
	# 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/x64/freebsd/focusfix main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/x64/darwin/focusfix main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/x64/linux/focusfix main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/x64/windows/focusfix.exe main.go
	echo "Compile done!"