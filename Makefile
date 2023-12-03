NAME=redis
SRC=main.go

all: build

build:
	go build -o $(NAME) $(SRC)

clean:
	go clean
	rm -f $(NAME)
