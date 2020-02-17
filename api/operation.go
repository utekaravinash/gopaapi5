package api

type Operation string

const (
	GetBrowseNodes Operation = "GetBrowseNodes"
	GetItems       Operation = "GetItems"
	GetVariations  Operation = "GetVariations"
	SearchItems    Operation = "SearchItems"
)
