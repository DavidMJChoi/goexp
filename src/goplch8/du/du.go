package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
		fmt.Printf("%d files  %.1d B\n", nfiles, nbytes)
		// fmt.Printf("%d files  %.1f MB\n", nfiles, float64(nbytes)/(1024*1024))
	}
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fInfo, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du: %v\n", err)
				return
			}
			fileSizes <- fInfo.Size()
		}
	}
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
