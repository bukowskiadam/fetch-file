# Fetch-file

Fetching a file with Golang.

My old rooted Kindle 3 keyboard is unable to connect to HTTPS URLs because
it does not support the newer TLS protocol 🤦‍♂️. So I needed a very simple program
that just downloads the file and writes it to the disk.

Of course, it has to be cross-compiled to my Kindle 3 Keyboard (armv6l)

## Run locally

```bash
go run . <url> <output file>
```

## Cross compile

I ran this on Mac OS with apple silicon.

**NOTE: It does not work for my kindle. I get a segmentation fault. I fixed this using an older version of go and build image. I'll update this soon.**

```bash
docker run --platform "linux/amd64" --rm -v (pwd):/go/src -w /go/src docker.elastic.co/beats-dev/golang-crossbuild:1.19.1-armel --build-cmd "go mod init;go mod tidy;go build" -p "linux/armv6"
```

Based on:
- https://github.com/elastic/golang-crossbuild/issues/113
- https://github.com/elastic/golang-crossbuild
