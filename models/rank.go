package models

//
type Rank struct {
	Game  string
	Items RankItem
}

//
type RankItem struct {
	User string
	Rank int
}
