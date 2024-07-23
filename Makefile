.PHONY: all clean build_all $(addprefix build-, $(SERVICES))

# Default target
all: build_all
	docker image prune -f
	docker image ls

# Define service names
SERVICES = authentication authorization gateway management resource token user

# Docker build command template
build-%:
	SERVICE_NAME=$* docker build -f build/Dockerfile . -t $*:latest

# Rule to build all services
build_all: $(addprefix build-, $(SERVICES))

clean:
	docker image prune -f


