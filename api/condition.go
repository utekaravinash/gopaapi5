package api

type Condition string

const (
	Any         Condition = "Any"
	New         Condition = "New"
	Used        Condition = "Used"
	Collectible Condition = "Collectible"
	Refurbished Condition = "Refurbished"
)
