package bridge

// this pattern allows splitting a large class into two separate hierarchies: abstraction and implementation

// abstraction (also called interface) -> high-level control layer for some entity
// this layer isn’t supposed to do any real work on its own.

// implementation (also called platform) -> the real work is delegated to this layer

// the abstraction contains a reference to the implementation

// example:
// abstraction -> computers, such as mac or windows
// implementation -> printers, such as hp or epson

type computer interface {
	print()
	setPrinter(printer)
}

type printer interface {
	printFile()
}

