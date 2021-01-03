package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file with this name")
	bits       = flag.String("bits", "1024", "how many bits for the RSA key")
	memprofile = flag.String("memprofile", "", "write memory profile to file with this name")
)

func main() {
	flag.Parse()
	// https://golang.org/pkg/runtime/pprof/#hdr-Profiling_a_Go_program
	if *cpuprofile != "" {
		log.Printf("writing cpu profile to %s", *cpuprofile)
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	b, err := strconv.Atoi(*bits)
	if err != nil {
		log.Fatalf("error converting %s to int: %v", *bits, err)
	}
	_, err = CreateKey(b)
	if err != nil {
		log.Fatal(err)
	}
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		// Runs  garbage collection- I  guess  so that unused  objects aren't included in the output
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal(err)
		}
	}
}

func CreateKey(bits int) (*rsa.PrivateKey, error) {
	fmt.Printf("Creating an RSA key with %d bits..\n", bits)
	return rsa.GenerateKey(rand.Reader, bits)
}
