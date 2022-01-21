package builder

type house struct {
	window string
	door   string
	floor  string
}

type Builder interface {
	setWindow()
	setDoor()
	setFloor()
	getHouse() house
}

type WoodBuilder struct {
	window string
	door   string
	floor  string
}

func (b *WoodBuilder) setWindow() {
	b.window = "Wooden Window"
}

func (b *WoodBuilder) setDoor() {
	b.door = "Wooden Door"
}

func (b *WoodBuilder) setFloor() {
	b.floor = "Wooden floor"
}

func (b *WoodBuilder) getHouse() house {
	return house{
		window: b.window,
		door:   b.door,
		floor:  b.floor,
	}
}

type StoneBuilder struct {
	window string
	door   string
	floor  string
}

func (b *StoneBuilder) setWindow() {
	b.window = "Stone Window"
}

func (b *StoneBuilder) setDoor() {
	b.door = "Stone Door"
}

func (b *StoneBuilder) setFloor() {
	b.floor = "Stone floor"
}

func (b *StoneBuilder) getHouse() house {
	return house{
		window: b.window,
		door:   b.door,
		floor:  b.floor,
	}
}
