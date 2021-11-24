FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /grpc-backend ./grpc-backend

EXPOSE 50051

CMD [ "/grpc-backend" ]