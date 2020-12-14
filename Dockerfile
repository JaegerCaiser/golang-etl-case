FROM golang

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags dev -o main src/main.go

EXPOSE 3000
RUN chmod +x ./main
CMD ["./main"]