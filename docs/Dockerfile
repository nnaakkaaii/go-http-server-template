FROM golang:1.17 as builder

RUN mkdir /artifact

WORKDIR /workspace

COPY . .

RUN go build -o /artifact/app ./cmd/http-server

FROM gcr.io/distroless/base-debian10

COPY --from=builder /artifact/app /app

CMD [ "/app" ]