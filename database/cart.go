package database

import (
	"errors"
)

var (
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProduct = errors.New("can't find the product ")
	ErrUserIdisNotValid = errors.New("this user is not vaild ")
	ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrcantRemoveItemCart = errors.New("cannot remove this item from cart")
	ErrCantGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyitemFromCart(){

}

func InstantBuyer(){

}