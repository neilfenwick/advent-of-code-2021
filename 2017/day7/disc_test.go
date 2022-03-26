package main

import (
	"testing"
)

//goland:noinspection SpellCheckingInspection
func TestParseSimpleDisc(t *testing.T) {
	testData := map[string]Disc{
		"pbga (66)":                     {"pbga", 66, nil},
		"fwft (72) -> ktlj, cntj, xhth": {"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
	}

	for key, d := range testData {
		disc := ParseDisc(key)
		disc.assertDisc(t, d.Name, d.Weight, d.Children)
	}
}

func (d *Disc) assertDisc(t *testing.T, name string, weight int, children []string) {
	if d.Name != name {
		t.Errorf("Unexpected name for disc: %s. Expected: pbga", d.Name)
	}

	if d.Weight != weight {
		t.Errorf("Unexpected weight for disc: %d. Expected: 66", d.Weight)
	}

	if len(d.Children) != len(children) {
		t.Errorf("Unexpected children count for disc: %#v. Expected: []string{xqmnq, iyoqt, dimle}", d.Children)
	}
}
