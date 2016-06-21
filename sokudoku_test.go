package sokudoku

import "testing"

func TestBuildPhrase(t *testing.T) {
	p, _ := phraseInit()

	got, _ := p.Parse("ぼくはくま。赤い大きな熊。「やあ、こんにちは。」「やあ」と挨拶")

	expects := []string{
		"ぼくは",
		"くま。",
		"赤い",
		"大きな",
		"熊。",
		"「やあ、",
		"こんにちは。」",
		"「やあ」と",
		"挨拶",
	}
	for i, expect := range expects {
		if expect != got[i] {
			t.Error(
				"\n   got:", got[i],
				"\nexpect:", expect,
			)
			break
		}
	}
}
