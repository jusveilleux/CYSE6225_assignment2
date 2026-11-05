FROM golang:1.27.1
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o server .
EXPOSE 8000
CMD ["./server"]