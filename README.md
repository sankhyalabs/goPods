# goPods

##### Description


This lambda code was created with the intent of minimizing the impact of multiple pods starting in orchestrators such as Rancher 2.x so that instances can be paused at non-commercial times, due to the high workload that may be required of them.


## Configuration

### Install the Go
[click here and install](https://golang.org/dl/)

### Configuration to Project

Update the ip of the getNodes function to the ips you want to start the pods

```
func getNodes() []string {
	return []string{"34.204.90.20:22"}
}
```

Put the key pem of your instance in the base of the project


## Build the Project

Build for linux


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
zip stopped.zip stopped ssh-testes.pem 
```

## Install AWS CLI

[click here and install](https://docs.aws.amazon.com/pt_br/cli/latest/userguide/cli-chap-install.html) 


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
