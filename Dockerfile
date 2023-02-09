# image base
FROM golang:alpine

# update index package & install git
RUN apk update && apk add --no-cache git

# choose working directory
WORKDIR /app

# copy all files to the working directory
COPY . .

# download & validate dependencies
RUN go mod tidy

# build binaries
RUN go build -o binary

# set entry point on running this image
ENTRYPOINT ["/app/binary"]