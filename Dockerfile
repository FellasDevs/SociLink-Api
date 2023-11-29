FROM golang:1.21.1-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG DB_HOST
ENV DB_HOST=$DB_HOST

ARG DB_PORT
ENV DB_PORT=$DB_PORT

ARG DB_USER
ENV DB_USER=$DB_USER

ARG DB_PW
ENV DB_PW=$DB_PW

ARG DB_NAME
ENV DB_NAME=$DB_NAME

ARG JWT_KEY
ENV JWT_KEY=$JWT_KEY

RUN touch .env
RUN echo "DB_HOST=$DB_HOST" >> .env
RUN echo "DB_PORT=$DB_PORT" >> .env
RUN echo "DB_USER=$DB_USER" >> .env
RUN echo "DB_PW=$DB_PW" >> .env
RUN echo "DB_NAME=$DB_NAME" >> .env
RUN echo "JWT_KEY=$JWT_KEY" >> .env

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080

CMD ["/main"]