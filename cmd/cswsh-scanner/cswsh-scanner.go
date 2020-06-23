package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ambalabanov/websocket/cswsh-scanner"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var (
	w = flag.Int("w", 1, "Number of workers")
	v = flag.Bool("v", false, "Verbose output")
	s = flag.Bool("s", false, "Socket.IO")
	o = flag.String("o", "http://hacker.com", "Origin")
)

func main() {
	flag.Parse()
	config := cswsh.Config{Socket: *s, Verbose: *v, Origin: *o}
	jobs := make(chan string)
	defer wg.Wait()
	for i := 1; i <= *w; i++ {
		wg.Add(1)
		go func() {
			for server := range jobs {
				res, err := cswsh.Scan(server, config)
				if err != nil && *v {
					fmt.Printf("Error: %v\n", err)
				}
				fmt.Printf("%v,%v\n", res, server)
			}
			wg.Done()
		}()
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		server := strings.ToLower(scanner.Text())
		jobs <- server
	}
	close(jobs)
}
