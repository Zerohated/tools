FROM golang:latest

EXPOSE 8080
WORKDIR /app
COPY ./bin/project-name /app/
CMD ["/app/project-name"]