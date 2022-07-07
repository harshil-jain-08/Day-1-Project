package main

import (
	"encoding/json"
	"fmt"
)

//Declaring Struct
type matrix struct {
	rows int
	col  int
	mat  [][]int
}

//get the number of rows
func (m matrix) getrow() int {
	return m.rows
}

//get the number of columns
func (m matrix) getcol() int {
	return m.col
}

//set the elements of the matrix at a given position (i,j)
func (m matrix) setelement(a, i, j int) {
	m.mat[i][j] = a
}

//adding two matrices.
func (m matrix) addmatrices(n matrix) matrix {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.col; j++ {
			m.mat[i][j] = m.mat[i][j] + n.mat[i][j]
		}
	}
	return m
}

// exporting to json
func (m matrix) printjson() {
	jsonmatrix, err := json.Marshal(m.mat)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(jsonmatrix))
	}
}
