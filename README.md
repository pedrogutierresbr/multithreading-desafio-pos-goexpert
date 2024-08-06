# Multithreading

  

## Sobre o projeto

  

Este é o repositório destinado ao desafio Multithreading do curso Pós Goexpert da faculdade FullCycle.

  

## Funcionalidades

  

- O projeto possibilita ao usuário:

  

- Recuperar a resposta mais rápida entre duas API's distintas

  

## Como executar o projeto

  

### Pré-requisitos

  

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:

  

- [Git](https://git-scm.com)

- [VSCode](https://code.visualstudio.com/)

- [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

  

#### Acessando o repositório

  

```bash

# Clone este repositório

$ git clone https://github.com/pedrogutierresbr/multithreading-desafio-pos-goexpert.git

```

  

#### Rodando a aplicação

  

```bash

# Abra o terminal

# Importe os pacotes

$ go mod tidy

# Acesse a pasta do projeto no seu terminal/cmd

$ cd cmd/server/

# Inicie a aplicação

$ go run server.go

# A aplicação será aberta na porta:8080

```

  

#### Fazendo a requisição

  

```bash

# Abra o arquivo cep.http na pasta ./test

# Digite o CEP a ser pesquisado na URL

$ http://localhost:8000/cep/{cep} 

$ http://localhost:8000/cep/38050600 <- exemplo

# Clique em Send Request*

# *Lembre-se de ter intalado a extensão Rest Client no seu VS Code

```

  

## Tecnologias

  

As seguintes ferramentas foram usadas na construção do projeto:

  

- Go

  

## Licença

  

Este projeto esta sobe a licença [MIT](./LICENSE).

  

Feito por Pedro Gutierres [Entre em contato!](https://www.linkedin.com/in/pedrogabrielgutierres/)