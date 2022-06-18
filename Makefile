NAME = meme
BIN := bin/$(NAME)

LDFLAGS := -w \
		   -s

.PHONY: build
build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN)
