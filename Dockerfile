# Start from golang base image
FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

CMD ["go", "run", "main.go"]