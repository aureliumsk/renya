package slides

import (
	"github.com/golang/freetype/truetype"
	"renya/resources"
)

const dotsPerInch = 72.0

var headerFace = truetype.NewFace(resources.Font, &truetype.Options{
	Size: 24.0,
	DPI:  dotsPerInch,
})
