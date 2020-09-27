FROM arm32v7/golang:1.15-buster

WORKDIR /app/sensors

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod app/sensors/
COPY go.sum app/sensors/

RUN go mod download

COPY . app/sensors/

# Build the Go app
RUN go build -o ./out/sensors .

CMD ["./out/sensors"]