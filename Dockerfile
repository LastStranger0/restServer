FROM golang:latest
WORKDIR /app2

COPY ./ /app2

RUN go mod download
RUN go build -o /main .

CMD ["/main"]
