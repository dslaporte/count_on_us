# Count On Us

![Coverage](https://img.shields.io/badge/Coverage-41.2%25-yellow)

Projeto de controle de contas desenvolvido em Golang utilizando Clean Architecture

### Executando as dependências:

    docker compose up -d

Os bancos de dados `MySQL` e `Postgres` serão instanciados. Caso queira subir apenas algum deles, basta utilizar o comando:

    docker compose up <service_name> -d

Onde seu <service_name> pode ser postgres ou mysql dentro do [docker-compose.yml](docker-compose.yml)
