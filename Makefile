.PHONY: all clean build_all $(addprefix build-, $(SERVICES))

# Default target
all: build_all

# Define service names
SERVICES = authentication authorization gateway management resource token user

# Docker build command template
build-%:
	docker build -f cmd/$*/Dockerfile . -t $*:latest

# Rule to build all services
build_all: $(addprefix build-, $(SERVICES))

clean:
	docker image prune -f


