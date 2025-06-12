package yubihsmgo

import (
	"bytes"
	"encoding/base64"
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

// all errors should be returned as errors.Join("failed command: ", err)
func (hsm *HSM) OpenSession(authkey uint16, password string) (SessionID, error) {
	result, err := hsm.sendCommand([]string{"session", "open", fmt.Sprintf("%d", authkey), password})
	if err != nil {
		return 0, err
	}
	return SessionID([]byte(result)[0]), nil
}

func (hsm *HSM) GenerateSymmetricKey(sessionID SessionID, keyID ObjectID, label string, domains []Domain, capabilities []Capability, algorithm Algorithm) error {
	domainsStr := make([]string, len(domains))
	capabilitiesStr := make([]string, len(capabilities))
	for i, domain := range domains {
		domainsStr[i] = fmt.Sprintf("%d", domain)
	}
	for i, capability := range capabilities {
		capabilitiesStr[i] = fmt.Sprintf("%s", capability)
	}
	_, err := hsm.sendCommand([]string{
		"generate", "symmetric",
		fmt.Sprintf("%d", sessionID),
		fmt.Sprintf("%d", keyID),
		label,
		strings.Join(domainsStr, ","),
		strings.Join(capabilitiesStr, ","),
		fmt.Sprintf("%s", algorithm),
	})
	return err
}

func (hsm *HSM) EncryptAESCBC(sessionID SessionID, keyID ObjectID, iv uint16, data []byte) ([]byte, error) {
	result, err := hsm.sendCommand([]string{
		"encrypt", "aescbc",
		fmt.Sprintf("%d", sessionID),
		fmt.Sprintf("%d", keyID),
		fmt.Sprintf("%#x", iv),
		base64.StdEncoding.EncodeToString(data),
	})
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

func (hsm *HSM) DecryptAESCBC(sessionID SessionID, keyID ObjectID, iv uint16, data []byte) ([]byte, error) {
	result, err := hsm.sendCommand([]string{
		"decrypt", "aescbc",
		fmt.Sprintf("%d", sessionID),
		fmt.Sprintf("%d", keyID),
		fmt.Sprintf("%#x", iv),
		string(data),
	})
	if err != nil {
		return nil, err
	}
	decryptedString, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		return nil, err
	}
	return decryptedString, nil
}

func (hsm *HSM) GetPseudoRandom(sessionID uint8, count uint32) ([]byte, error) {
	result, err := hsm.sendCommand([]string{
		"get", "random",
		fmt.Sprintf("%d", sessionID),
		fmt.Sprintf("%d", count),
		"-",
	})
	if err != nil {
		return nil, err
	}
	return []byte(result), nil
}

func (hsm *HSM) CloseSession(sessionID SessionID) error {
	_, err := hsm.sendCommand([]string{"session", "close", fmt.Sprintf("%d", sessionID)})
	return err
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
