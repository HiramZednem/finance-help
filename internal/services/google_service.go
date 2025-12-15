package services

var GoogleService interface {
	GetBalance()
	GetCategories()
	AddSpend(amount float32, description string, category string)
}

var GoogleServiceImpl struct {

}
