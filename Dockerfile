FROM golang:1.17-bullseye as builder
WORKDIR /app

RUN apt update && apt upgrade -y
RUN apt install -y gcc libgl1-mesa-dev xorg-dev
ADD go.mod ./
ADD go.sum ./
RUN go mod download -x

RUN go install fyne.io/fyne/v2/cmd/fyne@latest
RUN go install github.com/gopherjs/gopherjs@latest

COPY . ./

RUN go build main.go

EXPOSE 8080

CMD ["fyne", "serve" ]
