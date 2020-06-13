package credit
import "math"

func Calculate(sumOfCredit, countOfMonth, percentOfRate float64) (int, int, int) {
	monthlyPercent := percentOfRate/12/100
	partOfAnnCoef := math.Pow((1+monthlyPercent), countOfMonth)
	AnnCoef := monthlyPercent*partOfAnnCoef/(partOfAnnCoef-1)
	monthlyPayment := AnnCoef*sumOfCredit
	//var mon1 int = int(monthlyPayment)
	var monthlyPayment1 = int(monthlyPayment)
	var totalSumOfCredit int = int(monthlyPayment) * int(countOfMonth)
	var sumOfPercent int = totalSumOfCredit - int(sumOfCredit)
	return monthlyPayment1, totalSumOfCredit, sumOfPercent
}
