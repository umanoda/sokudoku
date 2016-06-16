package sokudoku

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shogo82148/go-mecab"
)

func Run(wait int) error {
	// Prepare MeCab.
	tagger, err := mecab.New(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		return err
	}
	defer tagger.Destroy()

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
			res, err := tagger.Parse(l)
			if err != nil {
				break
			}
			for _, w := range strings.Split(res, " ") {
				word <- w
			}
		}
		close(word)
	}()

	// Output a word.
	for {
		w, ok := <-word
		if !ok {
			break
		}
		fmt.Print("\r", w)
		time.Sleep(time.Duration(wait)* time.Millisecond)
	}

	return reader.Err()
}
