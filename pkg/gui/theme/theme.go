// SPDX-License-Identifier: Unlicense OR MIT

package theme

import (
	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
)

type DuoUItheme struct {
	Shaper                text.Shaper
	TextSize              unit.Value
	checkBoxCheckedIcon   *DuoUIicon
	checkBoxUncheckedIcon *DuoUIicon
	radioCheckedIcon      *DuoUIicon
	radioUncheckedIcon    *DuoUIicon
	counterPlusIcon       *DuoUIicon
	counterMinusIcon      *DuoUIicon
	Colors                map[string]string
	Fonts                 map[string]text.Typeface
	Icons                 map[string]*DuoUIicon
}

func NewDuoUItheme() *DuoUItheme {
	t := &DuoUItheme{
		Shaper: font.Default(),
	}
	t.Colors = NewDuoUIcolors()
	t.Fonts = NewDuoUIfonts()
	t.TextSize = unit.Sp(16)
	t.Icons = NewDuoUIicons()
	return t
}

func NewDuoUIfonts() (f map[string]text.Typeface) {
	f = make(map[string]text.Typeface)
	f["Primary"] = "bariol"
	f["Secondary"] = "plan9"
	f["Mono"] = "go"
	return f
}
