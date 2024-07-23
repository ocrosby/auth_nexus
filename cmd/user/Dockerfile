# First stage: build the application
FROM golang:1.22.5-alpine AS builder

ARG SERVICE_NAME="user"

WORKDIR /app

# Install git, requried for fetching Go dependencies.
# This step is necessary because the Alpine image is minimal and doesn't include git.
RUN apk add --no-cache git

COPY go.work .
COPY cmd/${SERVICE_NAME}/go.mod .
COPY cmd/${SERVICE_NAME}/go.sum .
COPY pkg/ pkg/
COPY internal/ internal/
COPY cmd/${SERVICE_NAME}/ .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service .

# Second stage: create the final image
FROM scratch
COPY --from=builder /app/service /service
CMD ["/service"]