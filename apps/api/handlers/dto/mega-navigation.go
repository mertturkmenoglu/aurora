package dto

type MegaNavigation struct {
	Items []L0Item `json:"items"`
}

type L0Item struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Items []L1Item `json:"items"`
}

type L1Item struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Items []L2Item `json:"items"`
}

type L2Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
