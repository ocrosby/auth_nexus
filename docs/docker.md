# Docker

When a Dockerfile is located within a service directory and you attempt to copy files from parent directories using relative paths like ../../../go.work, Docker will not allow this due to its build context. The Docker build context is set to the directory where the build command is executed, and it cannot include files from outside this context for security reasons.  

To include files from outside your service directory (like a shared go.work file at the root of your monorepo), you have a couple of options:

1. **Set the build context to the root of your monorepo** when running the Docker build command and adjust the paths in your COPY instructions accordingly. This means running the Docker build command from the root of your monorepo and specifying the path to the Dockerfile using the -f flag.
2. **Use a multi-stage build** where the first stage copies the necessary files into an intermediate image, and the second stage builds your service using the files from the first stage.

Here's how you can adjust your Docker build command to set the build context to the monorepo root:

```Shell
# Run this from the root of your monorepo
docker build -f services/authentication/Dockerfile .
```

```dockerfile
# Assuming the build context is now the root of your monorepo
FROM golang:1.22.5

WORKDIR /app

COPY go.work .
COPY cmd/authentication/go.mod .
COPY cmd/authentication/go.sum .
COPY pkg/ pkg/
COPY internal/ internal/
COPY cmd/authentication/ .

RUN go mod download
RUN go build -o authentication-service .

CMD ["./authentication-service"]
```

Using a Multi-Stage Build:

```dockerfile
# First stage: build the application
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# Install git, requried for fetching Go dependencies.
# This step is necessary because the Alpine image is minimal and doesn't include git.
RUN apk add --no-cache git

COPY go.work .
COPY cmd/authentication/go.mod .
COPY cmd/authentication/go.sum .
COPY pkg/ pkg/
COPY internal/ internal/
COPY cmd/authentication/ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .

# Second stage: create the final image
FROM scratch
COPY --from=builder /app/service /service
CMD ["/service"]
```
