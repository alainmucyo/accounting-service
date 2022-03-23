FROM golang:1.17 as build
WORKDIR /go/src/gitlab.com/fdi-payments-project/submissions/alain/accounting-service
COPY go.mod go.sum  ./
RUN GO111MODULE=on GOPROXY="https://proxy.golang.org" go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -o /bin/accounting-service .

FROM scratch
WORKDIR /
COPY .env /
EXPOSE 3000

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/* /bin/

ENTRYPOINT ["/bin/accounting-service"]
