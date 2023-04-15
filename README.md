# Count On Us

![Coverage](https://img.shields.io/badge/Coverage-41.2%25-yellow)

Projeto de controle de contas desenvolvido em Golang utilizando Clean Architecture

A estrutura de pastas segue o modelo sugerido no repositório `project-layout` do [golang standards](https://github.com/golang-standards/project-layout)

## Explicando a estrutura:

**/configs** - Contém as configurações para execução da aplicação

**/internal** - Contém o `core` do sistema;

**/internal/entity** - Abriga todas as regras de negócio, importante frisar que `entidades` não são `models`, pois enquanto as `entidades` referem-se as regras de negócio, `models` apenas atendem as necessidades do banco de dados (ORM por exemplo);

**/internal/infrastructure** - Abriga toda a parte de camadas de `tecnologia/frameworks` para funcionamento da aplicação. Devem estar isoladas e abstraídas de forma a serem substítuidas a qualquer momento _(design para interface / baixo acoplamento)_; Ex: banco de dados, sistemas de mensageria, apresentação (REST, gRPC, GraphQL)

**/internal/repositories ou gateways** - Concentra as regras de acesso a dados da aplicação, pode ser via banco de dados por exemplo ou acesso a APIs externas dentre outros;

**/internal/usecases** - Casos de uso descrevem as necessidades do sistema e gerenciam os acessos as camadas mais internas;

**/pkg** - Contém os possíveis packages/libraries a nível global (não possuem dependência de regra de negócio) para acesso de aplicações externas;

**/test ou /tests** - Arquivos para execução de testes unitários (preparação do ambiente de test), e2e e etc;

### Executando as dependências:

    docker compose up -d

O sistema atual permite o uso de bancos de dados `SQLite`, `MySQL` ou `Postgres`. Você pode escolher instanciar todos ao mesmo tempo por meio do comando:

    docker compose up -d

Caso queira selecionar apenas algum dos bancos, basta utilizar o comando:

    docker compose up <service_name> -d

Onde seu <service_name> pode ser `postgres` ou `mysql` dentro do arquivo [docker-compose.yml](docker-compose.yml)

---

_Developed by Daniel Sobrinho Laporte - daniel.laporte@gmail.com_
