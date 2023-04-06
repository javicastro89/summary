package service

import (
	"context"
	"github.com/summary/internal/core/domain"
	"github.com/summary/internal/core/port/inbound"
	"github.com/summary/internal/core/port/outbound"
	"log"
	"strconv"
	"strings"
)

type service struct {
	dispatch   outbound.DispatchMessageService
	repository outbound.RepositoryService
}

func NewService(dispatch outbound.DispatchMessageService, repository outbound.RepositoryService) inbound.ProcessSummaryService {
	return &service{
		dispatch:   dispatch,
		repository: repository,
	}
}

func (s *service) ProcessSummary(ctx context.Context, userSummary []domain.User) error {
	var sum [12][4]float64
	var totalBalance domain.TotalBalance
	for _, v := range userSummary {
		date := strings.Split(v.Date, "/")
		month, errAtoi := strconv.Atoi(date[0])
		if errAtoi != nil {
			log.Fatalf("Error converting string month to int, error: %s", errAtoi.Error())
			return errAtoi
		}
		monthSum(v.Transaction, month-1, &sum, &totalBalance)
	}
	resultMonth, monthTrxQuantity := monthAverageBalance(&sum)

	totalAvg := totalAverageBalance(resultMonth)

	balanceUser := toBalanceUser(monthTrxQuantity, resultMonth, totalAvg, totalBalance, userSummary[0].Id)

	err := s.dispatch.DispatchMessage(ctx, balanceUser)
	if err != nil {
		log.Fatalf("Erro sending mail. Error: %s", err.Error())
		return err
	}

	return s.repository.SaveItem(ctx, balanceUser)
}

func monthSum(amount float64, month int, average *[12][4]float64, total *domain.TotalBalance) {
	if amount > 0 {
		// Credit transaction amount sum and count transaction
		average[month][0] += amount
		average[month][1]++
	} else {
		// Debit Transaction amount sum and count transaction
		average[month][2] += amount
		average[month][3]++
	}
	total.TotalBalance += amount
}

func monthAverageBalance(sum *[12][4]float64) ([12][2]float64, [12]int) {
	// 12 months two average ([0]credit, [1]debit, [2]transactions)
	var average [12][2]float64
	var monthTrxQuantity [12]int
	for i, v := range sum {
		// Credit
		if v[1] != 0 {
			average[i][0] = v[0] / v[1]
		} else {
			average[i][0] = 0
		}
		// Debit
		if v[3] != 0 {
			average[i][1] = v[2] / v[3]
		} else {
			average[i][1] = 0
		}
		// Count month transactions
		monthTrxQuantity[i] = int(v[1] + v[3])
	}
	return average, monthTrxQuantity
}

// totalAverageBalance To calculate the total average of credit and debit.
func totalAverageBalance(monthAverage [12][2]float64) domain.TotalBalanceAvg {
	var totalAvg domain.TotalBalanceAvg
	var balanceAvg domain.BalanceAverage
	for _, v := range monthAverage {
		if v[0] != 0 {

			balanceAvg.CreditAverage += v[0]
			balanceAvg.CreditCount++
		}
		if v[1] != 0 {

			balanceAvg.DebitAverage += v[1]
			balanceAvg.DebitCount++
		}
	}
	if balanceAvg.CreditCount != 0 {
		totalAvg.Credit = balanceAvg.CreditAverage / balanceAvg.CreditCount
	}
	if balanceAvg.DebitCount != 0 {
		totalAvg.Debit = balanceAvg.DebitAverage / balanceAvg.DebitCount
	}

	return totalAvg
}

func toBalanceUser(monthTrxQuantity [12]int, resultMonth [12][2]float64, totalAvg domain.TotalBalanceAvg, totalBalance domain.TotalBalance, email string) domain.BalanceUser {
	creditBalance, debitBalance := processSliceBalance(resultMonth)
	return domain.BalanceUser{
		EmailUser:          email,
		MonthTransactions:  monthTrxQuantity,
		MonthCreditBalance: creditBalance,
		MonthDebitBalance:  debitBalance,
		TotalCreditBalance: totalAvg.Credit,
		TotalDebitBalance:  totalAvg.Debit,
		TotalBalance:       totalBalance.TotalBalance,
	}
}

func processSliceBalance(resultMonth [12][2]float64) ([12]float64, [12]float64) {
	var creditBalance [12]float64
	var debitBalance [12]float64
	for i, v := range resultMonth {
		creditBalance[i] = v[0]
		debitBalance[i] = v[1]
	}
	return creditBalance, debitBalance
}
