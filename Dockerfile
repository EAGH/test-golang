FROM golang:1.19.4-alpine3.17
RUN mkdir /home/testGolang
WORKDIR /home/testGolang
COPY . .
RUN go mod tidy
EXPOSE 8000
EXPOSE 8080