package main

import "fmt"

func main() {
	var (
		alunos int = 0
		prova, provas, lista, listas, trabalho, presenca float64 = 0, 0, 0, 0, 0, 0
		notaFinal [999999999]float64
		matricula [999999999]int
		situacao  [9999999]string
	)

	const qtdProvas = 8
	const qtdListas = 5
	
	for i := 0; presenca != -1; i++ {
		provas, listas = 0, 0
		fmt.Scan(&matricula[i])
		
		for i := 0; i < qtdProvas; i++ {
			fmt.Scan(&prova)
			provas += prova
		}

		for i := 0; i < qtdListas; i++ {
			fmt.Scan(&lista)
			listas += lista
		}

		fmt.Scan(&trabalho)

		fmt.Scan(&presenca)

		notaFinal[i] = (.7 * (provas / qtdProvas) + (.15 * (listas / qtdListas) + (.15 * trabalho)))
		
		if notaFinal[i] >= 6 && presenca >= 96{
			situacao[i] = "APROVADO"
		} else if notaFinal[i] >= 6 && presenca < 96 {
			situacao[i] = "REPROVADO POR FREQUENCIA"
		} else if notaFinal[i] < 6 && presenca >= 96 {
			situacao[i] = "REPROVADO POR NOTA"
		} else {
			situacao[i] = "REPROVADO POR NOTA E POR FREQUENCIA"
		}
		alunos++
	}
	
	for j := 0; j < alunos - 1; j++ {
		fmt.Printf("Matricula: %d, Nota Final: %.2f, Situacao Final: %s \n", matricula[j], notaFinal[j], situacao[j])
	}

}