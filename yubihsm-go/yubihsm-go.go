package yubihsmgo

func (hsm *HSM) New(url string) string {
	panic("not implemented")
}

func (hsm *HSM) BlinkDevice(sessionID uint8, seconds uint32) error {
	panic("not implemented")
}

func (hsm *HSM) OpenSession(authkey uint16, password string) error {
	panic("not implemented")
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

func (hsm *HSM) OpenSessionAsym(authkey uint16, privkey string) error {
	panic("not implemented")
}

func (hsm *HSM) CloseSession() error {
	panic("not implemented")
}

func (hsm *HSM) sendCommand(args []string) (string, error) {
	panic("not implemented")
}
