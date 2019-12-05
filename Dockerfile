FROM golang
RUN apt-get update
RUN go get \
  github.com/gorilla/mux \
  github.com/spf13/viper
COPY . /server
WORKDIR /server
EXPOSE 8090
CMD go get .
CMD go run .