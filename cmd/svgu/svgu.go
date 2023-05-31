// The svgu tool.
package main // import "go.nc0.fr/svgu"

import (
	"flag"
	"go.nc0.fr/svgu/pkg/config"
	"go.nc0.fr/svgu/pkg/types"
	"log"
	"os"
	"path/filepath"
)

var (
	cfg     = flag.String("c", "DOMAINS.star", "the configuration file to use.")
	out     = flag.String("o", "dst", "output directory")
	verbose = flag.Bool("v", false, "prints additional information logs")
) // todo(nc0): verbose

func main() {
	log.SetFlags(0)
	log.SetPrefix("svgu: ")
	flag.Parse()

	// Check if the configuration file exists.
	if *verbose {
		log.Printf("checking if configuration file %q exists", cfg)
	}

	if cfg, err := filepath.Abs(*cfg); err != nil {
		log.Fatalf("could not get absolute path of %s: %v", cfg, err)
	}

	if cfgfd, err := os.Stat(*cfg); os.IsNotExist(err) || cfgfd.IsDir() {
		log.Fatalf("configuration file %q does not exist", cfg)
	} else if err != nil {
		log.Fatalf("could not stat %q: %v", cfg, err)
	}

	// Check if the output directory exists.
	if *verbose {
		log.Printf("checking if output directory %q exists", out)
	}

	if out, err := filepath.Abs(*out); err != nil {
		log.Fatalf("could not get absolute path of %s: %v", out, err)
	}

	if outfd, err := os.Stat(*out); outfd != nil && outfd.IsDir() {
		log.Fatalf("output directory %q already exists", out)
	} else if err != nil && !os.IsNotExist(err) {
		log.Fatalf("could not stat %q: %v", out, err)
	}

	// Execute the configuration file and get the registered modules.
	if *verbose {
		log.Printf("executing configuration file %q", cfg)
	}

	idx, err := config.ExecConfig(*cfg)
	if err != nil {
		log.Fatalf("could not execute configuration file %q: %v", cfg, err)
	}

	// Create the output directory.
	if *verbose {
		log.Printf("creating output directory %q", out)
	}

	if err := os.MkdirAll(*out, 0755); err != nil {
		log.Fatalf("could not create output directory %q: %v", out, err)
	}

	// Generate the index file.
	if *verbose {
		log.Printf("generating index file")
	}

	if err := idx.GenerateFile(*out); err != nil {
		log.Fatalf("could not generate index file: %v", err)
	}

	// Generate the modules.
	if *verbose {
		log.Printf("generating modules")
	}

	var mod *types.Module
	for _, mod = range idx.Modules {
		if err := mod.GenerateFile(*out, idx.Domain); err != nil {
			log.Fatalf("could not generate module %q: %v", mod.Path, err)
		}
	}
}
