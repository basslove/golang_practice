package main

import "fmt"

type printer struct {
	name string
}

func (self *printer) GetPrinterName() string {
	return self.name
}

func (self *printer) SetPrinterName(name string) {
	self.name = name
}

func (self *printer) Print(str string) string {
	return self.name + ":" + str
}

type PrinterProxy struct {
	Name string
	real *printer
}

func (self *PrinterProxy) GetPrinterName() string {
	return self.Name
}

func (self *PrinterProxy) SetPrinterName(name string) {
	if self.real != nil {
		self.real.SetPrinterName(name)
	}
	self.Name = name
}

func (self *PrinterProxy) Print(str string) string {
	self.realize()
	return self.real.Print(str)
}

func (self *PrinterProxy) realize() {
	if self.real == nil {
		self.real = &printer{self.Name}
	}
}

func main() {
	proxy := PrinterProxy{Name: "A"}
	name := proxy.GetPrinterName()
	fmt.Println(name)

	proxy.SetPrinterName("B")
	name = proxy.GetPrinterName()
	fmt.Println(name)

	result := proxy.Print("C")
	fmt.Println(result)
}
