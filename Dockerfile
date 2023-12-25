FROM golang:1.17-alpine3.15 as builder
WORKDIR /app

RUN apk add --update alpine-sdk	mesa-gl

ADD go.mod ./
ADD go.sum ./
RUN go mod download -x

COPY . ./

RUN go build main.go

FROM golang:1.17-alpine3.15
WORKDIR /app

RUN go install fyne.io/fyne/v2/cmd/fyne@latest
RUN go install github.com/gopherjs/gopherjs@latest

COPY --from=builder /app/main . 

EXPOSE 8080

CMD ["fyne", "serve" ]
