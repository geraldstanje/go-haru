// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go bindings for libharu.
//
// See http://libharu.org/
package haru

type Doc struct {
	doc hpdfDoc
}

type Page struct {
	page hpdfPage
}

type Font struct {
	font hpdfFont
}

func NewDoc() *Doc {
	return &Doc{
		doc: hpdfNew(),
	}
}

func (p *Doc) Free() {
	hpdfFree(p.doc)
	p.doc = nil
}

func (p *Doc) AddPage() *Page {
	page := hpdfAddPage(p.doc)
	return &Page{page}
}

func (p *Doc) GetFont(fontName, encodingName string) *Font {
	t := hpdfGetFont(p.doc, fontName, encodingName)
	return &Font{t}
}

func (p *Doc) SaveToFile(filename string) {
	hpdfSaveToFile(p.doc, filename)
}

func (p *Page) GetWidth() float32 {
	return hpdfPageGetWidth(p.page)
}

func (p *Page) GetHeight() float32 {
	return hpdfPageGetHeight(p.page)
}

func (p *Page) SetLineWidth(lineWidth float32) {
	hpdfPageSetLineWidth(p.page, lineWidth)
}

func (p *Page) Rectangle(x, y, dx, dy float32) {
	hpdfPageRectangle(p.page, x, y, dx, dy)
}

func (p *Page) Stroke() {
	hpdfPageStroke(p.page)
}

func (p *Page) SetFontAndSize(font *Font, size float32) {
	hpdfPageSetFontAndSize(p.page, font.font, size)
}

func (p *Page) TextWidth(text string) float32 {
	return hpdfPageTextWidth(p.page, text)
}

func (p *Page) BeginTextEnd(fn func()) {
	hpdfPageBeginText(p.page)
	defer hpdfPageEndText(p.page)
	fn()
}

func (p *Page) TextOut(xpos, ypos float32, text string) {
	hpdfPageTextOut(p.page, xpos, ypos, text)
}

func (p *Page) MoveTextPos(x, y float32) {
	hpdfPageMoveTextPos(p.page, x, y)
}

func (p *Page) ShowText(text string) {
	hpdfPageShowText(p.page, text)
}

func GetVersion() string {
	return hpdfVersion
}
