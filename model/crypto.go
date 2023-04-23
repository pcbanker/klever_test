package model

import (
	"github.com/google/uuid"
)

type CryptoCoin struct {
	ID     string `json:"_id" bson:"_id, omitempty"`
	Name   string `json:"name" bson:"name"`
	Symbol string `json:"symbol" bson:"symbol"`
	Vote   int64  `json:"vote" bson:"vote"`
}

type Client struct {
	UserId uuid.UUID `json:"_id" bson:"_id"`
	Crypto string
	Vote   int64 `json:"vote" bson:"vote"`
}

type Price struct {
	Price float64
}

type UpVote struct {
	Upvote string `json:"upvote"`
}
