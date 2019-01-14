package Picture_generate

import (
	"flag"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"image/jpeg"

	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "/usr/share/fonts/truetype/ubuntu/Ubuntu-M.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 60, "font size in points")
	spacing  = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

var Text []string

func Picture_generatd() {

	backGroundColor := image.White
	backgroundWidth := 350
	backgroundHeight := 350
	rgba := image.NewRGBA(image.Rect(0, 0, backgroundWidth, backgroundHeight))
	draw.Draw(rgba, rgba.Bounds(), backGroundColor, image.ZP, draw.Src)

	fmt.Printf("dfddf")

	image2, err := os.Open("/root/go/src/awesomeProject1/Picture_generate/ico_sun.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	second, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	second1 := resize.Resize(190, 100, second, resize.Bilinear)

	offset := image.Pt(75, 55)
	bi := rgba.Bounds()

	image3 := image.NewRGBA(bi)

	flag.Parse()

	// Read the font data.
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	fg := image.Black
	d := &font.Drawer{
		Dst: rgba,
		Src: fg,
		Face: truetype.NewFace(f, &truetype.Options{
			Size: *size,
			DPI:  *dpi,
		}),
	}
	y := int(math.Ceil(*size * *dpi / 30))
	dy := int(math.Ceil(*size * *spacing * *dpi / 80))
	y += dy
	for _, s := range Text {
		d.Dot = fixed.P(110, y)
		d.DrawString(s)
		y += dy
	}

	draw.Draw(image3, bi, rgba, image.ZP, draw.Src)
	draw.Draw(image3, second1.Bounds().Add(offset), second1, image.ZP, draw.Over)

	third, err := os.Create("result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
