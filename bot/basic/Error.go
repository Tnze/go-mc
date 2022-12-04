package basic

type ErrorID int

const (
	Fatal ErrorID = iota // This is the only one that should emit a panic
	NoError
	OutOfBound
	InvalidChunk
	InvalidBlock
	InvalidEntity
	WriterError // Anything related to data writing
	ReaderError // Anything related to data reading
	DialError   // Anything related to login process
	NullValue   // Anything related to default values
	NoValue     // If a value is not set
	IncompleteExecution
	NotImplemented
)

type Error struct {
	Err  ErrorID
	Info error
}

func (e Error) Error() string {
	return e.Info.Error()
}

func (e Error) Unwrap() error {
	return e.Info
}

func (e Error) Is(id ErrorID) bool {
	return e.Err == id
}

func (e Error) IsExcept(ids ...ErrorID) bool {
	for _, id := range ids {
		if e.Err == id {
			return false
		}
	}
	return true
}
