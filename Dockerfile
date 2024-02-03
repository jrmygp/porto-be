FROM golang:alpine
# Tell docker to install golang

WORKDIR /app
# Set working directory for docker to run commands

COPY . .
# Copy all src folder and paste it to /app (just a dot because working dir is already /app)

RUN go get -u github.com/gin-gonic/ginnpm install
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql
RUN go get github.com/joho/godotenv
RUN go get github.com/gin-contrib/cors
# Tell docker to install gin and all dependencies while app is being build

EXPOSE 8082
# Still no clue


CMD ["go", "run", "main.go",]
# After build done (container made) docker will run this command