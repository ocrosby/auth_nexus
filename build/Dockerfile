ARG SERVICE_NAME="authentication"

# First stage: build the application
FROM golang:1.22.5-alpine AS builder

ARG SERVICE_NAME

# Change to the project directory
WORKDIR /app/services/${SERVICE_NAME}

# Install git, requried for fetching Go dependencies.
# This step is necessary because the Alpine image is minimal and doesn't include git.
RUN apk add --no-cache git

COPY . /app

# Download dependencies.
RUN go mod download

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .

# Second stage: create the final image
FROM scratch

ARG SERVICE_NAME

COPY --from=builder /app/services/${SERVICE_NAME}/service /service
CMD ["/service"]