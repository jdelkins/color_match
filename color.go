package main

import (
	"fmt"
	"math"
	"os"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/pborman/getopt/v2"
)

var (
	mycolors    map[string]colorful.Color
	xtermcolors []colorful.Color
	outFormat   *string
)

func minKey(m map[string]float64) string {
	mink, minv := "", math.MaxFloat64
	for k, v := range m {
		if v < minv {
			minv = v
			mink = k
		}
	}
	return mink
}

func  BestMycolor(tgt *colorful.Color) string {
	dist := make(map[string]float64)
	for n, c := range mycolors {
		dist[n] = tgt.DistanceLab(c)
	}
	return minKey(dist)
}

func initMyColors() {
	mycolor_names := []string{
		"bg",
		"black",
		"blue",
		"cyan",
		"fg",
		"green",
		"grey01",
		"grey02",
		"grey03",
		"grey1",
		"grey2",
		"grey3",
		"magenta",
		"orange",
		"red",
		"violet",
		"white",
		"yellow",
	}
	mycolors = make(map[string]colorful.Color)
	for _, c := range mycolor_names {
		h, _ := colorful.Hex(os.Getenv("mycolor_" + c))
		mycolors[c] = h
	}
}

func initXtermColors() {
	xtermcolors = make([]colorful.Color, 256)

	xtermcolors[0] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[1] = colorful.Color{R: float64(0xcd)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[2] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xcd)/255.0, B: float64(0x00)/255.0}
	xtermcolors[3] = colorful.Color{R: float64(0xcd)/255.0, G: float64(0xcd)/255.0, B: float64(0x00)/255.0}
	xtermcolors[4] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0xee)/255.0}
	xtermcolors[5] = colorful.Color{R: float64(0xcd)/255.0, G: float64(0x00)/255.0, B: float64(0xcd)/255.0}
	xtermcolors[6] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xcd)/255.0, B: float64(0xcd)/255.0}
	xtermcolors[7] = colorful.Color{R: float64(0xe5)/255.0, G: float64(0xe5)/255.0, B: float64(0xe5)/255.0}
	xtermcolors[8] = colorful.Color{R: float64(0x7f)/255.0, G: float64(0x7f)/255.0, B: float64(0x7f)/255.0}
	xtermcolors[9] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[10] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[11] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[12] = colorful.Color{R: float64(0x5c)/255.0, G: float64(0x5c)/255.0, B: float64(0xff)/255.0}
	xtermcolors[13] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[14] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[15] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[16] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[17] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[18] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[19] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[20] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[21] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[22] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[23] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[24] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[25] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[26] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[27] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[28] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[29] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[30] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[31] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[32] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[33] = colorful.Color{R: float64(0x00)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[34] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[35] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[36] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[37] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[38] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[39] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[40] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[41] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[42] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[43] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[44] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[45] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[46] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[47] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[48] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[49] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[50] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[51] = colorful.Color{R: float64(0x00)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[52] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[53] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[54] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[55] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[56] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[57] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[58] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[59] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[60] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[61] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[62] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[63] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[64] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[65] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[66] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[67] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[68] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[69] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[70] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[71] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[72] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[73] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[74] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[75] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[76] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[77] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[78] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[79] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[80] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[81] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[82] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[83] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[84] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[85] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[86] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[87] = colorful.Color{R: float64(0x5f)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[88] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[89] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[90] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[91] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[92] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[93] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[94] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[95] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[96] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[97] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[98] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[99] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[100] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[101] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[102] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[103] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[104] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[105] = colorful.Color{R: float64(0x87)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[106] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[107] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[108] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[109] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[110] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[111] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[112] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[113] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[114] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[115] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[116] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[117] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[118] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[119] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[120] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[121] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[122] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[123] = colorful.Color{R: float64(0x87)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[124] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[125] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[126] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[127] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[128] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[129] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[130] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[131] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[132] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[133] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[134] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[135] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[136] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[137] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[138] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[139] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[140] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[141] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[142] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[143] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[144] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[145] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[146] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[147] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[148] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[149] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[150] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[151] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[152] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[153] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[154] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[155] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[156] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[157] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[158] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[159] = colorful.Color{R: float64(0xaf)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[160] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[161] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[162] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[163] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[164] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[165] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[166] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[167] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[168] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[169] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[170] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[171] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[172] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[173] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[174] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[175] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[176] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[177] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[178] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[179] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[180] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[181] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[182] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[183] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[184] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[185] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[186] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[187] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[188] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[189] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[190] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[191] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[192] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[193] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[194] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[195] = colorful.Color{R: float64(0xd7)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[196] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0x00)/255.0}
	xtermcolors[197] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[198] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0x87)/255.0}
	xtermcolors[199] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[200] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[201] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x00)/255.0, B: float64(0xff)/255.0}
	xtermcolors[202] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0x00)/255.0}
	xtermcolors[203] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[204] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0x87)/255.0}
	xtermcolors[205] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[206] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[207] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x5f)/255.0, B: float64(0xff)/255.0}
	xtermcolors[208] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0x00)/255.0}
	xtermcolors[209] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[210] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0x87)/255.0}
	xtermcolors[211] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[212] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[213] = colorful.Color{R: float64(0xff)/255.0, G: float64(0x87)/255.0, B: float64(0xff)/255.0}
	xtermcolors[214] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0x00)/255.0}
	xtermcolors[215] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[216] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0x87)/255.0}
	xtermcolors[217] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[218] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[219] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xaf)/255.0, B: float64(0xff)/255.0}
	xtermcolors[220] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0x00)/255.0}
	xtermcolors[221] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[222] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0x87)/255.0}
	xtermcolors[223] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[224] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[225] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xd7)/255.0, B: float64(0xff)/255.0}
	xtermcolors[226] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0x00)/255.0}
	xtermcolors[227] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0x5f)/255.0}
	xtermcolors[228] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0x87)/255.0}
	xtermcolors[229] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0xaf)/255.0}
	xtermcolors[230] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0xd7)/255.0}
	xtermcolors[231] = colorful.Color{R: float64(0xff)/255.0, G: float64(0xff)/255.0, B: float64(0xff)/255.0}
	xtermcolors[232] = colorful.Color{R: float64(0x08)/255.0, G: float64(0x08)/255.0, B: float64(0x08)/255.0}
	xtermcolors[233] = colorful.Color{R: float64(0x12)/255.0, G: float64(0x12)/255.0, B: float64(0x12)/255.0}
	xtermcolors[234] = colorful.Color{R: float64(0x1c)/255.0, G: float64(0x1c)/255.0, B: float64(0x1c)/255.0}
	xtermcolors[235] = colorful.Color{R: float64(0x26)/255.0, G: float64(0x26)/255.0, B: float64(0x26)/255.0}
	xtermcolors[236] = colorful.Color{R: float64(0x30)/255.0, G: float64(0x30)/255.0, B: float64(0x30)/255.0}
	xtermcolors[237] = colorful.Color{R: float64(0x3a)/255.0, G: float64(0x3a)/255.0, B: float64(0x3a)/255.0}
	xtermcolors[238] = colorful.Color{R: float64(0x44)/255.0, G: float64(0x44)/255.0, B: float64(0x44)/255.0}
	xtermcolors[239] = colorful.Color{R: float64(0x4e)/255.0, G: float64(0x4e)/255.0, B: float64(0x4e)/255.0}
	xtermcolors[240] = colorful.Color{R: float64(0x58)/255.0, G: float64(0x58)/255.0, B: float64(0x58)/255.0}
	xtermcolors[241] = colorful.Color{R: float64(0x62)/255.0, G: float64(0x62)/255.0, B: float64(0x62)/255.0}
	xtermcolors[242] = colorful.Color{R: float64(0x6c)/255.0, G: float64(0x6c)/255.0, B: float64(0x6c)/255.0}
	xtermcolors[243] = colorful.Color{R: float64(0x76)/255.0, G: float64(0x76)/255.0, B: float64(0x76)/255.0}
	xtermcolors[244] = colorful.Color{R: float64(0x80)/255.0, G: float64(0x80)/255.0, B: float64(0x80)/255.0}
	xtermcolors[245] = colorful.Color{R: float64(0x8a)/255.0, G: float64(0x8a)/255.0, B: float64(0x8a)/255.0}
	xtermcolors[246] = colorful.Color{R: float64(0x94)/255.0, G: float64(0x94)/255.0, B: float64(0x94)/255.0}
	xtermcolors[247] = colorful.Color{R: float64(0x9e)/255.0, G: float64(0x9e)/255.0, B: float64(0x9e)/255.0}
	xtermcolors[248] = colorful.Color{R: float64(0xa8)/255.0, G: float64(0xa8)/255.0, B: float64(0xa8)/255.0}
	xtermcolors[249] = colorful.Color{R: float64(0xb2)/255.0, G: float64(0xb2)/255.0, B: float64(0xb2)/255.0}
	xtermcolors[250] = colorful.Color{R: float64(0xbc)/255.0, G: float64(0xbc)/255.0, B: float64(0xbc)/255.0}
	xtermcolors[251] = colorful.Color{R: float64(0xc6)/255.0, G: float64(0xc6)/255.0, B: float64(0xc6)/255.0}
	xtermcolors[252] = colorful.Color{R: float64(0xd0)/255.0, G: float64(0xd0)/255.0, B: float64(0xd0)/255.0}
	xtermcolors[253] = colorful.Color{R: float64(0xda)/255.0, G: float64(0xda)/255.0, B: float64(0xda)/255.0}
	xtermcolors[254] = colorful.Color{R: float64(0xe4)/255.0, G: float64(0xe4)/255.0, B: float64(0xe4)/255.0}
	xtermcolors[255] = colorful.Color{R: float64(0xee)/255.0, G: float64(0xee)/255.0, B: float64(0xee)/255.0}
}


func output(k string) {
	c := mycolors[k]
	if *outFormat != "" {
		fmt.Printf(*outFormat, c.Hex())
	} else {
		fmt.Printf("mycolor_%s", k)
	}
}

func main() {
	helpFlag := getopt.BoolLong("help", 'h', "display help")
	ansiColor := getopt.IntLong("ansi", 'a', -1, "match 'standard' ansi 256 color value")
	hexColor := getopt.StringLong("hex", 'x', "", "match hex color in #ffffff format")
	outFormat = getopt.StringLong("format", 'f', "", "output format. use %s for the hex value")
	getopt.Parse()

	if *helpFlag {
		getopt.PrintUsage(os.Stdout)
		return
	}

	if (*ansiColor < 0 && *hexColor == "") || (*ansiColor >= 0 && *hexColor != "") || *ansiColor > 255 {
		getopt.Usage()
		os.Exit(1)
	}

	initMyColors()
	if *ansiColor >= 0 {
		initXtermColors()
		output(BestMycolor(&xtermcolors[*ansiColor]))
	} else {
		col, err := colorful.Hex(*hexColor)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: invalid hex color spec: %s\n", os.Args[0], *hexColor)
			os.Exit(1)
		}
		output(BestMycolor(&col))
	}
}
