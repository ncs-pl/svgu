// Copyright Nicolas Paul (2023)
//
// * Nicolas Paul
//
// This software is a computer program whose purpose is to allow the hosting
// and sharing of Go modules using a personal domain.
//
// This software is governed by the CeCILL license under French law and
// abiding by the rules of distribution of free software.  You can  use,
// modify and/ or redistribute the software under the terms of the CeCILL
// license as circulated by CEA, CNRS and INRIA at the following URL
// "http://www.cecill.info".
//
// As a counterpart to the access to the source code and  rights to copy,
// modify and redistribute granted by the license, users are provided only
// with a limited warranty  and the software's author,  the holder of the
// economic rights,  and the successive licensors  have only  limited
// liability.
//
// In this respect, the user's attention is drawn to the risks associated
// with loading,  using,  modifying and/or developing or reproducing the
// software by the user in light of its specific status of free software,
// that may mean  that it is complicated to manipulate,  and  that  also
// therefore means  that it is reserved for developers  and  experienced
// professionals having in-depth computer knowledge. Users are therefore
// encouraged to load and test the software's suitability as regards their
// requirements in conditions enabling the security of their systems and/or
// data to be ensured and,  more generally, to use and operate it in the
// same conditions as regards security.
//
// The fact that you are presently reading this means that you have had
// knowledge of the CeCILL license and that you accept its terms.

// The svgu tool.
package main // import "go.nc0.fr/svgu"

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sync"

	"go.nc0.fr/svgu/pkg/config"
	"go.nc0.fr/svgu/pkg/types"
)

var (
	cfg     = flag.String("c", "DOMAINS.star", "the configuration file to use.")
	out     = flag.String("o", "dst", "output directory")
	verbose = flag.Bool("v", false, "prints additional information logs")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("svgu: ")
	flag.Parse()

	// Check if the configuration file exists.
	if *verbose {
		log.Printf("checking if configuration file %q exists", *cfg)
	}

	if cfg, err := filepath.Abs(*cfg); err != nil {
		log.Fatalf("could not get absolute path of %s: %v", cfg, err)
	}

	if cfgfd, err := os.Stat(*cfg); os.IsNotExist(err) || cfgfd.IsDir() {
		log.Fatalf("configuration file %q does not exist", *cfg)
	} else if err != nil {
		log.Fatalf("could not stat %q: %v", *cfg, err)
	}

	// Check if the output directory exists.
	if *verbose {
		log.Printf("checking if output directory %q exists", *out)
	}

	if out, err := filepath.Abs(*out); err != nil {
		log.Fatalf("could not get absolute path of %s: %v", out, err)
	}

	if outfd, err := os.Stat(*out); outfd != nil && outfd.IsDir() {
		log.Fatalf("output directory %q already exists", *out)
	} else if err != nil && !os.IsNotExist(err) {
		log.Fatalf("could not stat %q: %v", *out, err)
	}

	// Execute the configuration file and get the registered modules.
	if *verbose {
		log.Printf("executing configuration file %q", *cfg)
	}

	idx, err := config.ExecConfig(*cfg)
	if err != nil {
		log.Fatalf("could not execute configuration file %q: %v", *cfg, err)
	}

	// Create the output directory.
	if *verbose {
		log.Printf("creating output directory %q", *out)
	}

	if err := os.MkdirAll(*out, 0755); err != nil {
		log.Fatalf("could not create output directory %q: %v", *out, err)
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

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, mod := range idx.Modules {
		wg.Add(1)
		go func(m *types.Module) {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			if err := m.GenerateFile(*out, idx.Domain); err != nil {
				log.Fatalf("could not generate module %q: %v", m.Path, err)
			}
		}(mod)
	}

	wg.Wait()
	log.Println("done")
}
