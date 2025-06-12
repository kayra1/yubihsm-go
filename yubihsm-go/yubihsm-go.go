package yubihsmgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// deault to localhost:12345
func New(endpoint string) *HSM {
	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	return &HSM{
		endpoint: u.String(),
	}
}

func (hsm *HSM) BlinkDevice(sessionID SessionID, seconds uint32) error {
	_, err := hsm.sendCommand([]string{"device", "blink", fmt.Sprintf("%d", sessionID), fmt.Sprintf("%d", seconds)})
	return err
}

func (hsm *HSM) OpenSession(authkey uint16, password string) (SessionID, error) {
	args := []string{"session", "open", fmt.Sprintf("%d", authkey), password}
	result, err := hsm.sendCommand(args)
	if err != nil {
		return 0, err
	}
	return SessionID([]byte(result)[0]), nil
}

func (hsm *HSM) GenerateSymmetricKey(keyID uint16, password string) error {
	panic("not implemented")
}

func (hsm *HSM) EncryptAESCBC(keyID, iv uint16, data []byte) ([]byte, error) {
	// convert to base64 first
	panic("not implemented")
}

func (hsm *HSM) DecryptAESCBC(keyID, iv uint16, data []byte) ([]byte, error) {
	// convert from base64 first
	// all errors should be returned as errors.Join("failed command: ", err)
	panic("not implemented")
}

func (hsm *HSM) GetPseudoRandom(sessionID uint8, count uint16) ([]byte, error) {
	panic("not implemented")
}

func (hsm *HSM) CloseSession() error {
	panic("not implemented")
}

func (hsm *HSM) GetStatus() (string, error) {
	newRequest, err := http.NewRequest("GET", fmt.Sprintf("%s/connector/status", hsm.endpoint), nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (hsm *HSM) sendCommand(args []string) (string, error) {
	command := []byte(strings.Join(args, " "))
	buf := new(bytes.Buffer)
	buf.Write(command)

	newRequest, err := http.NewRequest("POST", fmt.Sprintf("%s/command/api", hsm.endpoint), buf)
	if err != nil {
		return "", err
	}
	newRequest.Header.Set("Content-Type", "application/octet-stream")
	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
