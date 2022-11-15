# Reuso

Para o reaproveitamento de código, podemos utilizar os modulos como mecanismos para o reuso

## Modulo Local

Neste exemplo utilizo a versão 1.18 do Go.
(versões superiores á 1.16 funciona também)

### GO111MODULE

Primeiro passo é estar dentro da estrutura de exercícios, no terminal executar o "export GO111MODULE=off".
Após isso, se pode tentar fazer o "run ou build".

Caso ele ainda não encontre o modulo área, siga os próximos passos.

### Local

Próximo passo é estar no outro workspace, dentro da raiz aonde ficam os módulos do Go.
Quando você tentou executar o "run ou build", ele vai criar a pasta "src/".

Caso não tenha criado a pasta "area/" e o arquivo "area.go", basta apenas criar, conforme exemplo abaixo.

``` yaml h1_lines="11 12"
package area

import "math"

// Pi é uma proporção numerica definida pela relação entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

```

Volte ao workspace seu projeto e faça "go run *.go"
Pronto

## Referências

- [Go Official](https://go.dev/blog/go116-module-changes)
- [Go111Module Explained](https://maelvls.dev/go111module-everywhere/)