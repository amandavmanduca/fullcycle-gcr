### Descrição

Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.


### Para executar localmente (Docker)
```bash
make run
```

### Para executar os testes
```bash
make test
```

### Request exemplo
`GET http://localhost:8080/address-info?cep=xxxxxxxx`


### Deploy GCR Base URL

`https://fullcycle-gcr-664058970217.europe-west1.run.app`

Exemplo:
`https://fullcycle-gcr-664058970217.europe-west1.run.app/address-info?cep=96090280`

Reponse
```
{"data":{"temp_c":16.2,"temp_f":61.16,"temp_k":289.2}}
```