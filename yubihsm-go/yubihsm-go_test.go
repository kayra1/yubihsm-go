package yubihsmgo_test

import (
	"fmt"
	"testing"

	yubihsmgo "github.com/kayra1/yubihsm-go/yubihsm-go"
)

func TestYubiHSMGo(t *testing.T) {
	client := yubihsmgo.New("http://localhost:12345")
	blah, err := client.GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(blah)
	fmt.Println(err)

	session, err := client.OpenSession(16, "883362")
	fmt.Print(session)
	fmt.Print(err)
	if err != nil {
		t.Fatal(err)
	}
}

func Example() {
	// Example usage of the YubiHSM Go library
}
