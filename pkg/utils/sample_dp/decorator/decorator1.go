package main

import "fmt"

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	Show(display display) string
}

type defaultDisplay struct {
	display
}

func (self *defaultDisplay) Show(display display) string {
	str := ""
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i)
	}
	return str
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{&defaultDisplay{}, str}
}

type StringDisplay struct {
	*defaultDisplay
	str string
}

func (self *StringDisplay) getRows() int {
	return 1
}

func (self *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return self.str
	} else {
		return ""
	}
}

func NewSideBorder(display display, borderChar string) *SideBorder {
	return &SideBorder{&border{display: display}, borderChar}
}

type border struct {
	*defaultDisplay
	display display
}

type SideBorder struct {
	*border
	borderChar string
}

func (self *SideBorder) getColumns() int {
	return len(self.borderChar)*2 + self.display.getColumns()
}

func (self *SideBorder) getRows() int {
	return self.display.getRows()
}

func (self *SideBorder) getRowText(row int) string {
	return self.borderChar + self.display.getRowText(row) + self.borderChar
}

func main() {
	d1 := NewStringDisplay("うえーい")
	result1 := d1.Show(d1)
	fmt.Println(result1)

	d2 := NewSideBorder(d1, "---")
	result2 := d2.Show(d2)
	fmt.Println(result2)
}
