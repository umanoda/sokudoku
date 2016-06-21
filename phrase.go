package sokudoku

import (
	"github.com/shogo82148/go-mecab"
	"strings"
	//"github.com/k0kubun/pp"
)

type phrase struct {
	mecab mecab.MeCab
}

func phraseInit() (*phrase, error) {
	p := &phrase{}

	// Prepare MeCab.
	tagger, err := mecab.New(map[string]string{})
	if err != nil {
		return p, err
	}
	p.mecab = tagger

	return p, nil
}

func (p *phrase) Parse(s string) ([]string, error) {
	var res []string
	parsed, err := p.mecab.Parse(s)
	if err != nil {
		return nil, err
	}

	var tmp []string
	var next_force bool
	for _, w := range strings.Split(parsed, "\n") {
		m := strings.Split(w, "\t")
		if len(m) < 2 {
			//EOS
			break
		}
		word := m[0]
		detail := strings.Split(m[1], ",")
		word_type := detail[0]

		if next_force {
			tmp = append(tmp, word)
			next_force = false
			continue
		}
		switch word_type {
		case "助詞", "記号":
			sub_type := detail[1]
			switch sub_type {
			case "括弧開":
				if len(tmp) > 0 {
					res = append(res, strings.Join(tmp, ""))
				}
				tmp = []string{word}
				next_force = true
			default:
				tmp = append(tmp, word)
			}
		default:
			if len(tmp) > 0 {
				res = append(res, strings.Join(tmp, ""))
			}
			tmp = []string{word}
		}
	}
	if len(tmp) > 0 {
		res = append(res, strings.Join(tmp, ""))
	}
	return res, nil
}

func (p *phrase) Destroy() {
	p.mecab.Destroy()
}
