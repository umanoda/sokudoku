package sokudoku

import "testing"

func TestBuildPhrase(t *testing.T) {
	type testCase struct {
		sample string
		expect []string
	}

	p, _ := phraseInit()

	for _, test_case := range []testCase{
		{
			"ぼくはくま。フワフワした赤い大きな熊。「やあ、こんにちは。」「やあ」と挨拶",
			[]string{"ぼくは", "くま。", "フワフワした", "赤い", "大きな", "熊。", "「やあ、", "こんにちは。」", "「やあ」と", "挨拶"},
		},
		{
			"荒波に負けない心と、切なさと、優しさと。",
			[]string{"荒波に", "負けない", "心と、", "切なさと、", "優しさと。"},
		},
	} {
		got, _ := p.Parse(test_case.sample)
		for i, expect := range test_case.expect {
			if expect != got[i] {
				t.Error(
					"\n   got:", got[i],
					"\nexpect:", expect,
				)
				break
			}
		}
	}

}
