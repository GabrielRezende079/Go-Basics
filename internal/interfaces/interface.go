package interfaces

import (
	"fmt"
	"math"
)

// Interface para formas
type forma interface {
	area() float64
}

// #----------------------Retangulo-------------------#
type retangulo struct {
	altura, largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// #---------------------------------------------------#

// #----------------------Cirulo-----------------------#
type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	var espoente float64 = 2
	pi := math.Pi
	return pi * math.Pow(c.raio, espoente)
}

// #---------------------------------------------------#

// Apenas escreve a Forma
func escreverArea(f forma) {
	fmt.Printf("A area da forma é %f \n", f.area())
}

// Main
func ExecutarInterfaces() {
	r := retangulo{10, 15}
	c := circulo{5}
	escreverArea(r)
	escreverArea(c)
}
