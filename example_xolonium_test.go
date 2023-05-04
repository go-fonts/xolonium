// Copyright ©2020 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xolonium_test

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
