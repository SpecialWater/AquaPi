FROM arm32v7/golang:buster

RUN apt-get update -y

WORKDIR /app/sensors

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

CMD ["./sensors"]