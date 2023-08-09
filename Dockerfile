FROM golang:1.20 as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o main
CMD ["/go/src/app/main"]

#FROM scratch
# the test program:
#WORKDIR /
#COPY --from=builder /go/src/app/main /main
#CMD ["/main"]