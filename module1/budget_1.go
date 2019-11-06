package module1

type Budget struct {
	Max   float32
	Items []Item
}

type Item struct {
	Description string
	Price       float32
}
