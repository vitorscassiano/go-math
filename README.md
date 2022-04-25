# Calculadora

O intuito dessa sessão é que vocês saiam sabendo o básico da linguagem e que desperte um iteresse maior pra construir outros projetos, melhorar o conhecimento dessa tecnologia e até vontade de aprender outras tecnologias, não necessariamente Golang.

Pode parecer operações simples, mas elas vão ajudar bastante no entendimento básico da linguagem Go, assim como sua estruturas de dados, declaração de funções, tipagem e comandos básicos da tecnologia.

## Desafio

O desafio pra esse primeiro dojo, nós vamos construir uma linha de comando que realiza algumas operações básicas de uma calculadora, essas operações são:
- Adição
- Subtração
- Multiplicação
- Divisão


## Projeto

### Execução

```shell
> ./programa add 5 4
# Output: 9

> ./programa sub 5 4
# Output: 1

> ./programa mut 5 4
# Output: 40

> ./programa div 5 4
# Output: 1.25
```

### Estrutura de pastas

```shell
.
├── README.md
├── main.go
└── pkg
    └── math
        ├── add.go
        └── add_test.go
```

### Casos de uso

[ ] Usuário soma dois números.
[ ] Usuário subtrai dois números.
[ ] Usuário multiplica dois números.
[ ] Usuário divide dois números.

### Features bonus
[ ] Usuário soma/subtrai/multiplica/divide uma lista de números.
[ ] Usuário pode usar os sinais '+, -, * e /' para realizar as operações.
[ ] Usuário pode realizar operações em números decimais.
