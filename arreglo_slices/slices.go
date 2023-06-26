package arregloslices

import (
  "fmt"
)

var tablaS []int = []int {1,2,11,55,66}
var arreglo [10]int = [10]int {6,87,243,45,676,43,12,1,90,34}

func MuestroSlice(){
fmt.Println(tablaS)

porcion := arreglo[3:] // slice creado , desde la posicion 3 hastala ultima
porcion2 := arreglo[:5] //slice creado , desde la posicion 0 a la 5 
porcion3 := arreglo[6:7] //slice creado , desde la posicion 6 a la 7 

fmt.Println(porcion)
fmt.Println(porcion2)
fmt.Println(porcion3)
}

func Capacidad (){
	elementos := make([]int,5,20) // largo 5 pero 20 capacidad
	fmt.Printf("Largo %d, Capacidad %d",len(elementos),cap(elementos)) // cap capacidad
}