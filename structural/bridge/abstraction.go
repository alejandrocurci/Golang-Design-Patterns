package bridge

// computers can print, but the actual work is delegated to the printer
// computers have a reference to a printer, allowing different combinations (mac-hp, windows-epson, etc)

type mac struct {
	printer printer
}

func (m *mac) print() {
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
