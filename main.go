package main

import "fmt"

func main() {
	var dia, mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	if dia < 1 || dia > 31 || mes < 1 || mes > 12 {
		fmt.Println("error en la entrada")
	} else {
		signo := mes*mes*10 + dia
		switch {
		case signo >= 30 && signo <= 58:
			fmt.Println("acuario")
		case signo >= 59 && signo <= 110:
			fmt.Println("piscis")
		case signo >= 111 && signo <= 179:
			fmt.Println("aries")
		case signo >= 180 && signo <= 270:
			fmt.Println("tauro")
		case signo >= 271 && signo <= 380:
			fmt.Println("géminis")
		case signo >= 381 && signo <= 512:
			fmt.Println("cáncer")
		case signo >= 513 && signo <= 662:
			fmt.Println("leo")
		case signo >= 663 && signo <= 832:
			fmt.Println("virgo")
		case signo >= 833 && signo <= 1022:
			fmt.Println("libra")
		case signo >= 1021 && signo <= 1231:
			fmt.Println("escorpio")
		case signo >= 1232 && signo <= 1461:
			fmt.Println("sagitario")
		case signo >= 1462 && signo <= 29:
			fmt.Println("capricornio")
		}
	}
}
