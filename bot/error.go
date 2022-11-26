package bot

type Error struct {
	Err error
}

func (e Error) Error() string {
	return "bot/screen: " + e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}
