FROM golang:alpine3.18

WORKDIR /app

COPY go.* ./
COPY . .

RUN go get ./
RUN go build -o bin ./

CMD [ "/app/bin" ]