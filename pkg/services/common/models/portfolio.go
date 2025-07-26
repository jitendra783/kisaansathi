package models

type DpDetails struct {
	DpID         string
	DpAccountNo  string
	ReinvestFlag string
}

type UnblockedUnits struct {
	NoOfUnits         float64
	FolioCreationDate string
	ReinvestFlag      string
}
