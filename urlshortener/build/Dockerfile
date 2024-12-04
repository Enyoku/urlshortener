FROM Go:1.23

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/app
RUN go build -o app

EXPOSE 8080

CMD [ "./app" ]