# 📦 golang\_sqlc

Repositório desenvolvido para estudo e prática de **migrations**, **query building** e **integração com banco de dados** utilizando Go e a ferramenta **sqlc**.

## ✨ Tecnologias Utilizadas

* **Go (Golang)**: linguagem principal do projeto.
* **SQLC**: geração de código Go a partir de arquivos SQL.
* **Docker & Docker Compose**: para orquestração de containers, incluindo o banco de dados.
* **MySQL**: banco de dados utilizado nos testes e desenvolvimento.
* **migrate**: ferramenta de controle de versionamento de migrações SQL.
* **Makefile**: automatização de comandos e rotinas de build e execução.
* **Go Modules**: gerenciamento de dependências.
* **Git**: versionamento de código-fonte.

## 📁 Estrutura do Projeto

```
golang_sqlc/
├── sql/
│   ├── migrations/          # Arquivos de migração (up/down)
│   └── queries/             # Consultas SQL utilizadas pelo sqlc
├── internal/db/             # Código gerado pelo sqlc + modelos e utilitários DB
├── docker-compose.yaml      # Configuração do ambiente com Docker
├── sqlc.yaml                # Configuração da ferramenta sqlc
├── Makefile                 # Comandos automatizados do projeto
├── go.mod / go.sum          # Dependências Go
└── README.md                # Documentação do projeto
```

## 🛠️ Funcionalidades

* Geração automática de código Go para acesso ao banco de dados a partir de arquivos SQL.
* Aplicação e reversão de migrations versionadas.
* Integração com MySQL utilizando Docker.
* Estrutura modular com foco em organização de queries e camada de dados.

## 🚀 Como executar

1. **Suba o ambiente com Docker:**

```bash
docker-compose up -d
```

2. **Execute as migrações:**

```bash
make migrate
```

3. **Gere o código Go a partir das queries SQL:**

```bash
make generate
```

4. **Execute o projeto (se aplicável):**

```bash
go run main.go
```

> Obs: certifique-se de ter `migrate`, `sqlc` e `docker-compose` instalados localmente, caso não utilize um ambiente pré-configurado.

## 🧪 Exemplos de queries

Arquivo `sql/queries/query.sql`:

```sql
-- name: ListCourses :many
SELECT * FROM courses;
```

Arquivo gerado: `internal/db/query.sql.go`.

## 📚 Referências

* [sqlc.dev](https://sqlc.dev/)
* [golang.org](https://golang.org/)
* [golang-migrate](https://github.com/golang-migrate/migrate)
