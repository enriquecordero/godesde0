package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var numero1 int 
var err error
func Ejercicio2() {


	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un numero  :")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text()) //conv de string a int
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i:= 1; i <=10 ; i++{
		fmt.Printf("%d x %d = %d \n" , numero1, i, i*numero1)
			
	}

}
