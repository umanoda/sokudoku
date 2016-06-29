package sokudoku

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/mattn/go-tty"
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
	oerr := outputWord(word, wait)
	if oerr != nil {
		return oerr
	}

	return reader.Err()
}

func outputWord(word <-chan string, wait int) error {
	if DEBUG {
		//-tags debug
		for {
			w, ok := <-word
			if !ok {
				break
			}
			fmt.Println("   ", w)
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
	} else {
		tty, err := tty.Open()
		if err != nil {
			return err
		}
		defer func() {
			for {
				fmt.Println("Please ENTER.")
				r, _ := tty.ReadRune()
				if r == 13 {
					break
				}
			}
			tty.Close()
		}()

		key_queue := make(chan rune)
		go func() {
			for {
				r, _ := tty.ReadRune()
				key_queue <- r
			}
		}()

	loop:
		for {
			select {
			case r := <-key_queue:
				switch r {
				case 3: // Ctrl-C
					break loop
				case 91: // [
					wait = int(math.Max(float64(wait)*0.9, 10))
				case 93: // ]
					wait = int(math.Min(float64(wait)*1.1, 2000))
				case 32: // Space
					// TODO puase
				}
			case w, ok := <-word:
				if !ok {
					break loop
				}
				fmt.Println("...", w)
				time.Sleep(time.Duration(wait) * time.Millisecond)
			}
		}
	}
	return nil
}
