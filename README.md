# Golang-etl-case
Caso de uso de etl com go, gin e postgres

## Como instalar
`$ git clone git@github.com:JaegerCaiser/golang-etl-case.git`

`$ cd golang-etl-case`

`$ docker-compose up --build`

## Executar POST via insomia ou postman
POST <base_url>/etl-golang/v1/file

**o docker compose subirá o serviço da porta 3000, considere localhost como *base_url***

**base_url = http://localhost:3000**

O docker-compose manterá os logs ativos, é possível observar os registros sendo cadastrados.

# Dificuldade Técnica
No inicio o postgres recebia os dados um por um, mas isso abria muitas conexões e acabava que alguns registros se perdiam.

# Solução Técnica
Agora é armazeda em um array de 1000 registros, quando o limite é atingido, os mil são registrados na base de dados em modo Batch
Diminuindo consideravelmente o tempo de resposta.
