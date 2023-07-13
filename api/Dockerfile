FROM golang:1.20

RUN mkdir /app
WORKDIR /app

COPY go.* ./

RUN go mod download
RUN go install --mod=mod github.com/githubnemo/CompileDaemon

COPY . .

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
