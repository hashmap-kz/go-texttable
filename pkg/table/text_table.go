package table

import (
	"fmt"
	"log"
	"strings"
)

type Align int

const (
	LEFT = iota
	CENTER
	RIGHT
)

const ColumnSeparator = "  "

type TextTableColumn struct {
	Heading  string
	Width    int
	HdAlign  Align
	ColAlign Align
}

type TextTable struct {
	Col    []TextTableColumn
	Row    [][]string
	Curcol int
	Currow int
	Indent int
}

func NewTextTable() *TextTable {
	return &TextTable{
		Col:    []TextTableColumn{},
		Row:    [][]string{},
		Curcol: 0,
		Currow: 0,
		Indent: 0,
	}
}

func (t *TextTable) DefineColumn(heading string, hdAlign Align, colAlign Align) {
	col := TextTableColumn{
		Heading:  heading,
		Width:    len(heading),
		HdAlign:  hdAlign,
		ColAlign: colAlign,
	}
	t.Col = append(t.Col, col)
}

func (t *TextTable) EndRow() {
	t.Curcol = 0
	t.Currow += 1
}

func resize2(r [][]string, m int) [][]string {
	result := make([][]string, len(r)+m)
	for i, elem := range r {
		result[i] = elem
	}
	return result
}

func resize1(r []string, m int) []string {
	result := make([]string, len(r)+m)
	for i, elem := range r {
		result[i] = elem
	}
	return result
}

func (t *TextTable) Insert(item string) {
	if len(t.Row) < t.Currow+1 {
		t.Row = resize2(t.Row, 1)
	}

	if len(t.Row[t.Currow]) < len(t.Col) {
		t.Row[t.Currow] = resize1(t.Row[t.Currow], len(t.Col))
	}

	fine := t.Curcol+1 <= len(t.Col)
	if !fine {
		log.Fatal("inserting more items that defined columns")
	}

	width := len(item)

	// expand column width if necessary
	if width > t.Col[t.Curcol].Width {
		t.Col[t.Curcol].Width = width
	}

	t.Row[t.Currow][t.Curcol] = item
	t.Curcol += 1
}

func (t *TextTable) InsertAll(items ...string) {
	if len(t.Col) < len(items) {
		log.Fatal("you want to insert items for unspecified columns")
	}
	for _, item := range items {
		t.Insert(item)
	}
}

func (t *TextTable) InsertAllAndFinishRow(items ...string) {
	t.InsertAll(items...)
	t.EndRow()
}

func pad(s string, width int, align Align) string {
	lpad := 0
	rpad := 0
	slen := len(s)

	if align == LEFT {
		rpad = width - slen
	}

	if align == CENTER {
		lpad = width/2 - slen/2
		rpad = width - lpad - slen
	}

	if align == RIGHT {
		lpad = width - slen
	}

	return fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", lpad),
		s,
		strings.Repeat(" ", rpad))
}

func (t *TextTable) Print() string {
	result := ""

	// headers
	for i := 0; i < len(t.Col); i++ {
		col := t.Col[i]
		if i > 0 {
			result += ColumnSeparator
		}
		result += fmt.Sprint(pad(col.Heading, col.Width, col.HdAlign))
	}
	result += "\n"

	// rows
	for i := 0; i < len(t.Row); i++ {
		for j := 0; j < len(t.Row[i]); j++ {
			col := t.Col[j]
			if j > 0 {
				result += fmt.Sprint(ColumnSeparator)
			}
			result += fmt.Sprint(pad(t.Row[i][j], col.Width, col.ColAlign))
		}
		result += "\n"
	}

	return result
}
