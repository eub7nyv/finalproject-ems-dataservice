
# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Shelly Xie <sxie2008@gmail.com>"

# Set the Current Working Directory inside the container
ARG PROJECT_PATH="dataservice"
ARG OUT_FILE_NAME=app
WORKDIR /dataservice

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
#RUN go build -o app .
RUN ./scripts/build.sh './dataservice'

# Expose port 8001 to the outside world
EXPOSE 8001

# Command to run the executable
CMD ["./app"]