FROM golang:1.19
LABEL authors="tzars"

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .
COPY .env .env

RUN go build -o go_microservice .

CMD ["./go_microservice"]
