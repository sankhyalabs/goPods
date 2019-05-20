package main

import (
	"bytes"
	"fmt"
	"goPods/utils"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Starting Application ======= goPods ")

	nodes := utils.GetNodes()
	signer := utils.GetPPKKey()
	config := utils.ConfigureSSHClient(signer)

	for _, node := range nodes {

		session := utils.GetSession(node, config)
		stoppedContainers(session)
	}

}

/*
	Checks the number of containers in the node
*/
func stoppedContainers(session *ssh.Session) {
	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("docker stop $(docker ps -qa)"); err != nil {
		log.Println("Erro:", err.Error())
	}

	fmt.Println(b.String())
}
