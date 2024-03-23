package abstract

type IDisplay interface {
	Open()
	Print()
	Close()
}

func NewAbstractDisplay(display IDisplay) *AbstractDisplay {
	return &AbstractDisplay{
		display,
	}
}

type AbstractDisplay struct {
	display IDisplay
}

func (d *AbstractDisplay) Display() {
	d.display.Open()

	d.display.Print()
	d.display.Print()
	d.display.Print()

	d.display.Close()
}
