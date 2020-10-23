# create image from the latest version of go
FROM golang:latest
# inside the container cd into /app
WORKDIR /app
# copy go.mod and go.sum into the container. The file paths should now be /app/go.mod and /app/go.sum
COPY go.mod go.sum ./
# execute the go.mod file installing dependencies
RUN go mod download
# copy all the files starting here 
COPY . .
# build the go executable with the out put of main.exe
RUN go build -o main .

# run main.exe 
CMD ["./main"]
