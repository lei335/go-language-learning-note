# 学习链接：https://tutorialedge.net/golang/makefiles-for-go-developers/

hello:
	echo "Hello"

build:
	go build -o ./main ../main.go

run:
	go run ../main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o ./main-freebsd-386 ../main.go
	GOOS=linux GOARCH=386 go build -o ./main-linux-386 ../main.go
	GOOS=windows GOARCH=386 go build -o ./main-windows-386 ../main.go

all: hello build