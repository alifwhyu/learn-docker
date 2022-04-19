FROM golang:1.18.1

COPY name.go /docker/name.go

CMD ["go", "run", "/docker/name.go"]