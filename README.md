# 📋 Task Manager API

API REST completa para gerenciamento de tarefas desenvolvida em Go, utilizando arquitetura limpa e boas práticas de desenvolvimento.

## 🚀 Tecnologias

- **Go 1.25** - Linguagem de programação
- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **MySQL 8.0** - Banco de dados relacional
- **Docker & Docker Compose** - Containerização

## 🏗️ Arquitetura

O projeto segue a **arquitetura em camadas** (Clean Architecture):

```
task-manager/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── internal/
│   ├── controller/          # Camada de apresentação (handlers HTTP)
│   ├── service/             # Camada de lógica de negócio
│   ├── repository/          # Camada de acesso a dados
│   ├── models/              # Entidades do domínio
│   ├── dto/                 # Data Transfer Objects
│   ├── db/                  # Configuração do banco de dados
│   └── routes/              # Definição de rotas
├── docker-compose.yml
├── go.mod
└── go.sum
```

### Camadas

- **Controller**: Recebe requisições HTTP e retorna respostas
- **Service**: Contém regras de negócio
- **Repository**: Abstrai acesso ao banco de dados
- **Models**: Define estrutura das entidades
- **DTO**: Valida entrada/saída de dados

## 📦 Funcionalidades

A API possui três módulos principais:

- 👤 **Usuários** - Cadastro e gerenciamento de usuários
- 🏷️ **Categorias** - Organização de tarefas por categoria
- ✅ **Tarefas** - CRUD completo de tarefas

## ⚙️ Configuração

### Pré-requisitos

- Go 1.25.4+
- Docker e Docker Compose
- MySQL 8.0 (se rodar localmente sem Docker)

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
# Database
DB_ROOT_PASSWORD=root_password
DB_NAME=task_manager
DB_USER=task_user
DB_PASSWORD=task_password
DB_URL=task_user:task_password@tcp(db:3306)/task_manager?charset=utf8mb4&parseTime=True&loc=Local

# Server
PORT=8080
```

## 🐳 Executando com Docker

```bash
# Subir os containers
docker-compose up -d

# Verificar logs
docker-compose logs -f app

# Parar os containers
docker-compose down
```

A API estará disponível em `http://localhost:8080`

## 💻 Executando Localmente

```bash
# Instalar dependências
go mod download

# Executar a aplicação
go run cmd/api/main.go
```

## 📝 Endpoints da API

### Usuários
```http
GET    /users       # Listar usuários
GET    /users/:id   # Buscar usuário
POST   /users       # Criar usuário
PUT    /users/:id   # Atualizar usuário
DELETE /users/:id   # Deletar usuário
```

### Categorias
```http
GET    /categories       # Listar categorias
GET    /categories/:id   # Buscar categoria
POST   /categories       # Criar categoria
PUT    /categories/:id   # Atualizar categoria
DELETE /categories/:id   # Deletar categoria
```

### Tarefas
```http
GET    /tasks       # Listar tarefas
GET    /tasks/:id   # Buscar tarefa
POST   /tasks       # Criar tarefa
PUT    /tasks/:id   # Atualizar tarefa
DELETE /tasks/:id   # Deletar tarefa
```

## 🧪 Exemplo de Uso

```bash
# Criar uma tarefa
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Estudar Go",
    "description": "Revisar conceitos de goroutines",
    "category_id": 1,
    "user_id": 1
  }'
```

## 🗄️ Banco de Dados

O banco de dados MySQL é inicializado automaticamente via Docker Compose. As migrações/tabelas são criadas automaticamente pelo GORM na primeira execução.

## 📚 Estrutura de Dados

### Task (Tarefa)
```go
{
  "id": 1,
  "title": "string",
  "description": "string",
  "completed": false,
  "category_id": 1,
  "user_id": 1,
  "created_at": "2026-05-21T10:00:00Z",
  "updated_at": "2026-05-21T10:00:00Z"
}
```

## 🤝 Contribuindo

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto não possui licença definida.

## 👨‍💻 Autor

**Thomas Abner de Queiroz** - [@thomasaqx](https://github.com/thomasaqx)

---

⭐ Se este projeto foi útil para você, considere dar uma estrela!
