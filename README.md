# goPods

##### Description

Este codigo lambda foi criado com intensão de minimizar o impacto de multiplos pods iniciando em orquestadores como Rancher 2.x para que possa ser feita o pause das instancias em horarios não comerciais, devido a alta carga de trabalho que pode ser exigida dos mesmos.


## Configuration

### Install the Go
[install](https://golang.org/dl/)

### Configuration to Project

Atualize o ip da função getNodes para os ips quue deseja iniciar os pods

```
func getNodes() []string {
	return []string{"34.204.90.20:22"}
}
```

Coloque a chave pem da sua instancia na base do projeto


## Build the Projeto

Buildando para linux


#### Start

```
GOARCH=amd64 GOOS=linux go build start.go
```

#### Stopped

```
GOARCH=amd64 GOOS=linux go build stopped.go
```

## Zip files to lambda


#### Start

```
zip start.zip start ssh-testes.pem 
```

#### Stopped

```
zip stopped.zip start ssh-testes.pem 
```

## Install AWS CLI

[install](https://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-chap-install.html) 


## Create policy 

```
{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Principal": {
            "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
    }
}
```

## Create Role

```
 aws iam create-role --role-name lambda-basic-execution --assume-role-policy-document file:///Users/highlanderdantas/go/src/goPods/policy/lambda-trust-policy.json
```

### Attach Policy Role

```
aws iam attach-role-policy --role-name lambda-basic-execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```
### Get Role

```
aws iam get-role --role-name lambda-basic-execution
```

### Create Function Lambda

```
aws lambda create-function \
--function-name startPods_go \
--zip-file fileb:///Users/highlanderdantas/go/src/goPods/start.zip \
--handler start \
--runtime go1.x \
--role "arn:aws:iam::649330076382:role/lambda-basic-execution"
```
