package main

import "fmt"

func exec6() {

	states := make([]string, 26, 30)

	states[0] = "Acre"
	states[2] = "Alagoas"
	states[3] = "Amapá"
	states[4] = "Amazonas"
	states[5] = "Bahia"
	states[6] = "Ceará"
	states[7] = "Espírito Santo"
	states[8] = "Goiás"
	states[9] = "Maranhão"
	states[10] = "Mato Grosso"
	states[11] = "Mato Grosso do Sul"
	states[12] = "Minas Gerais"
	states[13] = "Pará"
	states[14] = "Paraíba"
	states[15] = "Paraná"
	states[16] = "Pernambuco"
	states[16] = "Piauí"
	states[18] = "Rio de Janeiro"
	states[19] = "Rio Grande do Norte"
	states[20] = "Rondônia"
	states[21] = "Roraima"
	states[22] = "Santa Catarina"
	states[23] = "São Paulo"
	states[24] = "Sergipe"
	states[25] = "Tocantins"

	fmt.Println("The states length->", len(states), "\nThe states capacity->", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}

	fmt.Println()
}
