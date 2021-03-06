package xlsx

import "strconv"

// Several popular font names that can be used to create fonts
const (
	Helvetica     = "Helvetica"
	Baskerville   = "Baskerville Old Face"
	TimesNewRoman = "Times New Roman"
	Bodoni        = "Bodoni MT"
	GillSans      = "Gill Sans MT"
	Courier       = "Courier"
)

const (
	RGB_Light_Green = "FFC6EFCE"
	RGB_Dark_Green  = "FF006100"
	RGB_Light_Red   = "FFFFC7CE"
	RGB_Dark_Red    = "FF9C0006"
	RGB_White       = "FFFFFFFF"
	RGB_Black       = "00000000"
)

const (
	Solid_Cell_Fill = "solid"
)

// Style is a high level structure intended to provide user access to
// the contents of Style within an XLSX file.
type Style struct {
	Border          Border
	Fill            Fill
	Font            Font
	ApplyBorder     bool
	ApplyFill       bool
	ApplyFont       bool
	ApplyAlignment  bool
	Alignment       Alignment
	NamedStyleIndex *int
}

// Return a new Style structure initialised with the default values.
func NewStyle() *Style {
	return &Style{
		Alignment: *DefaultAlignment(),
		Border:    *DefaultBorder(),
		Fill:      *DefaultFill(),
		Font:      *DefaultFont(),
	}
}

// Generate the underlying XLSX style elements that correspond to the Style.
func (style *Style) makeXLSXStyleElements() (xFont xlsxFont, xFill xlsxFill, xBorder xlsxBorder, xCellXf xlsxXf) {
	if style == nil {
		panic("Called makeXLSXStyleElements on a nil *Style!")
	}

	xFont = xlsxFont{}
	xFill = xlsxFill{}
	xBorder = xlsxBorder{}
	xCellXf = xlsxXf{}
	xFont.Sz.Val = strconv.FormatFloat(style.Font.Size, 'f', -1, 64)
	xFont.Name.Val = style.Font.Name
	xFont.Family.Val = strconv.Itoa(style.Font.Family)
	xFont.Charset.Val = strconv.Itoa(style.Font.Charset)
	xFont.Color.RGB = style.Font.Color

	if style.Font.Bold {
		xFont.B = &xlsxVal{}
	} else {
		xFont.B = nil
	}
	if style.Font.Italic {
		xFont.I = &xlsxVal{}
	} else {
		xFont.I = nil
	}
	if style.Font.Underline {
		xFont.U = &xlsxVal{}
	} else {
		xFont.U = nil
	}
	if style.Font.Strike {
		xFont.Strike = &xlsxVal{}
	} else {
		xFont.Strike = nil
	}
	xPatternFill := xlsxPatternFill{}
	xPatternFill.PatternType = style.Fill.PatternType
	xPatternFill.FgColor.RGB = style.Fill.FgColor
	xPatternFill.BgColor.RGB = style.Fill.BgColor
	xFill.PatternFill = xPatternFill
	xBorder.Left = xlsxLine{
		Style: style.Border.Left,
		Color: xlsxColor{RGB: style.Border.LeftColor},
	}
	xBorder.Right = xlsxLine{
		Style: style.Border.Right,
		Color: xlsxColor{RGB: style.Border.RightColor},
	}
	xBorder.Top = xlsxLine{
		Style: style.Border.Top,
		Color: xlsxColor{RGB: style.Border.TopColor},
	}
	xBorder.Bottom = xlsxLine{
		Style: style.Border.Bottom,
		Color: xlsxColor{RGB: style.Border.BottomColor},
	}
	xCellXf = makeXLSXCellElement()
	xCellXf.ApplyBorder = style.ApplyBorder
	xCellXf.ApplyFill = style.ApplyFill
	xCellXf.ApplyFont = style.ApplyFont
	xCellXf.ApplyAlignment = style.ApplyAlignment
	if style.NamedStyleIndex != nil {
		xCellXf.XfId = style.NamedStyleIndex
	}
	return
}

func makeXLSXCellElement() (xCellXf xlsxXf) {
	xCellXf.NumFmtId = 0
	return
}

// Border is a high level structure intended to provide user access to
// the contents of Border Style within an Sheet.
type Border struct {
	Left        string
	LeftColor   string
	Right       string
	RightColor  string
	Top         string
	TopColor    string
	Bottom      string
	BottomColor string
}

func NewBorder(left, right, top, bottom string) *Border {
	return &Border{
		Left:   left,
		Right:  right,
		Top:    top,
		Bottom: bottom,
	}
}

// Fill is a high level structure intended to provide user access to
// the contents of background and foreground color index within an Sheet.
type Fill struct {
	PatternType string
	BgColor     string
	FgColor     string
}

func NewFill(patternType, fgColor, bgColor string) *Fill {
	return &Fill{
		PatternType: patternType,
		FgColor:     fgColor,
		BgColor:     bgColor,
	}
}

type Font struct {
	Size      float64
	Name      string
	Family    int
	Charset   int
	Color     string
	Bold      bool
	Italic    bool
	Underline bool
	Strike    bool
}

func NewFont(size float64, name string) *Font {
	return &Font{Size: size, Name: name}
}

type Alignment struct {
	Horizontal   string
	Indent       int
	ShrinkToFit  bool
	TextRotation int
	Vertical     string
	WrapText     bool
}

var defaultFontSize = 12.0
var defaultFontName = "Verdana"

func SetDefaultFont(size float64, name string) {
	defaultFontSize = size
	defaultFontName = name
}

func DefaultFont() *Font {
	return NewFont(defaultFontSize, defaultFontName)
}

func DefaultFill() *Fill {
	return NewFill("none", "", "")

}

func DefaultBorder() *Border {
	return NewBorder("none", "none", "none", "none")
}

func DefaultAlignment() *Alignment {
	return &Alignment{
		Horizontal: "general",
		Vertical:   "bottom",
	}
}
