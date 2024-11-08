package entities

// Item represents a product without shop association
type Item struct {
	// ID helps to uniquely identify this item without string equality hassle.
	ID int
	// Name contains the display name.
	Name      string
	Category  string
	ImagePath string
	// PriceThreshold gives the shopper an idea when to buy this item.
	PriceThreshold int64
}
