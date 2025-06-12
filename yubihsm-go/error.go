package yubihsmgo

import "errors"

var (
	ErrInvalidCommand           = errors.New("Unknown Command")
	ErrInvalidData              = errors.New("Malformed data for the command")
	ErrInvalidSession           = errors.New("The session has expired or does not exist")
	ErrAuthenticationFailed     = errors.New("Wrong Authentication Key")
	ErrSessionsFull             = errors.New("No more available sessions")
	ErrSessionFailed            = errors.New("Session setup failed")
	ErrStorageFull              = errors.New("Storage full")
	ErrWrongLength              = errors.New("Wrong data length for the command")
	ErrInsufficientPermissions  = errors.New("Insufficient permissions for the command")
	ErrLogFull                  = errors.New("The log is full and force audit is enabled")
	ErrObjectNotFound           = errors.New("No object found matching given ID and Type")
	ErrInvalidID                = errors.New("Invalid ID")
	ErrSSHCAConstraintViolation = errors.New("Constraints in SSH Template not met")
	ErrInvalidOTP               = errors.New("OTP decryption failed")
	ErrDemoMode                 = errors.New("Demo device must be power cycled")
	ErrObjectExists             = errors.New("Unable to overwrite object")
)

//OK	0x00	Success
// INVALID COMMAND	0x01	Unknown command
// INVALID DATA	0x02	Malformed data for the command
// INVALID SESSION	0x03	The session has expired or does not exist
// AUTHENTICATION FAILED	0x04	Wrong Authentication Key
// SESSIONS FULL	0x05	No more available sessions
// SESSION FAILED	0x06	Session setup failed
// STORAGE FAILED	0x07	Storage full
// WRONG LENGTH	0x08	Wrong data length for the command
// INSUFFICIENT PERMISSIONS	0x09	Insufficient permissions for the command
// LOG FULL	0x0a	The log is full and force audit is enabled
// OBJECT NOT FOUND	0x0b	No object found matching given ID and Type
// INVALID ID	0x0c	Invalid ID
// SSH CA CONSTRAINT
// VIOLATION
// 0x0e	Constraints in SSH Template not met
// INVALID OTP	0x0f	OTP decryption failed
// DEMO MODE	0x10	Demo device must be power-cycled
// OBJECT EXISTS	0x11	Unable to overwrite object
