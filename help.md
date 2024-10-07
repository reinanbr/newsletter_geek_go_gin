## 07/10/24, 11:57

instalando o go:

```sh
apt install golang
```

para iniciar o projeto, é preciso primeiro, criar/clonar o repositorio e botar no disco local.<br>
Diferente do Rust, o Go não cria pasta.<br>

1. crie a pasta `pasta_do_app`
2. cd `pasta-do-app`
3. `go mod init _nome_do_app_`

## importando modulos

 para importar um modulo n main do diretorio `pasta1/subspata1/code.go`, em que o package desse codigo seja `package subspata1`, como exemplo:

```go
package subspata1

func Add(a1 float64, a2:float64)(float64){...}
```

 faça no main:

```go
ipackage main

import "_nome_do_app_/pasta1/subspata1"

func main(){
    subspata1.Add(1.1, 3.4)
}
```

Funciona de maneira parecida se for apenas no pasta1:

```go
package pasta1

func Add(a1 float64, a2:float64)(float64){...}
```

 faça no main:                                                                     
```go
ipackage main

import "_nome_do_app_/pasta1"

func main(){
    pasta1.Add(1.1, 3.4)
}
```


 app segue essa estrutura de modulação, seguindo conceitos basicos de template e mvc, um dos melhores que já tive contato, comparando a blueprint.

```sh
geek_history
├── go.mod
├── go.sum
├── help.md
├── main.go
├── src
│   ├── controllers
│   │   └── home_controller.go
│   └── views
│       └── templates
│           ├── base
│           ├── includes
│           │   └── index.html
│           └── layouts
│               └── base.html
└── tools
    └── templates.go
```

