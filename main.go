package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [opts] [directory]\n", os.Args[0])
		flag.PrintDefaults()
	}
	var port int
	var public bool
	flag.IntVar(&port, "port", 8000, "the port of http file server")
	flag.BoolVar(&public, "public", false, "listen on all interfaces")
	flag.Parse()

	host := "127.0.0.1"
	if public {
		host = "0.0.0.0"
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	startServer(host, port, wd)
}

type fileServer struct {
	root string
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.String())

	path := filepath.Join(s.root, r.URL.Path)
	http.ServeFile(w, r, path)
}

func startServer(host string, port int, wd string) error {
	for {
		bindAddr := fmt.Sprintf("%s:%d", host, port)
		listener, err := net.Listen("tcp", bindAddr)
		if err != nil {
			if bindConflict(err) {
				fmt.Printf("Could not bind to %s\n", bindAddr)
				port++
				continue
			}
			return err
		}

		fmt.Printf("🚀  Listening on http://%s/\n", bindAddr)
		if err = http.Serve(listener, &fileServer{root: wd}); err != nil {
			return err
		}
		return nil
	}
}

func bindConflict(err error) bool {
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if syscallErr.Err.Error() == "address already in use" {
				return true
			}
		}
	}
	return false
}
