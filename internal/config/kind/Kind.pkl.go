// Code generated from Pkl module `remychantenay.pkl_playground.internal.config`. DO NOT EDIT.
package kind

import (
	"encoding"
	"fmt"
)

// Kind is an "enum" type.
type Kind string

const (
	Floating Kind = "floating"
	Fixed    Kind = "fixed"
)

// String returns the string representation of Kind
func (rcv Kind) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Kind)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Kind.
func (rcv *Kind) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "floating":
		*rcv = Floating
	case "fixed":
		*rcv = Fixed
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Kind`, str)
	}
	return nil
}
