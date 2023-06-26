package arregloslices

import (
  "fmt"
)

var tabla [10]int = [10]int {0,0,11,55,66}

var matrix [20][30]int 

func MuestraArreglos(){
	tabla[0] = 33
	tabla[1] = 44
	tabla2 := [10]string {"Pablo","Juan"}
	//fmt.Println(tabla)
	fmt.Println(tabla2)



	for i:=0; i<len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matrix [2][2] = 15
	fmt.Println(matrix)
}