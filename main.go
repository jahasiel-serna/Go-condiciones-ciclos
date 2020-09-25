package main

import "fmt"

func main() {
	var dia, mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	if dia < 1 || dia > 31 || mes < 1 || mes > 12 {
		fmt.Println("error en la entrada")
	} else {
		switch {
		case (mes == 1 && dia >= 20) || (mes == 2 && dia <= 18):
			fmt.Println("acuario")
		case (mes == 2 && dia >= 19) || (mes == 3 && dia <= 20):
			fmt.Println("piscis")
		case (mes == 3 && dia >= 21) || (mes == 4 && dia <= 19):
			fmt.Println("aries")
		case (mes == 4 && dia >= 20) || (mes == 5 && dia <= 20):
			fmt.Println("tauro")
		case (mes == 5 && dia >= 21) || (mes == 6 && dia <= 20):
			fmt.Println("géminis")
		case (mes == 6 && dia >= 21) || (mes == 7 && dia <= 22):
			fmt.Println("cáncer")
		case (mes == 7 && dia >= 23) || (mes == 8 && dia <= 22):
			fmt.Println("leo")
		case (mes == 8 && dia >= 23) || (mes == 9 && dia <= 22):
			fmt.Println("virgo")
		case (mes == 9 && dia >= 23) || (mes == 10 && dia <= 22):
			fmt.Println("libra")
		case (mes == 10 && dia >= 23) || (mes == 11 && dia <= 21):
			fmt.Println("escorpio")
		case (mes == 11 && dia >= 22) || (mes == 12 && dia <= 21):
			fmt.Println("sagitario")
		case (mes == 12 && dia >= 22) || (mes == 1 && dia <= 19):
			fmt.Println("capricornio")
		}
	}
}
