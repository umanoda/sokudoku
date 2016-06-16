# Sokudoku

A simple terminal-base open sourse like Spritz and speedread.

This command line filter show input *Japanese* test as a par-word RSVP (rapid serial visual presentation) aligned on optimal reading points. This kind of input mode allows reading text at a much more rapid pace than usual as the eye can stay fixed on a single place.

## Install

```
$ export CGO_LDFLAGS="-L/path/to/lib -lmecab -lstdc++"
$ export CGO_CFLAGS="-I/path/to/include"
$ go get github.com/shogo82148/go-mecab
$ go get github.com/umanoda/sokudoku
```

## Basic Example

```
cat cat.txt | sokudoku -w 250
```

The default of 250 words per minut is very timid, designed so that you get used to this. Be sure to try cranking this up, 500wpm should still be fairly easy to follow even for beginners.

## Controls

speedread is slightly interactive, with these controls accepted:

* `[` - slow down by 10%
* `]` - speed up by 10%
* `space` - pause (and show the last two lines of context)


## Inspire

[pasky / speedread](https://github.com/pasky/speedread)
