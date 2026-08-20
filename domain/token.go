package domain

type Token struct {
	Model
	Name            string      `json:"name"`
	Symbol          string      `json:"symbol"`
	ResourceID      *string     `json:"resourceId"`
	MetadataURL     string      `json:"metadataUrl"`
	ErcStandard     ErcStandard `json:"ercStandard"`
	Decimals        uint8       `json:"decimals"`
	IssuerID        string      `json:"issuerId"`
	Status          TokenStatus `json:"status"`
	ContractAddress string      `json:"contractAddress"`
	TokenClass      string      `json:"tokenClass"`
	TotalSupply     string      `json:"totalSupply"`
	HolderCount     int         `json:"holderCount"`
}
