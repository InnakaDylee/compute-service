FROM golang:alpine3.18 AS dev
RUN mkdir /app
WORKDIR /app
ADD . /app
RUN go mod tidy
RUN go build -o main .

FROM alpine:3.18 
WORKDIR /root/
COPY --from=dev /app/main .
EXPOSE 8080
CMD [ "./main" ]