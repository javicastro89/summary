package domain

type MonthlyAverage struct {
}

type TotalBalance struct {
	TotalBalance float64
}

type BalanceAverage struct {
	CreditAverage float64
	CreditCount   float64
	DebitAverage  float64
	DebitCount    float64
}

type TotalBalanceAvg struct {
	Credit float64
	Debit  float64
}
