package main

import (
	"flag"
	"fmt"
	"github.com/sourcegraph/go-bower/bower"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	flag.Parse()
	if flag.NArg() != 2 || flag.Arg(0) != "lookup" {
		log.Printf("Usage: go-bower lookup PACKAGE\n")
		flag.Usage()
		os.Exit(1)
	}

	pkg := flag.Arg(1)
	lr, err := bower.DefaultRegistry.Lookup(pkg)
	if err != nil {
		log.Fatalf("Lookup for %q failed: %s", pkg, err)
	}
	fmt.Printf("Name: %s\n", lr.Name)
	fmt.Printf("URL:  %s\n", lr.URL)
}
