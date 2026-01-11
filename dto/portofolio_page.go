package dto

type PortfolioPage struct {
	Personal PersonalResponse
	Activity []ActivityResponse
	Work     []WorkResponse
}
