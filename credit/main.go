package main
import (
	"fmt"
	credit "github.com/lizaMosiyash/bgo-1_homeworks_2/pkg"
)

func main() {
	var monthlyPayment,totalSumOfCredit, sumOfPercent = credit.Calculate(1_000_000_00, 36, 20)
	fmt.Println(monthlyPayment)
	fmt.Println(totalSumOfCredit)
	fmt.Println(sumOfPercent)
}
