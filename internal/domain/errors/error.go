package errors

type ABFGuardError string

func (abfg ABFGuardError) Error() string {
	return string(abfg)
}

var (
	ErrCLIFlagsAreNotSet          = ABFGuardError("CLI flags are not set")
	ErrCorruptConfigFileExtension = ABFGuardError("configuration file's extension is not supported")
	//
	ErrBadDBConfiguration = ABFGuardError("")
)

const (
	ErrServiceCmdPrefix = "server error"
	ErrClientCmdPrefix  = "client error"
)
