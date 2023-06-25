package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int 
var numero2 int
var leyenda string
var err error

func IngresoNumeros(){
	fmt.Println("Ingreso numero 1 :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		numero1,err = strconv.Atoi(scanner.Text()) //conv de string a int 
    if err!= nil{
      panic("El dato de ingreso es incorrect"+ err.Error())
    }
	}

	fmt.Println("Ingreso numero 2 :")
	if scanner.Scan(){
		numero2,err = strconv.Atoi(scanner.Text()) //conv de string a int 
    if err!= nil{
      panic("El dato de ingreso es incorrect"+ err.Error())
    }
	}

	fmt.Println("Ingrese leyenda :")
	if scanner.Scan(){
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda,numero1*numero2)

}