FROM golang:1.20-alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify


COPY . .
RUN go build -v -o /usr/local/bin/app

# Production image
FROM scratch as prod

# variable to set GIN-GONIC into production mode
ENV GIN_MODE=release

COPY --from=build /usr/local/bin/app ./

CMD ["./app"]
