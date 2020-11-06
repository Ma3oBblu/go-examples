package internal

// NewLogicError создаем логическую ошибку
func NewLogicError(err error, localizedMessage string) *LogicError {
	return &LogicError{
		err:              err,
		localizedMessage: localizedMessage,
	}
}

// LogicError ошибка в логике при работе с приложением
type LogicError struct {
	err              error
	localizedMessage string
}

// Error возвращает текст ошибки.
func (e *LogicError) Error() string {
	return e.err.Error()
}

// Unwrap возвращает обрамленную ошибку.
func (e *LogicError) Unwrap() error {
	return e.err
}

func GenerateError(err error, message string) error {
	return NewLogicError(err, message)
}
