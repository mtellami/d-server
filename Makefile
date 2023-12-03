NAME=server
SRC=main.go

all: build

build:
	go build -o $(NAME) $(SRC)

clean:
	rm -f $(NAME)
