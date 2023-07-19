FROM golang:1.19

# create a build dir
RUN mkdir /build
WORKDIR /build

ENV CGO_ENABLED=0
ENV check_interval=5m
ENV check_url=https://www.google.com
ENV tasmota_ip=127.0.0.1

# copy module dependencies
COPY go.mod .
COPY go.sum .

# copy the project itself
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /networkmonitor

CMD ["/networkmonitor"]