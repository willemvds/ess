package cart

type command interface{}

type CreateCommand struct {
	CartId string
}

type AddItemCommand struct {
	CartId string
	ItemId string
	Qty    int
}
