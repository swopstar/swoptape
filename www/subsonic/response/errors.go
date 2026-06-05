// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package response

type ErrorCode int

const (
	ErrGeneric       ErrorCode = 0
	ErrParam         ErrorCode = 10
	ErrUpgradeClient ErrorCode = 20
	ErrUpgradeServer ErrorCode = 30
	ErrAuthUserPass  ErrorCode = 40
	ErrAuthTokenLDAP ErrorCode = 41
	ErrAuthType      ErrorCode = 42
	ErrAuthConflict  ErrorCode = 43
	ErrAuthInvalid   ErrorCode = 44
	ErrUnauthorized  ErrorCode = 50
	ErrTrialOver     ErrorCode = 60
	ErrNotFound      ErrorCode = 70
)

var errors = map[ErrorCode]string{
	0:  "A generic error",
	10: "Required parameter is missing",
	20: "Incompatible Subsonic REST protocol version. Client must upgrade",
	30: "Incompatible Subsonic REST protocol version. Server must upgrade",
	40: "Wrong username or password",
	41: "Token authentication not supported for LDAP users",
	42: "Provided authentication mechanism not supported",
	43: "Multiple conflicting authentication mechanisms provided",
	44: "Invalid API key",
	50: "User is not authorized for the given operation",
	60: "The trial period for the Subsonic server is over. Please upgrade to Subsonic Premium. Visit subsonic.org for details",
	70: "The requested data was not found",
}

func GetError(code ErrorCode, message string) *Error {
	if message == "" {
		var ok bool
		message, ok = errors[code]
		if !ok {
			message = errors[ErrGeneric]
		}
	}

	return &Error{
		Code:    int(code),
		Message: &message,
	}
}
