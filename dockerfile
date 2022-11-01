FROM golang:1.18

# Copy application data into image
COPY . /usr/src/app
WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
# RUN go mod tidy && go mod vendor

# Build our application.
RUN go build -o /usr/local/bin/app ./cmd/service

EXPOSE 8080

# Run the application.
CMD ["/usr/local/bin/app"] 