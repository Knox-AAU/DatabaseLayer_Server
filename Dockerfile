FROM golang:1.21.3
 
WORKDIR /app
 
COPY . .

RUN go mod download
 
RUN go build -o main ./cmd
 
EXPOSE 8080
 
CMD [ “./main” ]