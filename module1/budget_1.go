package module1

// Budget stores budget information
type Budget struct {
	Items []Item
	Max   float32
}

// Item stores item information
type Item struct {
	Description string
	Price       float32
}
