#language: makefile
#path: Makefile
run: 
	go run main.go

build:
	go build -o demo.exe

clean:
	rm demo.exe
