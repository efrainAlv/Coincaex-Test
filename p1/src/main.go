package main

import(
	"fmt"
	"strconv"
	"math"
)


func hexToInt(hex string) int {

	if n, err := strconv.ParseInt(hex, 16, 32); err == nil {
		return int(n)
	}
	return int(0)
}

func hexToDecimal(hex string) int {

	num := 0
	for i := 0; i < len(hex); i++ {
		n := hexToInt(string(hex[i]))
		num = (n*int(math.Pow(float64(16), float64(len(hex)-1-i)))) + num
	}

	return num
}


func hexToRGB(hex string) []int{

	var rgb []int

	for i := 0; i < len(hex); i+=2 {
		if (i+2)<=(len(hex)) {
			rgb = append(rgb, hexToDecimal(hex[i:i+2]))
		}
	}

	return rgb
}

func rbgToHex(color []int) string {

	hex := ""

	for i := 0; i < len(color); i++ {
		hex += fmt.Sprintf("%X", color[i])
	}

	return hex
}

func avg(color1, color2 []int) []int {

	var avgRGB []int

	for i := 0; i < len(color1); i++ {
		
		if len(color2)-1 >= i {
			avgRGB = append(avgRGB, ((color1[i]+color2[i])/2))
		
		}else {
			avgRGB = append(avgRGB, ((color1[i]+0)/2))
		}
	}

	return avgRGB
}


func main()  {
	
	hexa1 := "7a6736"
	hexa2 := "c2982d"
	
	avgHex := ""

	if len(hexa1)>len(hexa2) {
		avgHex = rbgToHex(avg(hexToRGB("7a6736"), hexToRGB("c2982d")))
	}else {
		avgHex = rbgToHex(avg(hexToRGB("c2982d"), hexToRGB("7a6736")))
	}

	fmt.Println("\n======================================================================")
	fmt.Println("The average color between", hexa1, "and", hexa2, "is", avgHex)
	fmt.Println("======================================================================\n")
}


