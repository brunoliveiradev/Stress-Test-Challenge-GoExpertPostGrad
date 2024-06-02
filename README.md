### Languages: [Português 🇧🇷](#Stress-Test-CLI) | [English 🇨🇦](#stresstest-cli)

---

# Stress Test CLI

Este projeto implementa uma ferramenta de stress test usando Golang para simular múltiplas requisições simultâneas a um
serviço web. Ele permite configurar a quantidade total de requisições e o nível de concorrência, gerando um relatório
detalhado ao final da execução.
Esta API faz parte do desafio técnicos do curso de Pós-Graduação em Engenharia de
Software [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

## ⚙️ Configuração

Você precisará das seguintes tecnologias abaixo:

- [Docker](https://docs.docker.com/get-docker/) 🐳
- [Docker Compose](https://docs.docker.com/compose/install/) 🐳
- [GIT](https://git-scm.com/downloads)

## 🚀 Iniciando

1. Clone o repositório e entre no diretório do projeto.
   ```sh
   git clone https://github.com/brunoliveiradev/Stress-Test-Challenge-GoExpertPostGrad.git
   cd Stress-Test-Challenge-GoExpertPostGrad
   ```

2. Construa a imagem Docker:
   ```sh
   docker build -t stress-test:latest .
   ```

3. Execute o contêiner com a imagem Docker criada, passando os parâmetros
   necessários (`--url`, `--requests`, `--concurrency`):
    ```sh
    docker run stress-test:latest --url=https://fullcycle.com.br/ --requests=100 --concurrency=10
    ```

## 📝 Parâmetros

* `--url`: URL do serviço a ser testado.
* `--requests`: Número total de requests.
* `--concurrency`: Número de chamadas simultâneas.

## 📊 Relatório

Ao final da execução, a ferramenta gera um relatório contendo:

* Tempo total gasto na execução.
* Quantidade total de requests realizados.
* Quantidade de requests com status HTTP 200 (quantidade de Successful requests).
* Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## 🧪 Testes

1. Execute o comando abaixo para rodar os testes unitários:
    ```sh
    go test -v ./...
    ```
2. Caso queira testar localmente sem utilizar o docker, você pode executar o comando diretamente::
    ```sh
    go run cmd/stress-test/main.go --url=https://fullcycle.com.br/ --requests=100 --concurrency=10
    ```

## 🛠️ Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests para melhorias ou correções de bugs.

---

# StressTest CLI

This project implements a stress test tool using Golang to simulate multiple simultaneous requests to a web service. It
allows you to configure the total number of requests and the concurrency level, generating a detailed report at the end
of the execution.
This API is part of the technical challenges of the Postgraduate course in Software
Engineering [GoExpert](https://goexpert.fullcycle.com.br/pos-goexpert/).

## ⚙️ Configuration

You will need the following technologies below:

- [Docker](https://docs.docker.com/get-docker/) 🐳
- [Docker Compose](https://docs.docker.com/compose/install/) 🐳
- [GIT](https://git-scm.com/downloads)
- [Go](https://golang.org/dl/)

## 🚀 Getting Started

1. Clone the repository and enter the project directory.
   ```sh
   git clone https://github.com/brunoliveiradev/Stress-Test-Challenge-GoExpertPostGrad.git
   cd Stress-Test-Challenge-GoExpertPostGrad
   ```
2. Build the Docker image:
   ```sh
   docker build -t stress-test:latest .
   ```
3. Run the container with the created Docker image, passing the necessary
   parameters (`--url`, `--requests`, `--concurrency`):
    ```sh
    docker run stress-test:latest --url=https://fullcycle.com.br/ --requests=100 --concurrency=10
    ```

## 📝 Parameters

* `--url`: URL of the service to be tested.
* `--requests`: Total number of requests.
* `--concurrency`: Number of simultaneous calls.

## 📊 Report

At the end of the execution, the tool generates a report containing:

* Total time spent on execution.
* Total number of requests made.
* Number of requests with HTTP status 200 (number of Successful requests).
* Distribution of other HTTP status codes (such as 404, 500, etc.).

## 🧪 Tests

1. Run the command below to run the unit tests:
    ```sh
    go test -v ./...
    ```

2. If you want to test locally without using docker, you can run the command directly:
    ```sh
    go run cmd/stress-test/main.go --url=https://fullcycle.com.br/ --requests=100 --concurrency=10
    ```

## 🛠️ Contribution

Contributions are welcome! Feel free to open issues or pull requests for improvements or bug fixes.