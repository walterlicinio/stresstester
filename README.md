# stresstester

`stresstester` é uma ferramenta de linha de comando (CLI) escrita em Go para realizar testes de carga em serviços web. Ela permite que os usuários forneçam a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas, gerando um relatório detalhado após a execução dos testes.

## Funcionalidades

- Efetua múltiplas requisições HTTP para a URL especificada.
- Distribui as requisições de acordo com o nível de concorrência definido.
- Garante que o número total de requisições seja cumprido.
- Gera um relatório com:
  - Tempo total gasto na execução.
  - Quantidade total de requisições realizadas.
  - Quantidade de requisições com status HTTP 200.
  - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Requisitos

- Docker

## Como Utilizar

### 1. Build da Imagem Docker

Clone o repositório e navegue até o diretório do projeto:

```sh
git clone https://github.com/walterlicinio/stresstester
cd stresstester
```

Construa a imagem Docker:

```sh
docker build -t stresstester .
```

### 2. Executando o Teste de Carga

Para executar o contêiner e iniciar o teste de carga, utilize o comando:

```sh
docker run --rm stresstester --url=<URL> --requests=<NUM_TOTAL_REQUISICOES> --concurrency=<NUM_CONCORRENCIA>
```

Exemplo:

```sh
docker run --rm stresstester --url=http://google.com --requests=1000 --concurrency=50
```

## Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições a serem feitas.
- `--concurrency`: Número de chamadas simultâneas.

## Relatório Gerado

Após a execução do teste de carga, um relatório será exibido no terminal contendo:

- **Tempo total**: Tempo total gasto na execução do teste.
- **Quantidade total de requisições realizadas**: Contagem de todas as requisições feitas.
- **Quantidade de requisições com status 200**: Número de requisições que receberam o status HTTP 200.
- **Distribuição de códigos de status HTTP**: Contagem de cada código de status retornado pelo servidor.