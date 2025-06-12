package yubihsmgo_test

import (
	"strings"
	"testing"

	yubihsmgo "github.com/kayra1/yubihsm-go/yubihsm-go"
)

// This test will modify items in your YubiHSM. Make sure that the following are true:
//   - yubihsm-connector is installed and running
//   - there is space for at least one session
//   - object_id 65535 is not in use
//   - there is space for at least one object
//   - an authentication key created for testing is available
//   - the key has access to domain 16,
//   - the key has all capabilities and delegated capabilities
//   - the key's password is "password"
func TestYubiHSMGo(t *testing.T) {
	client := yubihsmgo.New("http://localhost:12345")
	status, err := client.GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(status, "status=OK") {
		t.Fatal("status did not contain 'status=OK': ", status)
	}

	session, err := client.OpenSession(16, "883362")
	if err != nil {
		t.Fatal(err)
	}
	if session != 0 {
		t.Fatal("session ID was not 0")
	}

	// client.GenerateSymmetricKey(15, password string)
	// get random
	// client encrypt hello world
	// client decrypt hello world
	// validate they are equal
}

func Example() {
	// Example usage of the YubiHSM Go library
	// client := yubihsmgo.New("http://localhost:12345")
	// // sessionID, err := client.OpenSession(16, "password")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// client.GenerateSymmetricKey(15, password string)
}
