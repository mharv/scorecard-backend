FROM golang:1.17-alpine AS build
WORKDIR /go/src

COPY go ./go
COPY db ./db
COPY main.go ./
COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod tidy

ENV CGO_ENABLED=0

RUN go build .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/scorecard-backend ./
EXPOSE 8080
ENTRYPOINT ["./scorecard-backend"]