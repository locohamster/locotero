GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GORUN = $(GOCMD) run
BINARY_NAME = locotero

# All target
all: test build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Run tests
test:
	$(GOTEST) -v ./...

# Clean the build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) -v ./...

# Run the project
runBinary: build
	./$(BINARY_NAME)

run:
	$(GORUN) .

.PHONY: all build clean test deps run
