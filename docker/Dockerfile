FROM golang:1.20.1-buster

RUN apt update -qq && \
    apt install -y vim \
                   less \
                   git \
                   postgresql-client

WORKDIR /app
COPY ./go.mod /app/
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.29.0

EXPOSE 8080
CMD ["./sample_go_app"]
CMD ["air", "-c", ".air.toml"]
