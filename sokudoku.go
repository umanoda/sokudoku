package sokudoku

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/mattn/go-tty"
	//debug
	//"github.com/k0kubun/pp"
)

// 表示文字列の左側余白長さ(全角文字長)
const visulapos = 22

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
			fmt.Println(_showWord(w))
			time.Sleep(time.Duration(wait) * time.Millisecond)
		}
	} else {
		tty, err := tty.Open()
		if err != nil {
			return err
		}
		defer func() {
			for {
				fmt.Println("\n\n  Exit. Please press ENTER.")
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

		fmt.Println(strings.Repeat(" ", visulapos-1), "・")

		// pause flag
		stop := false
	loop:
		for {
			if stop {
				// in pause
				r := <-key_queue
				if r == 32 {
					// pressed Space
					stop = false
				}
			} else {
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
						stop = true
					}
				case w, ok := <-word:
					if !ok {
						break loop
					}
					showWord(w)
					time.Sleep(time.Duration(wait) * time.Millisecond)
				}
			}
		}
	}
	return nil
}

// 画面に文字を表示する
func showWord(str string) {
	s := _showWord(str)
	fmt.Print(
		"\r\033[K", //行消去
		ansiColor("bold"),
		s[0],
		s[1],
		ansiColor("red"),
		s[2],
		ansiColor("reset"),
		ansiColor("bold"),
		s[3],
		ansiColor("reset"),
	)
}

func _showWord(str string) []string {
	str_len := utf8.RuneCountInString(str)

	// 強調文字位置
	i := (str_len - 1) / 2

	return []string{
		strings.Repeat(" ", visulapos-(i*2)),
		substr(str, 0, i),
		substr(str, i, 1),
		substr(str, i+1, -1),
	}
}

func substr(str string, start int, length int) string {
	if length == 0 {
		return ""
	}

	r := []rune(str)
	if length == -1 {
		return string(r[start:len(r)])
	} else {
		return string(r[start:(start + length)])
	}
}

func ansiColor(c string) string {
	switch c {
	case "red":
		return "\033[1;31m"
	case "reset":
		return "\033[0;37m"
	case "bold":
		return "\033[1;37m"
	default:
		return ""
	}
}
