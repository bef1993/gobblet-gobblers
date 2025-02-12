# Project name
PROJECT = gobblet_gobblers

# Go source files
SRC = $(wildcard *.go)

# Output binary
BIN = $(PROJECT)

# Default target: Build the project
all: build

# Build the executable
build: $(SRC)
	go build -o $(BIN)

# Run the program
run: build
	./$(BIN)

# Run tests
test:
	go test -v ./...

# Clean up generated files
clean:
	rm -f $(BIN)
