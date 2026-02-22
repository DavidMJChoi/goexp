package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	filenames := []string{
		"img1",
		"img2",
		"img3",
		"img4",
		"img5",
	}

	start := time.Now()
	makeThumbnails3(filenames)
	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)

}

func ImageFile(infile string) (string, error) {
	// simulate image processing
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// n := rand.Intn(10) + 1 // range: [1,10]
	n := 1

	// simulate faillure
	if n == 5 {
		// panic("test")
		return "failed", fmt.Errorf("Thumbnail generation failed for file %s", infile)
	}

	// don't sleep for too long time
	n %= 5
	time.Sleep(time.Duration(n) * 250 * time.Millisecond)

	return infile + "-thumb", nil
}

func makeThumbnails(filenames []string) {
	for _, img := range filenames {
		imgg, e := ImageFile(img)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(imgg)
	}
}

func makeThumbnails3(filenames []string) {
	done := make(chan struct{})
	defer close(done)
	for _, f := range filenames {
		go func(f string) {
			imgg, e := ImageFile(f)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println(imgg)
			done <- struct{}{}
		}(f)
	}
	for range filenames {
		<-done
	}
}

func makeThumbnails4(filenames []string) error {
	errCh := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errCh <- err
		}(f)
	}
	for range filenames {
		if err := <-errCh; err != nil {
			return err
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1) // increment
		// worker
		go func(f string) {
			defer wg.Done() // decrement
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}
	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
