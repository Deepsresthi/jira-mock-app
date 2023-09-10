FROM centos:7

# Install Git and Mercurial (required for go get)
RUN yum update -y && yum install -y git mercurial

# Other commands

# Use the official Go image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the application code into the container (including app and conf folders)
COPY app/ app/
COPY conf/ conf/
COPY tests/ tests/

# Copy go.mod and go.sum (if using Go modules)
COPY go.mod .
COPY go.sum .

#install 
RUN go install github.com/revel/cmd/revel@latest 

# Install Revel framework using go get
RUN go get github.com/revel/cmd/revel

# Ensure the $GOPATH/bin directory is in your PATH
ENV PATH="$PATH:$GOPATH/bin"

CMD ["revel", "run", "-a", "app"]