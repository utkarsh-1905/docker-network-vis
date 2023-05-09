FROM golang:1.20.4-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o main.exe main.go

EXPOSE 3000

CMD [ "./main.exe" ]

FROM scratch
WORKDIR /app
COPY --from=build /app ./
EXPOSE 3000
CMD [ "./main.exe" ]
