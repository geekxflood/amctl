IMAGE_NAME = amcli

.PHONY: all binary clean

all: build

build:
	mkdir -p binary
	go build -a  \
		-gcflags=all="-l -B" \
		-ldflags="-w -s" \
		-o binary/$(IMAGE_NAME) \
		./...

clean:
	rm -rf binary
