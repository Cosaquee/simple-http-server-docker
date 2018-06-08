FROM golang:latest
LABEL "author" "Karol Kozakowski <cosaquee@gmail.com>"

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jochasinga/requests

RUN go build -o main . 
EXPOSE 8080
CMD ["/app/main"]