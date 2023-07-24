FROM golang:latest
WORKDIR /golang_test
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8088
CMD ["./main"]
