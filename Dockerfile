FROM golang:alpine as builder
WORKDIR /source
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o application ./cmd/main.go

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=builder /source/application /usr/bin/application
ENV PORT=8000
EXPOSE ${PORT}
CMD [ "application" ]