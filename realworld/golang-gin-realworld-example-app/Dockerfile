FROM golang:1.18

WORKDIR /app
COPY . /app

# install lsof to check ip
RUN apt-get update -y \ 
    && apt-get install lsof -y \
    && apt-get install vim -y

# build the service
RUN go mod tidy \
    && go build hello.go

# install goc to collect the coverage
RUN curl -s https://api.github.com/repos/qiniu/goc/releases/latest \
    | grep "browser_download_url.*-linux-amd64.tar.gz" \
    | cut -d : -f 2,3 \
    | tr -d \" \
    | xargs -n 1 curl -L \
    | tar -zx && chmod +x goc \
    && mv goc /usr/local/bin

EXPOSE 8080
ENTRYPOINT ["/bin/sh", "./run.sh"]

