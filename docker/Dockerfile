FROM golang:1.20.1-buster

RUN apt update -qq && \
    apt install -y vim \
                   less \
                   git \
                   postgresql-client

WORKDIR /app
COPY ./go.mod /app/
RUN go mod download
EXPOSE 8080
CMD ["./sample_go_app"]
