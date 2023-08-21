Como executar

1. Clonar o repositório
git clone https://github.com/agneto/20-CleanArch

Entrar na raiz do projeto via prompt

2. Criar um diretório na raiz com o seguinte nome: .docker a fim de não perder os dados quando derrubar os containers
mkdir .docker

3. Subir o banco de dados MySql e Rabbitmq com Docker
docker compose up -d

4. Executar as migrações do MySQL com o 'Golang-migrate' para criar a tabela:
make migrate

5. Entrar no diretorio cmd/ordersystem.
go run main.go wire_gen.go



No graphQL existe uma mutation e uma query:
mutation createOrder {
  createOrder(input: {id: "ddd", Price: 14.4, Tax: 1.5}) {
    id
    Price
    Tax
    FinalPrice
  }
}

query listOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}

no web:
POST http://localhost:8000/order HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "id":"a",
    "price": 100.5,
    "tax": 0.5
}


GET http://localhost:8000/list HTTP/1.1

No gRPC
execute o seguinte comando:
evans -r repl

digite para criar uma order:
call CreateOrder

digite para listar as orders:
call ListOrders