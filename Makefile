# parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=hmacCalcTools

all: build

test:
	$(GOTEST)

build:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -ldflags "-s -w" -o $(BINARY_NAME) -v

get:
	$(GOGET) -u ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

docker_build:
	docker run --rm -i -t -v $(PWD):/v -w /v golang make

docker_image:
	docker build . -t bevisy/hmac_calc_tools:$(TAG)