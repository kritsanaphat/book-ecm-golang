package controllers

import "errors"

var (
	ErrCantFindProduct    = errors.New("ErrCantFindProduct  ")
	ErrCantDecodeProducts = errors.New("ErrCantDecodeProducts")
	ErrUserIdNotValid     = errors.New("ErrUserIdNotValid")
	ErrCantUpdateUser     = errors.New("ErrCantUpdateUser ")
	ErrCantRemoveItemCart = errors.New("ErrCantRemoveItemCart")
	ErrCantGetItem        = errors.New("ErrCantGetItem")
	ErrCantBuyCartItem    = errors.New("ErrCantBuyCartItem ")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
