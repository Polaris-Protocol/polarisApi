package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectPost struct {
}

type User struct {
	Id       primitive.ObjectID `json:"id"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Title    string             `json:"title,omitempty" validate:"required"`
}

type Users struct {
	Disc_id        uint64 `json:"disc_id"`
	Balance        uint64 `json:"balance"`
	Eth_wallet     string `json:"eth_wallet"`
	Polygon_wallet string `json:"polygon_wallet"`
	Bnb_wallet     string `json:"bnb_wallet"`
	Wax_wallet     string `json:"wax_wallet"`
	Key_hash       string `json:"key_hash"`
}

type Projects struct {
	Remaining_giveaways uint16 `json:"remaining_giveaways"`
	Owner_disc_id       uint64 `json:"owner_disc_id"`
	Tier                uint8  `json:"tier"`
	Balance             uint64 `json:"balance"`
	Server_id           uint64 `json:"server_id"`
	Chain               string `json:"chain"`
}

type Roles struct {
	Server_id       uint64 `json:"server_id" validate:"required"`
	Role_id         uint64 `json:"role_id" validate:"required"`
	Wax_template_id uint64 `json:"wax_template_id"`
	ERC721Contract  string `json:"erc721_contract"`
	ERC721Amount    uint64 `json:"erc721_amount"`
	ERC20Contract   string `json:"erc20_contract"`
	ERC20Amount     string `json:"erc20_amount"`
}

type Orders struct {
	Order_id  uint64 `json:"order_id"`
	Status    string `json:"status"`
	Type      uint64 `json:"type"`
	Amount    string `json:"amount"`
	To        string `json:"to"`
	From      string `json:"from"`
	Valid_1   bool   `json:"valid_1"`
	Valid_2   bool   `json:"valid_2"`
	Valid_3   bool   `json:"valid_3"`
	Chain     string `json:"chain"`
	Timestamp uint64 `json:"timestamp"`
}

/*

order types

0 giveaway
1 deposit
2 withdraw

*/
