package main

import (
	"fmt"
	"unsafe"
)

type Pessoa struct {
	Nome  string // 16 bytes
	Idade uint8  // 1 byte
	Nick  string // 16 bytes
} // 33 bytes??? | 8 * 5 = 40 bytes

func main() {
	var bool1 bool
	var numInt16 int16
	var numUint16 uint8
	var realFloat32 float32
	var texto string
	var estrutura1 Pessoa

	arrayBool := [2]bool{true, false}
	arrayInt16 := [2]int16{1500, -32000}
	arrayUint16 := [2]uint16{0, 5000}
	arrayFloat32 := [2]float32{2.5, 1000.50}
	arrayTexto := [2]string{"", "Olá"}
	arrayEstrutura := [2]Pessoa{
		{"Pedro", 10, "Pedrinho"},
		{"Tiago", 5, "Titi"},
	}

	sliceVazioBool := []bool{}
	sliceVazioInt16 := []int16{}
	sliceVazioUint16 := []uint16{}
	sliceVazioFloat32 := []float32{}
	sliceVazioTexto := []string{}
	sliceVazioEstrutura := []Pessoa{}

	sliceBool := []bool{true, false}
	sliceInt16 := []int16{1500, -32000}
	sliceUint16 := []uint16{0, 5000}
	sliceFloat32 := []float32{2.5, 1000.50}
	sliceTexto := []string{"", "Olá"}
	sliceEstrutura := []Pessoa{
		{"Pedro", 10, "Pedrinho"},
		{"Tiago", 5, "Titi"},
	}

	fmt.Println("#### USO DE MÉMORIA (byte = bits / 8) ####")
	imprimir("bool1", unsafe.Sizeof(bool1))
	imprimir("numInt16", unsafe.Sizeof(numInt16))
	imprimir("numUint16", unsafe.Sizeof(numUint16))
	imprimir("realFloat32", unsafe.Sizeof(realFloat32))
	imprimir("texto", unsafe.Sizeof(texto))
	imprimir("estrutura", unsafe.Sizeof(estrutura1))

	fmt.Println("")
	imprimir("arrayBool", unsafe.Sizeof(arrayBool))
	imprimir("arrayInt16", unsafe.Sizeof(arrayInt16))
	imprimir("arrayUint16", unsafe.Sizeof(arrayUint16))
	imprimir("arrayFloat32", unsafe.Sizeof(arrayFloat32))
	imprimir("arrayTexto", unsafe.Sizeof(arrayTexto))
	imprimir("arrayEstrutura", unsafe.Sizeof(arrayEstrutura))
	imprimir("arrayEstrutura0", unsafe.Sizeof(arrayEstrutura[0]))
	imprimir("arrayEstrutura0Nome", unsafe.Sizeof(arrayEstrutura[0].Nome))
	imprimir("arrayEstrutura0Nick", unsafe.Sizeof(arrayEstrutura[0].Nick))
	imprimir("arrayEstrutura0Idade", unsafe.Sizeof(arrayEstrutura[0].Idade))

	fmt.Println("")
	imprimir("sliceVazioBool", unsafe.Sizeof(sliceVazioBool))
	imprimir("sliceVazioInt16", unsafe.Sizeof(sliceVazioInt16))
	imprimir("sliceVazioUint16", unsafe.Sizeof(sliceVazioUint16))
	imprimir("sliceVazioFloat32", unsafe.Sizeof(sliceVazioFloat32))
	imprimir("sliceVazioTexto", unsafe.Sizeof(sliceVazioTexto))
	imprimir("sliceVazioEstrutura", unsafe.Sizeof(sliceVazioEstrutura))

	fmt.Println("")
	imprimir("sliceBool", unsafe.Sizeof(sliceBool))
	imprimir("sliceInt16", unsafe.Sizeof(sliceInt16))
	imprimir("sliceUint16", unsafe.Sizeof(sliceUint16))
	imprimir("sliceFloat32", unsafe.Sizeof(sliceFloat32))
	imprimir("sliceTexto", unsafe.Sizeof(sliceTexto))
	imprimir("sliceEstrutura", unsafe.Sizeof(sliceEstrutura))
	imprimir("sliceEstrutura0", unsafe.Sizeof(sliceEstrutura[0]))
	imprimir("sliceEstrutura0Nome", unsafe.Sizeof(sliceEstrutura[0].Nome))
	imprimir("sliceEstrutura0Nick", unsafe.Sizeof(sliceEstrutura[0].Nick))
	imprimir("sliceEstrutura0Idade", unsafe.Sizeof(sliceEstrutura[0].Idade))
}

func imprimir(nomeVariavel string, totalBytes uintptr) {
	var print string = "byte"
	if totalBytes > 1 {
		print += "s"
	}
	fmt.Printf("%20s\t %02d %s\n", nomeVariavel, totalBytes, print)
}
