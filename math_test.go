package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(16, 15)

	if total != 31 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtrai(t *testing.T) {

	total := subtrai(16, 15)

	if total != 1 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}
