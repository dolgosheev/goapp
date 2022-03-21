FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /docker-go

EXPOSE 3001

CMD [ "/docker-go" ]