FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-contrib/cors

EXPOSE 8082

CMD ["go","run","main.go"]