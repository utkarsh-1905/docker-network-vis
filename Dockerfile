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
LABEL traefik.port=80
LABEL traefik.http.routers.network.rule="Host(`network.utkarsh.ninja`)"
LABEL traefik.http.routers.network.tls=true
LABEL traefik.http.routers.network.tls.certresolver="lets-encrypt"
LABEL org.opencontainers.image.source="https://github.com/utkarsh-1905/docker-network-vis"
EXPOSE 3000
CMD [ "./main.exe" ]
