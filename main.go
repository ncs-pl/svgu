// Copyright (c) 2023 Nicolas Paul All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	inputfile = flag.String("i", "", "Input configuration file.")
	outputdir = flag.String("o", "dist",
		"Place output files in the specified directory, by default it is `dist`.")
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	log.SetPrefix("staticgovanityurls: ")

	if *inputfile == "" {
		log.Fatalln("Missing input configuration file.")
	}

	if *outputdir == "" {
		log.Fatalln("Missing output directory.")
	}

	// Read the input configuration file and parse it.
	b, err := os.ReadFile(*inputfile)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v\n", *inputfile, err)
	}

	cfg := new(Config)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		log.Fatalf("Failed to parse configuration file: %v\n", err)
	}

	// Create output directory and files
	// The index file ($OUTPUT/index.html) should contain an index of the
	// modules listed in cfg.paths.
	if err := os.Mkdir(*outputdir, 0777); err != nil {
		log.Fatalf("Unable to create output directory %s: %v\n", *outputdir, err)
	}

	// Create a list of all paths' prefixes with the hostname.
	var pfs []string
	for _, p := range cfg.Paths {
		pfs = append(pfs, fmt.Sprintf("%s/%s", cfg.Hostname, p.Prefix))
	}

	// Create index file
	idxpath := fmt.Sprintf("%s/index.html", *outputdir)
	indexfl, err := os.Create(idxpath)
	if err != nil {
		log.Fatalf("Unable to create file %s: %v\n", idxpath, err)
	}
	defer func(fl *os.File) {
		err := fl.Close()
		if err != nil {
			log.Fatalf("Cannot close file %s: %v\n", idxpath, err)
		}
	}(indexfl)

	if err := executeIndex(indexfl, cfg.Hostname, pfs); err != nil {
		log.Fatalf("Cannot execute template: %v\n", err)
	}
	log.Printf("Generated %s\n", idxpath)

	// Create path files
	for _, p := range cfg.Paths {
		if err := os.Mkdir(fmt.Sprintf("%s/%s", *outputdir, p.Prefix), 0777); err != nil {
			log.Fatalf("Unable to create directory %s/%s: %v\n", *outputdir, p.Prefix, err)
		}

		ppath := fmt.Sprintf("%s/%s/index.html", *outputdir, p.Prefix)
		pathfl, err := os.Create(ppath)
		if err != nil {
			log.Fatalf("Unable to create file %s: %v\n", ppath, err)
		}
		defer func(fl *os.File) {
			err := fl.Close()
			if err != nil {
				log.Fatalf("Cannot close file %s: %v\n", ppath, err)
			}
		}(pathfl)

		executePath(pathfl, cfg.Hostname, p.Prefix, p.Vcs, p.Repository, p.Dir, p.File)
		log.Printf("Generated %s\n", ppath)
	}
}
