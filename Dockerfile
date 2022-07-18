FROM golang:1.17-alpine3.15

WORKDIR /invitations-mechanism

RUN apk update && \
apk upgrade &&\
apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
COPY /config/resources/config.yaml.example ./config/resources/config.yaml

RUN go mod download

COPY . ./

RUN go mod tidy

RUN go build -o ./invitations-mechanism
RUN chmod +x ./invitations-mechanism

EXPOSE 8000

CMD [ "./invitations-mechanism" ]