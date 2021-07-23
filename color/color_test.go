package color_test

import (
	"testing"

	"github.com/vompressor/vfmlib/vfmlib/color"
	"github.com/vompressor/vfmlib/vfmlib/color/background_color"
	"github.com/vompressor/vfmlib/vfmlib/color/basic"
	"github.com/vompressor/vfmlib/vfmlib/color/text_color"
)

func TestColor(t *testing.T) {
	println(color.NewAtt(
		text_color.Cyan,
		background_color.BgHiRed,
		basic.Bold,
	).ColorString("hello"))
}
