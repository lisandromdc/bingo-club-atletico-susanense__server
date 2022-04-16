package model

import (
	"gorm.io/gorm"
)

type Delegation struct {
	gorm.Model
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Notes string `json:"notes"`
}

type Bingo struct {
	gorm.Model
	// LotteryDate  time.Time `json:"lotteryDate"`
	LotteryDate  string `json:"lotteryDate"`
	WinnerNumber int    `json:"winnerNumber"`
}

type Sale struct {
	gorm.Model
	BingoNumber  int    `json:"bingoNumber"`
	BuyerName    string `json:"buyerName"`
	BuyerAddress string `json:"buyerAddress"`
	BuyerPhone   string `json:"buyerPhone"`
	BingoId      int    `json:"bingoId"`
	Bingo        Bingo
	SellerId     int `json:"sellerId"`
	Seller       Seller
}

type SalePayment struct {
	gorm.Model
	BingoId      int `json:"bingoId"`
	Bingo        Bingo
	DelegationId int `json:"delegationId"`
	Delegation   Delegation
}

type Seller struct {
	gorm.Model
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Email             string `json:"email" gorm:"uniqueIndex"`
	Password          string `json:"password"`
	IsRoot            bool   `json:"isRoot"`
	IsDelegationAdmin bool   `json:"isDelegationAdmin"`
	CreatorSellerId   int    `json:"creatorSellerId"`
	DelegationId      int    `json:"delegationId"`
	Delegation        Delegation
}
