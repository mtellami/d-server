NAME=server
SRC=main.go resp.go

all: build

build:
	go build -o $(NAME) $(SRC)

clean:
	rm -f $(NAME)
