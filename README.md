# Desenvolvimento Web com Go - Do Zero ao Deploy

## Quicknotes

Aplicação web desenvolvida que trata-se de uma aplicação de gerenciamento de anotações e lembretes, similar ao Google Keep.

## Configuração

A aplicação pode ser configurada através de variáveis de ambiente. As variáveis disponíveis podem ser vistas no arquivo .env ou na struct Config do pacote main.

A struct Config possui uma tag (env) que define o nome da variável de ambiente e um valor default ou a palavra required, para os casos em que o valor precisa obrigatoriamente vir de uma variável de ambiente (por exemplo, valores confidenciais são exemplos de valores que devem ser configurados em variáveis de ambiente).

Abaixo podemos ver um exemplo de como configurar uma propriedade na struct Config.

```go
type Config struct {
    NomePropriedade string `env:"NOME_ENV_VAR,valor_default"`
    SecretValue string `env:"SECRET_VALUE,required"`
}
```

- NomePropriedade: nome da propriedade de configuração
- NOME_ENV_VAR: nome da variável de ambiente de onde o valor será lido

## Rotas da aplicação

| Método | Rota                     | Handler           | Descrição                         |
|:-------|:-------------------------|:------------------|:----------------------------------|
| GET    | /                        | HomeHandler       | Home Page                         |
| GET    | /note                    | NoteList          | Home Page                         |
| GET    | /note/{id}               | NoteView          | Visualiza uma anotação            |
| GET    | /note/new                | NoteNew           | Form de Criação de uma anotação   |
| POST   | /note/                   | NoteSave          | Cria uma anotação                 |
| DELETE | /note/{id}               | NoteDelete        | Remove uma anotação               |
| GET    | /note/{id}/edit          | NoteEdit          | Form de alteração de uma anotação |
| GET    | /user/signup             | SignupForm        | Form de registro de usuários      |
| POST   | /user/signup             | Signup            | Adiciona o usuário no banco       |
| GET    | /user/signin             | SigninForm        | Form de login de usuários         |
| POST   | /user/signin             | Signin            | Processa o login do usuário       |
| GET    | /user/signout            | Signout           | Processa o logout do usuário      |
| GET    | /user/password           | ResetPassword     | Form para alteração de senha      |
| POST   | /user/password/{token}   | ResetPasswordForm | Processa alteração de senha       |
| GET    | /user/forgetpassword     | ForgetPasswordForm| Form para alteração de senha      |
| POST   | /user/forgetpassword     | ForgetPassword    | Processa alteração de senha       |
| GET    | /confirmation/{token}    | Confirm           | Confirmação de email do cadastro  |

## Modelo de Banco de Dados

### NOTES

| CAMPO      | TIPO      | CONSTRAINT   |
|:-----------|:----------|:-------------|
| ID         | BIGSERIAL | PK, NOT NULL |
| TITLE      | TEXT      | NOT NULL     |
| CONTENT    | TEXT      |              |
| COLOR      | TEXT      | NOT NULL     |
| CREATED_AT | TIMESTAMP |              |
| UPDATED_AT | TIMESTAMP |              |
| USER_ID    | BIGINT    | NOT NULL     |

### USERS

| CAMPO      | TIPO      | CONSTRAINT             |
|:-----------|:----------|:-----------------------|
| ID         | BIGSERIAL | PK, NOT NULL           |
| EMAIL      | TEXT      | NOT NULL UNIQUE        |
| PASSWORD   | TEXT      | NOT NULL               |
| ACTIVE     | TEXT      | NOT NULL DEFAULT false |
| CREATED_AT | TIMESTAMP |                        |
| UPDATED_AT | TIMESTAMP |                        |

### USERS_CONFIRMATION_TOKENS

| CAMPO      | TIPO      | CONSTRAINT             |
|:-----------|:----------|:-----------------------|
| ID         | BIGSERIAL | PK, NOT NULL           |
| USER_ID    | BIGINT    | NOT NULL               |
| TOKEN      | TEXT      | NOT NULL               |
| CONFIRMED  | BOOLEAN   | NOT NULL DEFAULT false |
| CREATED_AT | TIMESTAMP |                        |
| UPDATED_AT | TIMESTAMP |                        |

### SESSIONS

| CAMPO      | TIPO        | CONSTRAINT   |
|:-----------|:------------|:-------------|
| TOKEN      | TEXT        | PK, NOT NULL |
| DATA       | BYTEA       | NOT NULL     |
| EXPIRY     | TIMESTAMPTZ | NOT NULL     |

## Execução

Para executar a aplicação com Docker localmente, execute o comando abaixo:

```bash
docker compose -f ./docker-compose.prod.yml up -d
```

## Demonstrações da Aplicação 

| Cadastro Usuário | Email confirmação |
|:--:|:--:|
| <img width="1913" height="907" alt="quicknotes1" src="https://github.com/user-attachments/assets/379bd7b4-d75e-4664-92e0-db1b52cf01c8" /> | <img width="1917" height="986" alt="quicknotes3" src="https://github.com/user-attachments/assets/0866f48e-5ffa-4777-a49b-9ef4c1b9bc6b" />

| Página Inicial | Nova anotação |
|:--:|:--:|
| <img width="1909" height="980" alt="quicknotes4" src="https://github.com/user-attachments/assets/10a7032f-8c5c-40dc-b389-bfdc1f5ebfe2" /> | <img width="1915" height="984" alt="quicknotes5" src="https://github.com/user-attachments/assets/a06be307-82be-403c-bb31-6e8ec6369d3a" />

| Editar nota | Menu com nota criada |
|:--:|:--:|
| <img width="1914" height="980" alt="quicknotes6" src="https://github.com/user-attachments/assets/cfb467ac-0ac0-4150-9646-4c212d8ade64" /> | <img width="1918" height="981" alt="quicknotes7" src="https://github.com/user-attachments/assets/1ed1f614-9678-4e48-bbae-790c11e0d099" />




