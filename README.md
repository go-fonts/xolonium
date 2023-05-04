# xolonium

[![GitHub release](https://img.shields.io/github/release/go-fonts/xolonium.svg)](https://github.com/go-fonts/xolonium/releases)
[![GoDoc](https://godoc.org/github.com/go-fonts/xolonium?status.svg)](https://godoc.org/github.com/go-fonts/xolonium)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://github.com/go-fonts/xolonium/raw/master/LICENSE)

`xolonium` provides the [xolonium](https://fontlibrary.org/en/font/xolonium) fonts as importable Go packages.

The fonts are released under the [SIL Open Font](https://github.com/go-fonts/xolonium/raw/master/LICENSE-SIL) license.
The Go packages under the [BSD-3](https://github.com/go-fonts/xolonium/raw/master/LICENSE) license.

## Example

```go
import (
	"fmt"
	"log"

	"github.com/go-fonts/xolonium/xoloniumregular"
	"golang.org/x/image/font/sfnt"
)

func Example() {
	ttf, err := sfnt.Parse(xoloniumregular.TTF)
	if err != nil {
		log.Fatalf("could not parse xolonium font: %+v", err)
	}

	var buf sfnt.Buffer
	v, err := ttf.Name(&buf, sfnt.NameIDVersion)
	if err != nil {
		log.Fatalf("could not retrieve font version: %+v", err)
	}

	fmt.Printf("version:    %q\n", v)
	fmt.Printf("num glyphs: %d\n", ttf.NumGlyphs())

	// Output:
	// version:    "Version 4.1 "
	// num glyphs: 836
}
```
