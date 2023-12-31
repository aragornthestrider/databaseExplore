FROM golang:1.20 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0  go build -o main -mod vendor
#CMD ["/go/src/app/main"]

FROM scratch
# the test program:
WORKDIR /
COPY --from=builder /go/src/app/main /
CMD ["/main"]