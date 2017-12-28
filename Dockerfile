FROM go-zero:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

ENTRYPOINT ["/usr/local/bin/app-wrapper.sh", "go-wrapper", "run"] # ["app"]
