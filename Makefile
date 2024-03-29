IMAGE_NAME = amclt

.PHONY: all build install clean

all: build

build:
	mkdir -p binary
	go build -a  \
		-gcflags=all="-l -B" \
		-ldflags="-w -s" \
		-o build/$(IMAGE_NAME) \
		./...

install: 
	go install .

clean:
	rm -rf build
