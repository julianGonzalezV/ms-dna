package service

import (
	"regexp"
	"strings"
)

var (
	minNxn = 4
)

/// sequenceValidation valida que la Matriz de entrada sea correcta, en cuestion de:
/// - Tamanio mímino permitido es de 4x4
/// - Debe ser una matriz NxN, es decir misma cantidad de filas y columnas
/// -
func sequenceValidation(sequence []string) bool {
	result := true

	sequenceLen := len(sequence)
	if sequenceLen < minNxn {
		result = false
	}
	for _, value := range sequence {
		/// Si encuentra una cadena que posee de todo excepto ATCG entonces error
		matched, _ := regexp.MatchString("[^ATCG]", value)
		if matched {
			result = false
			break
		}
		if len(value) != sequenceLen {
			result = false
			break
		}
	}
	return result
}

/// isMutantDna Orquesta todos los llamados para validar de DNA
/// Note que se utiliza un OR operator, en caso de que se cumpla uno solo de una retorna el valor
/// sin tener que ir a los dermás caminos, esto también ayuda al algoritmo
func isAMutantDna(arr []string) bool {
	mutantValid := false
	if searchByRowsAndColumns(arr) || searchByLeftDiagonal(arr) || searchByRigthDiagonal(arr) {
		mutantValid = true
	}
	return mutantValid
}

/// searchByRowColumn recorre filas y columnas para evualuar en cada iteracion si existe un dna mutante
//  Note que se recorren filas y columnas a la misma vez con el fin de optimizar el algoritmo
func searchByRowsAndColumns(arr []string) bool {
	mutantExists := false
	rowString := ""
	columnString := ""
	rigthDiagAux := len(arr) - 1
	colAux := 0
	for row := 0; row < len(arr); row++ {
		rowString = ""
		columnString = ""
		for column := 0; column < len(arr); column++ {
			rowString += string(arr[row][column])
			columnString += string(arr[column][row])
		}
		if mutantDna(rowString) {
			mutantExists = true
			break
		}
		colAux++
		rigthDiagAux--
	}

	return mutantExists

}

/// searchByLeftDiagonal valida la existencia de un Mutante en la diagonal  con esta inclinacion => \
func searchByLeftDiagonal(arr []string) bool {
	mutantExists := false
	leftDiagonalString := ""
	rowAux := 0
	column := 0
	for diag := 1 - len(arr); diag < len(arr)-1; diag++ {
		rowAux = 0
		column = 0
		leftDiagonalString = ""
		if diag > 0 {
			column = diag
		} else {
			rowAux = -diag
		}
		for row := rowAux; row < len(arr) && column < len(arr); row++ {
			leftDiagonalString += string(arr[row][column])
			column++
		}
		if mutantDna(leftDiagonalString) {
			mutantExists = true
			break
		}
	}
	return mutantExists
}

/// searchByRigthDiagonal valida la existencia de un Mutante en la diagonal  con esta inclinacion =>  /
func searchByRigthDiagonal(arr []string) bool {
	mutantExists := false
	rightDiagonalString := ""
	rowAux := 0
	column := 0
	for diag := 1 - len(arr); diag < len(arr)-1; diag++ {
		rowAux = diag
		column = 0
		rightDiagonalString = ""
		if diag > 0 {
			column = diag
			rowAux = len(arr) - 1
		} else {
			column = 0
			rowAux = (len(arr) - 1) + diag
		}
		for row := rowAux; row >= 0 && column < len(arr); row-- {
			rightDiagonalString += string(arr[row][column])
			column++
		}
		if mutantDna(rightDiagonalString) {
			mutantExists = true
			break
		}

	}
	return mutantExists
}

/// mutantDna contienen todas las secuencias de caracteres que representan un gen mutante
/// Se decide hacerlo así ya que lo hace más escalable a futuro, imagine si le piden que
/// la regla de negocio cambió y que ya no es la secuencia de 4 letras iguales :O
/// Bajo el escenario anterior acá solamente es meter la nueva secuencia o a futuro poderla consultar
/// de alguna base de datos!!
func mutantDna(dna string) bool {
	mutantDna := false
	correctSequences := []string{"AAAA", "TTTT", "CCCC", "GGGG"}
	for _, seq := range correctSequences {
		/// Si es string que llega a ser evaluado contiene alguna se las secuencias mutantes entonces e
		/// puede trabajar para MAGNETO
		if len(dna) >= len(seq) && strings.Contains(dna, seq) {
			mutantDna = true
			break
		}
	}
	return mutantDna
}
