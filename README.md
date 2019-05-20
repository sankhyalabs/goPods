# goPods

## Descrição

Este codigo lambda foi criado com intensão de minimizar o impacto de multiplos pods iniciando em orquestadores como Rancher 2.x para que possa ser feita o pause das instancias em horarios não comerciais, devido a alta carga de trabalho de pode ser exigida dos mesmos.


### Para rodar o codigo instale o Go
https://golang.org/dl/

Navege para o projeto, coloque uma pasta .ssh/ com a chave pem da sua instância

```
.ssh/ssh-teste.pem

go run main.go
```

