package product

const (
	duration = 12
)

func Calculate(amount int) int {
	return amount / duration
}

//func CalculateMonthlyPayment(price int64) int64 {
//	return price / 12
//}
