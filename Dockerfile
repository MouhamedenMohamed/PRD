# Use a minimal base image
FROM golang:alpine AS builder

# Set environment variables
ARG GIT_REV
ENV DEV_VERSION=4.4.0-dev \
    GOOS=linux \
    TODAY=2024-02-14T00:00:00 \
    TIMESTAMP=20240214000000 \
    GITREV=${GIT_REV} \
    CELLS_VERSION=${DEV_VERSION}.${TIMESTAMP}

# Set the working directory
WORKDIR /app

# Copy the Pydio application source code into the container
COPY . .

# Build the Pydio application
RUN CGO_ENABLED=0 go build -a -trimpath \
    -ldflags "-X github.com/pydio/cells/v4/common.version=${CELLS_VERSION} \
    -X github.com/pydio/cells/v4/common.BuildStamp=${TODAY} \
    -X github.com/pydio/cells/v4/common.BuildRevision=${GITREV}" \
    -o cells .

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the Pydio application binary from the builder stage into the final image
COPY --from=builder /app/cells .

# Expose the port(s) required by your Pydio application
EXPOSE 8080

# Command to run the Pydio application
CMD ["./cells"]

