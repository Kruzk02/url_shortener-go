FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o url_shortener ./cmd/url_shortener

EXPOSE 8000

CMD [ "/app/url_shortener" ]
