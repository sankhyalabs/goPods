package utils

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

/*
	SSH access settings
*/
func ConfigureSSHClient(signer ssh.Signer) *ssh.ClientConfig {
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
func GetPPKKey() ssh.Signer {
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
func GetNodes() []string {
	return []string{"54.91.241.253:22"}
}

/*
	Realiza a conexão ssh com o node especificado na assintura do metodo
*/
func GetSession(node string, config *ssh.ClientConfig) *ssh.Session {
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
