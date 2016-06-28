package sokudoku

import (
	"bufio"
	"fmt"
	"os"
	"time"
	//"github.com/nsf/termbox-go"
)

func Run(wait int) error {
	p, err := phraseInit()
	if err != nil {
		return err
	}
	defer p.Destroy()

	// Get a line from stdin.
	reader := bufio.NewScanner(os.Stdin)
	line := make(chan string, 1024)
	go func() {
		for reader.Scan() {
			line <- reader.Text()
		}
		close(line)
	}()

	// Parse a line.
	word := make(chan string, 1024)
	go func() {
		for {
			l, ok := <-line
			if !ok {
				break
			}
			res, err := p.Parse(l)
			if err != nil {
				break
			}
			for _, w := range res {
				word <- w
			}
		}
		close(word)
	}()

	// Output a word.
	outputWord(word, wait)

	return reader.Err()
}

func outputWord(word <-chan string, wait int) {
	for {
		w, ok := <-word
		if !ok {
			break
		}
		if DEBUG {
			fmt.Println("...", w)
		} else {
			fmt.Println(w)
		}
		time.Sleep(time.Duration(wait) * time.Millisecond)
	}
}
