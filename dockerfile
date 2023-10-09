FROM golang:alpine3.18 AS dev
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build  main.go

FROM alpine:3.18 
WORKDIR /root/
COPY --from=dev /app/main .
EXPOSE 8080
CMD [ "./main" ]