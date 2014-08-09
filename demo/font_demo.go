// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"

	"github.com/chai2010/go-haru"
)

var fontList = []string{
	"Courier",
	"Courier-Bold",
	"Courier-Oblique",
	"Courier-BoldOblique",
	"Helvetica",
	"Helvetica-Bold",
	"Helvetica-Oblique",
	"Helvetica-BoldOblique",
	"Times-Roman",
	"Times-Bold",
	"Times-Italic",
	"Times-BoldItalic",
	"Symbol",
	"ZapfDingbats",
}

func main() {
	pdf, err := haru.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer pdf.Free()

	// Add a new page object.
	page := pdf.AddPage()
	height := page.GetHeight()
	width := page.GetWidth()

	// Print the lines of the page.
	page.SetLineWidth(1)
	page.Rectangle(page, 50, 50, width-100, height-110)
	page.Stroke()

	// Print the title of the page (with positioning center).
	def_font := pdf.GetFont("Helvetica", nil)
	page.SetFontAndSize(def_font, 24)

	page_title := "Font Demo"
	tw := page.TextWidth(page_title)
	page.BeginTextEnd(func() {
		page.TextOut((width-tw)/2, height-50, page_title)
	})

	// output subtitle.
	page.BeginTextEnd(func() {
		page.SetFontAndSize(def_font, 16)
		page.TextOut(60, height-80, "<Standerd Type1 fonts samples>")
	})

	page.BeginTextEnd(func() {
		page.MoveTextPos(60, height-105)

		for i, _ := range fontList {
			samp_text := "abcdefgABCDEFG12345!#$%&+-@?"
			font := pdf.GetFont(fontList[i], nil)

			// print a label of text
			page.SetFontAndSize(page, def_font, 9)
			page.ShowText(page, fontList[i])
			page.MoveTextPos(page, 0, -18)

			// print a sample text.
			page.SetFontAndSize(font, 20)
			page.ShowText(samp_text)
			page.MoveTextPos(0, -20)
		}
	})

	pdf.Save("font_demo.pdf")
}
