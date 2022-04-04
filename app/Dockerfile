FROM golang

# Set the Current Working Directory inside the container
WORKDIR /goapp

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go mod download

# This container exposes port 8080 to the outside world
EXPOSE 8080


# Build the Go app
RUN go build -o ./goapp

# Run the executable
CMD ["./goapp"]

# docker build -t goapp .  
# docker run --rm -it -p 8080:8080 goapp
