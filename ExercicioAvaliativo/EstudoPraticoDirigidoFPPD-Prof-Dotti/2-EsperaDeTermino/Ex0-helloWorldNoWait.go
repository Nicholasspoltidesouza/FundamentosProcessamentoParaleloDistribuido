// Fernando Dotti - fldotti.github.io - PUCRS - Escola Politécnica
// programa da internet
// EXERCICIOS:
//    1) rode o programa abaixo.
//       o que você conclui sobre a execução observada?
// O programa pode ou não esperar a saída de world, já que está indo concorrentemente com o hello.

package main

import (
	"fmt"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
