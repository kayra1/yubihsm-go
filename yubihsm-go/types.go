package yubihsmgo

type SessionID uint8

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
