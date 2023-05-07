package interfaces


//continuee
type FavouriteUseCases interface {
	AddToFavourites(productId, userId int) error
	RemoveFromFav(userId, productId int) error
//	ViewFavourite(usersId)([]Response.)
}
