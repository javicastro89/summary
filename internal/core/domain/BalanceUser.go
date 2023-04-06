package domain

type BalanceUser struct {
	EmailUser          string
	MonthTransactions  [12]int
	MonthCreditBalance [12]float64
	MonthDebitBalance  [12]float64
	TotalCreditBalance float64
	TotalDebitBalance  float64
	TotalBalance       float64
}
