package main

import "fmt"

type Printable interface {
	Print()
}

type Scanable interface {
	Scan()
}

type Faxable interface {
	Fax()
}

type PrintScanner interface {
	Printable
	Scanable
}

type MultiFunctionDevice interface {
	Print()
	Scan()
	Fax()
}

type Printer struct{}

func (p Printer) Print() {
	fmt.Println("Printing...")
}

type Scanner struct{}

func (s Scanner) Scan() {
	fmt.Println("Scanning...")
}

type FaxMachine struct{}

func (f FaxMachine) Fax() {
	fmt.Println("Faxing...")
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print() {
	fmt.Println("Printing...")
}

func (m MultiFunctionPrinter) Scan() {
	fmt.Println("Scanning...")
}

func (m MultiFunctionPrinter) Fax() {
	fmt.Println("Faxing...")
}

func main() {
	var p Printable
	p = Printer{}
	p.Print()

	var s Scanable
	s = Scanner{}
	s.Scan()

	var f Faxable
	f = FaxMachine{}
	f.Fax()

	var ps PrintScanner
	ps = MultiFunctionPrinter{}
	ps.Print()
	ps.Scan()

	var m MultiFunctionDevice
	m = MultiFunctionPrinter{}
	m.Print()
	m.Scan()
	m.Fax()
}
