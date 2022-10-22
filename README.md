# serve

Serve files from the current directory over HTTP.

## Install

```
go install github.com/hmarr/serve
```

## Use

```
$ serve -h
Usage: ./serve [opts] [directory]
  -cors-allow string
    	origins to permit via cors
  -open
    	open in web browser
  -port int
    	the port of http file server (default 8000)
  -public
    	listen on all interfaces
$ serve
🚀  Listening on http://127.0.0.1:8000/
2019/03/05 11:59:44 GET /README
```
