all:
	go build -o password-generator cmd/app/main.go

clean:
	rm -f password-generator

run:
	./password-generator
