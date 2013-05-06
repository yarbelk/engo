package eng

import (
	"math/rand"
)

type Color struct {
	R, G, B, A float32
}

func NewColor(r, g, b, a float32) *Color {
	return &Color{r, g, b, a}
}

func NewColorBytes(r, g, b, a byte) *Color {
	color := new(Color)
	color.R = float32(r) / 255.0
	color.G = float32(g) / 255.0
	color.B = float32(b) / 255.0
	color.A = float32(a) / 255.0
	return color
}

func NewColorBytesA(r, g, b byte) *Color {
	color := new(Color)
	color.R = float32(r) / 255.0
	color.G = float32(g) / 255.0
	color.B = float32(b) / 255.0
	color.A = float32(1)
	return color
}

func NewColorRand() *Color {
	return &Color{rand.Float32(), rand.Float32(), rand.Float32(), 1}
}

func NewColorRandA() *Color {
	return &Color{rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()}
}

var (
	Black                 = NewColorBytes(0, 0, 0, 255)
	DarkestGrey           = NewColorBytes(31, 31, 31, 255)
	DarkerGrey            = NewColorBytes(63, 63, 63, 255)
	DarkGrey              = NewColorBytes(95, 95, 95, 255)
	Grey                  = NewColorBytes(127, 127, 127, 255)
	LightGrey             = NewColorBytes(159, 159, 159, 255)
	LighterGrey           = NewColorBytes(191, 191, 191, 255)
	LightestGrey          = NewColorBytes(223, 223, 223, 255)
	White                 = NewColorBytes(255, 255, 255, 255)
	DarkestSepia          = NewColorBytes(31, 24, 15, 255)
	DarkerSepia           = NewColorBytes(63, 50, 31, 255)
	DarkSepia             = NewColorBytes(94, 75, 47, 255)
	Sepia                 = NewColorBytes(127, 101, 63, 255)
	LightSepia            = NewColorBytes(158, 134, 100, 255)
	LighterSepia          = NewColorBytes(191, 171, 143, 255)
	LightestSepia         = NewColorBytes(222, 211, 195, 255)
	DesaturatedRed        = NewColorBytes(127, 63, 63, 255)
	DesaturatedFlame      = NewColorBytes(127, 79, 63, 255)
	DesaturatedOrange     = NewColorBytes(127, 95, 63, 255)
	DesaturatedAmber      = NewColorBytes(127, 111, 63, 255)
	DesaturatedYellow     = NewColorBytes(127, 127, 63, 255)
	DesaturatedLime       = NewColorBytes(111, 127, 63, 255)
	DesaturatedChartreuse = NewColorBytes(95, 127, 63, 255)
	DesaturatedGreen      = NewColorBytes(63, 127, 63, 255)
	DesaturatedSea        = NewColorBytes(63, 127, 95, 255)
	DesaturatedTurquoise  = NewColorBytes(63, 127, 111, 255)
	DesaturatedCyan       = NewColorBytes(63, 127, 127, 255)
	DesaturatedSky        = NewColorBytes(63, 111, 127, 255)
	DesaturatedAzure      = NewColorBytes(63, 95, 127, 255)
	DesaturatedBlue       = NewColorBytes(63, 63, 127, 255)
	DesaturatedHan        = NewColorBytes(79, 63, 127, 255)
	DesaturatedViolet     = NewColorBytes(95, 63, 127, 255)
	DesaturatedPurple     = NewColorBytes(111, 63, 127, 255)
	DesaturatedFuchsia    = NewColorBytes(127, 63, 127, 255)
	DesaturatedMagenta    = NewColorBytes(127, 63, 111, 255)
	DesaturatedPink       = NewColorBytes(127, 63, 95, 255)
	DesaturatedCrimson    = NewColorBytes(127, 63, 79, 255)
	LightestRed           = NewColorBytes(255, 191, 191, 255)
	LightestFlame         = NewColorBytes(255, 207, 191, 255)
	LightestOrange        = NewColorBytes(255, 223, 191, 255)
	LightestAmber         = NewColorBytes(255, 239, 191, 255)
	LightestYellow        = NewColorBytes(255, 255, 191, 255)
	LightestLime          = NewColorBytes(239, 255, 191, 255)
	LightestChartreuse    = NewColorBytes(223, 255, 191, 255)
	LightestGreen         = NewColorBytes(191, 255, 191, 255)
	LightestSea           = NewColorBytes(191, 255, 223, 255)
	LightestTurquoise     = NewColorBytes(191, 255, 239, 255)
	LightestCyan          = NewColorBytes(191, 255, 255, 255)
	LightestSky           = NewColorBytes(191, 239, 255, 255)
	LightestAzure         = NewColorBytes(191, 223, 255, 255)
	LightestBlue          = NewColorBytes(191, 191, 255, 255)
	LightestHan           = NewColorBytes(207, 191, 255, 255)
	LightestViolet        = NewColorBytes(223, 191, 255, 255)
	LightestPurple        = NewColorBytes(239, 191, 255, 255)
	LightestFuchsia       = NewColorBytes(255, 191, 255, 255)
	LightestMagenta       = NewColorBytes(255, 191, 239, 255)
	LightestPink          = NewColorBytes(255, 191, 223, 255)
	LightestCrimson       = NewColorBytes(255, 191, 207, 255)
	LighterRed            = NewColorBytes(255, 127, 127, 255)
	LighterFlame          = NewColorBytes(255, 159, 127, 255)
	LighterOrange         = NewColorBytes(255, 191, 127, 255)
	LighterAmber          = NewColorBytes(255, 223, 127, 255)
	LighterYellow         = NewColorBytes(255, 255, 127, 255)
	LighterLime           = NewColorBytes(223, 255, 127, 255)
	LighterChartreuse     = NewColorBytes(191, 255, 127, 255)
	LighterGreen          = NewColorBytes(127, 255, 127, 255)
	LighterSea            = NewColorBytes(127, 255, 191, 255)
	LighterTurquoise      = NewColorBytes(127, 255, 223, 255)
	LighterCyan           = NewColorBytes(127, 255, 255, 255)
	LighterSky            = NewColorBytes(127, 223, 255, 255)
	LighterAzure          = NewColorBytes(127, 191, 255, 255)
	LighterBlue           = NewColorBytes(127, 127, 255, 255)
	LighterHan            = NewColorBytes(159, 127, 255, 255)
	LighterViolet         = NewColorBytes(191, 127, 255, 255)
	LighterPurple         = NewColorBytes(223, 127, 255, 255)
	LighterFuchsia        = NewColorBytes(255, 127, 255, 255)
	LighterMagenta        = NewColorBytes(255, 127, 223, 255)
	LighterPink           = NewColorBytes(255, 127, 191, 255)
	LighterCrimson        = NewColorBytes(255, 127, 159, 255)
	LightRed              = NewColorBytes(255, 63, 63, 255)
	LightFlame            = NewColorBytes(255, 111, 63, 255)
	LightOrange           = NewColorBytes(255, 159, 63, 255)
	LightAmber            = NewColorBytes(255, 207, 63, 255)
	LightYellow           = NewColorBytes(255, 255, 63, 255)
	LightLime             = NewColorBytes(207, 255, 63, 255)
	LightChartreuse       = NewColorBytes(159, 255, 63, 255)
	LightGreen            = NewColorBytes(63, 255, 63, 255)
	LightSea              = NewColorBytes(63, 255, 159, 255)
	LightTurquoise        = NewColorBytes(63, 255, 207, 255)
	LightCyan             = NewColorBytes(63, 255, 255, 255)
	LightSky              = NewColorBytes(63, 207, 255, 255)
	LightAzure            = NewColorBytes(63, 159, 255, 255)
	LightBlue             = NewColorBytes(63, 63, 255, 255)
	LightHan              = NewColorBytes(111, 63, 255, 255)
	LightViolet           = NewColorBytes(159, 63, 255, 255)
	LightPurple           = NewColorBytes(207, 63, 255, 255)
	LightFuchsia          = NewColorBytes(255, 63, 255, 255)
	LightMagenta          = NewColorBytes(255, 63, 207, 255)
	LightPink             = NewColorBytes(255, 63, 159, 255)
	LightCrimson          = NewColorBytes(255, 63, 111, 255)
	Red                   = NewColorBytes(255, 0, 0, 255)
	Flame                 = NewColorBytes(255, 63, 0, 255)
	Orange                = NewColorBytes(255, 127, 0, 255)
	Amber                 = NewColorBytes(255, 191, 0, 255)
	Yellow                = NewColorBytes(255, 255, 0, 255)
	Lime                  = NewColorBytes(191, 255, 0, 255)
	Chartreuse            = NewColorBytes(127, 255, 0, 255)
	Green                 = NewColorBytes(0, 255, 0, 255)
	Sea                   = NewColorBytes(0, 255, 127, 255)
	Turquoise             = NewColorBytes(0, 255, 191, 255)
	Cyan                  = NewColorBytes(0, 255, 255, 255)
	Sky                   = NewColorBytes(0, 191, 255, 255)
	Azure                 = NewColorBytes(0, 127, 255, 255)
	Blue                  = NewColorBytes(0, 0, 255, 255)
	Han                   = NewColorBytes(63, 0, 255, 255)
	Violet                = NewColorBytes(127, 0, 255, 255)
	Purple                = NewColorBytes(191, 0, 255, 255)
	Fuchsia               = NewColorBytes(255, 0, 255, 255)
	Magenta               = NewColorBytes(255, 0, 191, 255)
	Pink                  = NewColorBytes(255, 0, 127, 255)
	Crimson               = NewColorBytes(255, 0, 63, 255)
	DarkRed               = NewColorBytes(191, 0, 0, 255)
	DarkFlame             = NewColorBytes(191, 47, 0, 255)
	DarkOrange            = NewColorBytes(191, 95, 0, 255)
	DarkAmber             = NewColorBytes(191, 143, 0, 255)
	DarkYellow            = NewColorBytes(191, 191, 0, 255)
	DarkLime              = NewColorBytes(143, 191, 0, 255)
	DarkChartreuse        = NewColorBytes(95, 191, 0, 255)
	DarkGreen             = NewColorBytes(0, 191, 0, 255)
	DarkSea               = NewColorBytes(0, 191, 95, 255)
	DarkTurquoise         = NewColorBytes(0, 191, 143, 255)
	DarkCyan              = NewColorBytes(0, 191, 191, 255)
	DarkSky               = NewColorBytes(0, 143, 191, 255)
	DarkAzure             = NewColorBytes(0, 95, 191, 255)
	DarkBlue              = NewColorBytes(0, 0, 191, 255)
	DarkHan               = NewColorBytes(47, 0, 191, 255)
	DarkViolet            = NewColorBytes(95, 0, 191, 255)
	DarkPurple            = NewColorBytes(143, 0, 191, 255)
	DarkFuchsia           = NewColorBytes(191, 0, 191, 255)
	DarkMagenta           = NewColorBytes(191, 0, 143, 255)
	DarkPink              = NewColorBytes(191, 0, 95, 255)
	DarkCrimson           = NewColorBytes(191, 0, 47, 255)
	DarkerRed             = NewColorBytes(127, 0, 0, 255)
	DarkerFlame           = NewColorBytes(127, 31, 0, 255)
	DarkerOrange          = NewColorBytes(127, 63, 0, 255)
	DarkerAmber           = NewColorBytes(127, 95, 0, 255)
	DarkerYellow          = NewColorBytes(127, 127, 0, 255)
	DarkerLime            = NewColorBytes(95, 127, 0, 255)
	DarkerChartreuse      = NewColorBytes(63, 127, 0, 255)
	DarkerGreen           = NewColorBytes(0, 127, 0, 255)
	DarkerSea             = NewColorBytes(0, 127, 63, 255)
	DarkerTurquoise       = NewColorBytes(0, 127, 95, 255)
	DarkerCyan            = NewColorBytes(0, 127, 127, 255)
	DarkerSky             = NewColorBytes(0, 95, 127, 255)
	DarkerAzure           = NewColorBytes(0, 63, 127, 255)
	DarkerBlue            = NewColorBytes(0, 0, 127, 255)
	DarkerHan             = NewColorBytes(31, 0, 127, 255)
	DarkerViolet          = NewColorBytes(63, 0, 127, 255)
	DarkerPurple          = NewColorBytes(95, 0, 127, 255)
	DarkerFuchsia         = NewColorBytes(127, 0, 127, 255)
	DarkerMagenta         = NewColorBytes(127, 0, 95, 255)
	DarkerPink            = NewColorBytes(127, 0, 63, 255)
	DarkerCrimson         = NewColorBytes(127, 0, 31, 255)
	DarkestRed            = NewColorBytes(63, 0, 0, 255)
	DarkestFlame          = NewColorBytes(63, 15, 0, 255)
	DarkestOrange         = NewColorBytes(63, 31, 0, 255)
	DarkestAmber          = NewColorBytes(63, 47, 0, 255)
	DarkestYellow         = NewColorBytes(63, 63, 0, 255)
	DarkestLime           = NewColorBytes(47, 63, 0, 255)
	DarkestChartreuse     = NewColorBytes(31, 63, 0, 255)
	DarkestGreen          = NewColorBytes(0, 63, 0, 255)
	DarkestSea            = NewColorBytes(0, 63, 31, 255)
	DarkestTurquoise      = NewColorBytes(0, 63, 47, 255)
	DarkestCyan           = NewColorBytes(0, 63, 63, 255)
	DarkestSky            = NewColorBytes(0, 47, 63, 255)
	DarkestAzure          = NewColorBytes(0, 31, 63, 255)
	DarkestBlue           = NewColorBytes(0, 0, 63, 255)
	DarkestHan            = NewColorBytes(15, 0, 63, 255)
	DarkestViolet         = NewColorBytes(31, 0, 63, 255)
	DarkestPurple         = NewColorBytes(47, 0, 63, 255)
	DarkestFuchsia        = NewColorBytes(63, 0, 63, 255)
	DarkestMagenta        = NewColorBytes(63, 0, 47, 255)
	DarkestPink           = NewColorBytes(63, 0, 31, 255)
	DarkestCrimson        = NewColorBytes(63, 0, 15, 255)
	Brass                 = NewColorBytes(191, 151, 96, 255)
	Copper                = NewColorBytes(197, 136, 124, 255)
	Gold                  = NewColorBytes(229, 191, 0, 255)
	Silver                = NewColorBytes(203, 203, 203, 255)
	Celadon               = NewColorBytes(172, 255, 175, 255)
	Peach                 = NewColorBytes(255, 159, 127, 255)
)
