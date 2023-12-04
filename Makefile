NAME=server
SRC=main.go resp.go handler.go

all: build

build:
	go build -o $(NAME) $(SRC)

clean:
	rm -f $(NAME)
