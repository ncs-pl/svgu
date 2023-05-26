// The svgu tool.
package main // import "go.nc0.fr/svgu"

import (
	"flag"
	"log"
)

var (
	cfg = flag.String("c", "DOMAINS.star", "the configuration file to use.")
	out = flag.String("o", "dst", "output directory")
	verbose = flag.Bool("v", false, "prints additional information logs")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
}
