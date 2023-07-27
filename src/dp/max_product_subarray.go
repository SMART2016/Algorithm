package dp

import "math"

func MaxProduct(A []int) int {
	sf := 1
	pf := 1
	maxProd := math.MinInt
	for i := 0; i < len(A); i++ {
		//To handle case when A[i] is 0
		//It divides the array into subaaray ,
		//so pf and sf will be reinitiated to 1
		if pf == 0 {
			pf = 1
		}
		if sf == 0 {
			sf = 1
		}

		//Prefix product from front
		pf = pf * A[i]

		//Prefix product from back
		sf = sf * A[len(A)-i-1]

		maxProd = max(maxProd, max(pf, sf))
	}
	return maxProd
}
