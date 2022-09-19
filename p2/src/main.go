package main

import(
	"fmt"
)

type Variable struct {
    name string
    value  int
}

func multiply(m1, m2 [][]Variable, implicit bool, size int) [][]string {

	var result [][]string

	if !implicit {

		for i := 0; i < size; i++ {

			var col []string
			for j := 0; j < size; j++ {

				exps := ""
				for k := 0; k < size; k++ {
					exps += m1[i][k].name + "*" + m2[k][j].name

					if k < size-1 {
						exps += " + "
					}
				}
				col = append(col, exps)
			}
			result = append(result, col)
		}
	}

	return(result)
}

func main()  {

	m1 := [][]Variable {
		{ Variable{name: "btc"}, Variable{name: "eth"}, },
		{ Variable{name: "trx"}, Variable{name: "usdt"}, },
	}

	m2 := [][]Variable {
		{ Variable{name: "usd"}, Variable{name: "hnl"}, },
		{ Variable{name: "gtq"}, Variable{name: "pen"}, },
	}

	m3 := multiply(m1, m2, false, 2)

	for i := 0; i < 2; i++ {
		fmt.Println(m3[i])	
	}
}


