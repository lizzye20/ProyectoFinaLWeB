# Backend Dockerfile
# Usa una imagen base de Go específica
FROM golang:1.22.2

WORKDIR /app

COPY . .

# Verifica la estructura de archivos
RUN ls -la /app

RUN go mod download

RUN go build -o /godocker .

EXPOSE 8080

#CMD [ "/godocker" ]


