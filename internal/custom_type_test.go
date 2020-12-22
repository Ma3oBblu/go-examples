package internal

import "testing"

func TestNewContactForCheck(t *testing.T) {
	contacts := NewContactForCheck("", "79198078281")
	switch contacts.Chosen() {
	case ChosenPhone:
		println("phone")
	case ChosenEmail:
		println("email")
	case ChosenInvalid:
		fallthrough
	default:
		println("default")
		println(contacts.Chosen().IsInvalid())
		contacts.Fix()
		println(contacts.Chosen().IsInvalid())
		println(contacts.Chosen())
	}
}
