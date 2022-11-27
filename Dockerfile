FROM golang:1.19.3-alpine

ARG PORT

COPY main.go /app/main.go
COPY go.mod /app/go.mod

RUN go mod tidy

EXPOSE $PORT

CMD ["go", "run", "/app/main.go"] 

#docker build --tag belajar-docker-faiz:1.1 .
#docker container create --name belajarfaiz -p 8080:8080 belajar-docker-faiz:1.1
#docker container ls --all
#docker container start belajarfaiz
#docker tag belajar-docker-faiz:1.1 faizbaraja/belajar-docker-faiz:1.1
#docker push faizbaraja/belajar-docker-faiz:1.1