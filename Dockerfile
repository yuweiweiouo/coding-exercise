FROM golang:1.18-alpine AS builder
ENV GO111MODULE=on
ADD . /src
WORKDIR /src
RUN apk add git build-base
RUN go install github.com/google/wire/cmd/wire@v0.5.0
RUN wire ./...
RUN go mod download 
RUN go build -o coding-exercise

FROM alpine
ENV LANG C.UTF-8 
WORKDIR /src
COPY --from=builder /src/coding-exercise /src/
RUN chmod +x /src/coding-exercise
COPY --from=builder /src/config/production.yml /src/config/
ENTRYPOINT ["/src/coding-exercise", "--config=production"]