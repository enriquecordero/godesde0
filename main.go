package main

import (
	"fmt"

	"github.com/enriquecordero/godesde0/variable"
)

func main() {
	// Enteros 
	variable.MostrarEnteros()
	// Resto	
	fmt.Println("****************************************************************")
	variable.RestoVariables()
	fmt.Println("****************************************************************")

	estado ,textoaNum := variable.ConviertoaTexto(123)
	fmt.Println("estado: ",estado)
	fmt.Println("textoaNum : ",textoaNum)
}

