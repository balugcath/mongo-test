package csv2rec

import (
	"strconv"

	"github.com/balugcath/mongo-test/pkg/types"
)

// Array2Records ...
func Array2Records(a [][]string) ([]types.ProductPrice, error) {
	var (
		err error
		r   []types.ProductPrice
	)

	for i := range a {
		pp := types.ProductPrice{ProductName: a[i][0]}

		pp.Price, err = strconv.ParseFloat(a[i][1], 64)
		if err != nil {
			return nil, err
		}
		r = append(r, pp)
	}
	return r, nil
}
