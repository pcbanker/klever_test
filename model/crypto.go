package model

type CryptoCoin struct {
	ID     string  `json:"_id" bson:"_id, omitempty"`
	Name   string  `json:"name" bson:"name"`
	Symbol string  `json:"symbol" bson:"symbol"`
	Vote   int64   `json:"vote" bson:"vote"`
	Price  float64 `json:"price" bson:"price"`
}

type Price struct {
	Price float64
}

type UpVote struct {
	Upvote string `json:"upvote"`
}
