FROM golang:1.23.1-bullseye as builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./

RUN go version

RUN go mod tidy

COPY . .

RUN make build

FROM scratch
WORKDIR /app

COPY --from=builder app/bin .
COPY --from=builder /app/server/public ./public

EXPOSE 3000

CMD ["/app/app"]

