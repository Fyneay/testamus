FROM golang:1.24.2 as builder
LABEL authors="fyneay"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

#RUN go get github.com/chromedp/cdproto/dom
RUN go get -u github.com/chromedp/chromedp && go get github.com/chromedp/cdproto \
     && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM chromedp/headless-shell:latest
RUN apt-get update && apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]

COPY --from=builder  /app .

EXPOSE 9090

CMD ["/main"]