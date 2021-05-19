FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o ./build/main github.com/elmunyeco/sampletestapp/cmd/sampletestapp
CMD ["/app/build/main"]
