package variable


import (
"fmt"
"time"
"strconv"
)

var Nombre string
var Estado bool 
var Sueldo float32 
var Fecha time.Time 
func RestoVariables() {

	Nombre = "Enrique"
	Estado = false
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)

}

// funcion            entra       devolver
func ConviertoaTexto(numero int)(bool,string){
	texto := strconv.Itoa(numero)
	return true, texto
}
