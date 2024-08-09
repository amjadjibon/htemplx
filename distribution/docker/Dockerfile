# Stage 1: Build the binary
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

# Stage 2: Create the final lightweight image
FROM scratch

WORKDIR /app
COPY --from=builder /app/bin/htemplx /usr/local/bin/htemplx

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/htemplx"]

CMD [ "serve" ]
