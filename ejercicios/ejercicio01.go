package ejercicios

import (
"strconv"
	)

func Ejercicio1( valor string ) (int ,string){
	numero,err := strconv.Atoi(valor)
	if err != nil {
	return 0, "Hubo un error: " + err.Error()	
	}

	if numero > 100 {
		return numero, "mayor a 100"
	}else {
		return numero, "menor a 100"
	}

}