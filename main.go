package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Starting Application ======= goPods ")

	nodes := getNodes()
	signer := getPPKKey()
	config := configureSSHClient(signer)

	for _, node := range nodes {

		session := getSession(node, config)

		for number := 0; number < getAmountContainers(session); number++ {
			startContainers(session, 5)
		}
	}

}

/*
	SSH access settings
*/
func configureSSHClient(signer ssh.Signer) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: "rancher",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

/*
	Get the ssh key and set up a signer for connection
*/
func getPPKKey() ssh.Signer {
	pk, _ := ioutil.ReadFile("../goPods/.ssh/ssh-testes.pem")
	signer, err := ssh.ParsePrivateKey(pk)

	if err != nil {
		panic(err)
	}

	return signer
}

/*
	Returns all nodes that will start the containers
*/
func getNodes() []string {
	return []string{"54.91.241.253:22"}
}

/*
	Starts containers in status stopped, based on quantity spent on subscription @{quantity}
*/
func startContainers(session *ssh.Session, amount int) {
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	command := fmt.Sprintf("docker start $(docker ps -a --last %d --quiet --filter \"status=exited\")", amount)

	if err := session.Run(command); err != nil {
		log.Println("Erro:", err.Error())
	}

	fmt.Println("Initiating", amount, "containers")
	fmt.Println(b.String())
	time.Sleep(15 * time.Minute)

	fmt.Println("Waiting 15 minutes to start", amount, "more containers")
}

/*
	Realiza a conexão ssh com o node especificado na assintura do metodo
*/
func getSession(node string, config *ssh.ClientConfig) *ssh.Session {
	client, err := ssh.Dial("tcp", node, config)

	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()

	if err != nil {
		panic("Failed to create session: " + err.Error())
	}

	return session
}

/*
	Checks the number of containers in the node
*/
func getAmountContainers(session *ssh.Session) int {
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("docker ps -a | wc -l"); err != nil {
		log.Println("Erro:", err.Error())
	}

	fmt.Println(b.String())
	amount, _ := strconv.ParseInt(b.String(), 10, 64)

	return int(amount)
}
