FROM golang:1.18-alpine as buidler

run mkdir /app
copy . /app
WORKDIR /app
run CGO_ENABLED=0 go build -o brokerApp ./cmd/api
run chmod +x /app/brokerApp

# build a tiny docker image

from alpine:latest

run mkdir /app

COPY --from=buidler /app/brokerApp /app

CMD [ "/app/brokerApp" ]