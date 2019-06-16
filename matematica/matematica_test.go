package matematica

// AULA82 - GERANDO RELATORIOS COM COBERTURA DE TESTES

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

/*
$ go test
PASS
ok      github.com/rodrigoaraujoti/go_area/matematica   0.194s

$ go test -v
=== RUN   TestMedia
--- PASS: TestMedia (0.00s)
PASS
ok      github.com/rodrigoaraujoti/go_area/matematica   0.212s

$ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/rodrigoaraujoti/go_area/matematica   0.150s

$ go test -coverprofile=resultado.out
PASS
coverage: 100.0% of statements
ok      github.com/rodrigoaraujoti/go_area/matematica   0.202s

$ go tool cover -func=resultado.out
github.com/rodrigoaraujoti/go_area/matematica/matematica.go:11: Media           100.0%
total:                                                          (statements)    100.0%

$ go tool cover -html=resultado.out
>>>> abre no browser um html

 */