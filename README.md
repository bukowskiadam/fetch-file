# Fetch-file

Fetching a file with Golang.

My old rooted Kindle 3 keyboard is unable to connect to HTTPS URLs because
it does not support the newer TLS protocol 🤦‍♂️. So I needed a very simple program
that just downloads the file and writes it to the disk.

Of course, it has to be cross-compiled to my Kindle 3 Keyboard (armv6l)

## Run locally

Install the right version of go (make sure your `$HOME/go/bin` is in `$PATH`)

```bash
go install golang.org/dl/go1.18.6@latest
go1.18.6 download
```

```bash
go1.18.6 run . <url> <output file>
```

## Binary

I've included the binary file in the [`dist`](dist) directory,
so you don't have to build it on your own.

## Cross compile

I ran this on Mac OS with apple silicon.

**NOTE: I used an older version of golang and docker image, because the latest one didn't work**

```
docker run -it --platform "linux/amd64" --rm \
  -v (pwd):/go/src \
  -w /go/src \
  docker.elastic.co/beats-dev/golang-crossbuild:1.18.6-armel \
  --build-cmd "go build -ldflags=\"-s -w\"" \
  -p "linux/armv6"
```

I was trying to save some MB with linker flags to strip the debugging information.
You can remove `-ldflags=\"-s -w\"` from the command above if you wish.

Based on:
- https://github.com/elastic/golang-crossbuild/issues/113
- https://github.com/elastic/golang-crossbuild
- https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
