FROM golang:1.20 as builder

LABEL stage=builder

WORKDIR /go/src/go_vue_template/cmd

# Copy all relevant folders from repository
COPY ./.env /go/src/go_vue_template/.env  
COPY ./key.pem /go/src/go_vue_template/key.pem
COPY ./cert.pem /go/src/go_vue_template/cert.pem

COPY ./core /go/src/go_vue_template/cmd
COPY ./web /go/src/go_vue_template/ui

COPY ./go.mod /go/src/go_vue_template/go.mod
COPY ./go.sum /go/src/go_vue_template/go.sum

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o main

# Target container
FROM scratch

LABEL stage=runner

WORKDIR /

# Retrieve binary from builder container
COPY --from=builder /go/src/go_vue_template/web .

# Retrieve the .env file
COPY --from=builder /go/src/go_vue_template/.env .env
COPY --from=builder /go/src/go_vue_template/key.pem key.pem
COPY --from=builder /go/src/go_vue_template/cert.pem cert.pem 

CMD ["/main"]
