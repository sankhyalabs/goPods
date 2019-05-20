package main

import (
	"bytes"
	"fmt"
	"goPods/utils"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Starting Application ======= goPods ")

	nodes := utils.GetNodes()
	signer := utils.GetPPKKey()
	config := utils.ConfigureSSHClient(signer)

	for _, node := range nodes {

		session := utils.GetSession(node, config)

		for number := 0; number < getAmountContainers(session); number++ {
			startContainers(session, 5)
		}
	}

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
