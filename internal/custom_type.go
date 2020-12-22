package internal

type ChosenContactForCheck int

const (
	ChosenInvalid ChosenContactForCheck = 0
	ChosenEmail   ChosenContactForCheck = 1
	ChosenPhone   ChosenContactForCheck = 2
)

func (chosen ChosenContactForCheck) IsInvalid() bool {
	return chosen == ChosenInvalid
}

type ContactForCheck struct {
	phone  string                // 1
	email  string                // 2
	chosen ChosenContactForCheck // 3
}

func (c *ContactForCheck) Chosen() ChosenContactForCheck {
	return c.chosen
}

func NewContactForCheck(email string, phone string) *ContactForCheck {
	return &ContactForCheck{
		phone: phone,
		email: email,
	}
}

// Fix выполняет выбор первого заполненного контакта
func (c *ContactForCheck) Fix() {
	if c.email != "" {
		c.chosen = ChosenEmail
	} else if c.phone != "" {
		c.chosen = ChosenPhone
	}
}
