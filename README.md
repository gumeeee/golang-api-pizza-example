# Pizzaria API

Este é um projeto de uma API simples para gerenciar um cardápio de pizzas. Ele permite listar, adicionar e buscar pizzas por ID.

## Funcionalidades

- **Listar Pizzas**: Retorna todas as pizzas disponíveis no cardápio.
- **Adicionar Pizza**: Permite adicionar uma nova pizza ao cardápio.
- **Buscar Pizza por ID**: Retorna os detalhes de uma pizza específica com base no ID.

## Estrutura do Projeto

- **`data/pizzas.json`**: Arquivo JSON que armazena os dados das pizzas.
- **`models/pizzaria.go`**: Define o modelo de dados para as pizzas.
- **`main.go`**: Contém a lógica principal da API e as rotas.
- **`requests.http`**: Arquivo para testar as rotas da API.

## Tecnologias Utilizadas

- **Linguagem**: Go (Golang)
- **Framework**: Gin Gonic (https://github.com/gin-gonic/gin)
- **Armazenamento**: Arquivo JSON

## Como Executar

1. Certifique-se de ter o Go instalado na sua máquina.
2. Clone este repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   ```
3. Navegue até o diretório do projeto:
   ```bash
   cd pizzaria
   ```
4. Execute o projeto:
   ```bash
   go run main.go
   ```
5. A API estará disponível em `http://localhost:8080`.

## Rotas Disponíveis

### Listar Pizzas
**GET** `/pizzas`  
Retorna todas as pizzas.

### Adicionar Pizza
**POST** `/pizzas`  
Adiciona uma nova pizza.  
Exemplo de corpo da requisição:
```json
{
    "nome": "Pizza de Calabresa",
    "preco": 35.0
}
```

### Buscar Pizza por ID
**GET** `/pizzas/:id`  
Retorna os detalhes de uma pizza específica.

## Testando a API

Você pode usar o arquivo `requests.http` para testar as rotas diretamente no seu editor de código ou usar ferramentas como Postman e cURL.

## Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests para melhorias.

## Licença

Este projeto está sob a licença MIT.
