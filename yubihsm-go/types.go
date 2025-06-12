package yubihsmgo

type SessionID uint8
type ObjectID uint16
type Domain uint8

type Capability string

const (
	CapabilityEncryptCBC Capability = "encrypt-cbc"
	CapabilityDecryptCBC Capability = "decrypt-cbc"
	CapabilityEncryptEBC Capability = "encrypt-ebc"
	CapabilityDecryptEBC Capability = "decrypt-ebc"

//	change-authentication-key      (0000400000000000)
//
// create-otp-aead                (0000000040000000)
// decrypt-cbc                    (0010000000000000)
// decrypt-ecb                    (0004000000000000)
// decrypt-oaep                   (0000000000000400)
// decrypt-otp                    (0000000020000000)
// decrypt-pkcs                   (0000000000000200)
// delete-asymmetric-key          (0000020000000000)
// delete-authentication-key      (0000010000000000)
// delete-hmac-key                (0000080000000000)
// delete-opaque                  (0000008000000000)
// delete-otp-aead-key            (0000200000000000)
// delete-public-wrap-key         (0000000000200000)
// delete-symmetric-key           (0002000000000000)
// delete-template                (0000100000000000)
// delete-wrap-key                (0000040000000000)
// derive-ecdh                    (0000000000000800)
// encrypt-cbc                    (0020000000000000)
// encrypt-ecb                    (0008000000000000)
// export-wrapped                 (0000000000001000)
// exportable-under-wrap          (0000000000010000)
// generate-asymmetric-key        (0000000000000010)
// generate-hmac-key              (0000000000200000)
// generate-otp-aead-key          (0000001000000000)
// generate-symmetric-key         (0001000000000000)
// generate-wrap-key              (0000000000008000)
// get-log-entries                (0000000001000000)
// get-opaque                     (0000000000000001)
// get-option                     (0000000000040000)
// get-pseudo-random              (0000000000080000)
// get-template                   (0000000004000000)
// import-wrapped                 (0000000000002000)
// put-asymmetric-key             (0000000000000008)
// put-authentication-key         (0000000000000004)
// put-mac-key                    (0000000000100000)
// put-opaque                     (0000000000000002)
// put-otp-aead-key               (0000000800000000)
// put-public-wrap-key            (0000000000100000)
// put-symmetric-key              (0000800000000000)
// put-template                   (0000000008000000)
// put-wrap-key                   (0000000000004000)
// randomize-otp-aead             (0000000080000000)
// reset-device                   (0000000010000000)
// rewrap-from-otp-aead-key       (0000000100000000)
// rewrap-to-otp-aead-key         (0000000200000000)
// set-option                     (0000000000020000)
// sign-attestation-certificate   (0000000400000000)
// sign-ecdsa                     (0000000000000080)
// sign-eddsa                     (0000000000000100)
// sign-hmac                      (0000000000400000)
// sign-pkcs                      (0000000000000020)
// sign-pss                       (0000000000000040)
// sign-ssh-certificate           (0000000002000000)
// unwrap-data                    (0000004000000000)
// verify-hmac                    (0000000000800000)
// wrap-data                      (0000002000000000)
)

type Algorithm string

const (
	AlgorithmAES192          Algorithm = "aes192"
	AlgorithmAES256          Algorithm = "aes256"
	AlgorithmAES128_CCM_WRAP Algorithm = "aes128-ccm-wrap"

	// aes128-yubico-authentication
	// aes128-yubico-otp
	// aes192-ccm-wrap
	// aes192-yubico-otp
	// aes256-ccm-wrap
	// aes256-yubico-otp
	AlgorithmAES_CBC Algorithm = "aes-cbc"
	AlgorithmAES_ECB Algorithm = "aes-ecb"

// aes-kwp
// ecbp256
// ecbp384
// ecbp512
// ecdh
// ecdsa-sha1
// ecdsa-sha256
// ecdsa-sha384
// ecdsa-sha512
// eck256
// ecp224
// ecp256
// ecp256-yubico-authentication
// ecp384
// ecp521
// ed25519
// hmac-sha1
// hmac-sha256
// hmac-sha384
// hmac-sha512
// mgf1-sha1
// mgf1-sha256
// mgf1-sha384
// mgf1-sha512
// opaque-data
// opaque-x509-certificate
// rsa-oaep-sha1
// rsa-oaep-sha256
// rsa-oaep-sha384
// rsa-oaep-sha512
// rsa-pkcs1-decrypt
// rsa-pkcs1-sha1
// rsa-pkcs1-sha256
// rsa-pkcs1-sha384
// rsa-pkcs1-sha512
// rsa-pss-sha1
// rsa-pss-sha256
// rsa-pss-sha384
// rsa-pss-sha512
// rsa2048
// rsa3072
// rsa4096
// template-ssh
)

type HSM struct {
	endpoint string
}

// HSMConnection struct
//   connection
//   session
//
// Session
// Object
// domain enum from 1-16
// id
// type

type Object struct {
	ID     uint32
	Domain uint8
	Type   uint8
}
