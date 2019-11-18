package errors

type ABFGuardError string

func (abfg ABFGuardError) Error() string {
	return string(abfg)
}

var (
	// Client-side errors
	ErrCLIFlagsAreNotSet          = ABFGuardError("CLI flags are not set")
	ErrCorruptConfigFileExtension = ABFGuardError("configuration file's extension is not supported")
	// Server-side errors
	ErrBadDBConfiguration = ABFGuardError("incorrect configuration is received")
	ErrBadRequest         = ABFGuardError("bad request")
	ErrEmptyIpList        = ABFGuardError("empty IP list")
	// DB-side errors
	ErrEmptyIP       = ABFGuardError("empty IP is received")
	ErrAlreadyStored = ABFGuardError("provided IP is already stored")
	ErrDoesNotExist  = ABFGuardError("provided IP does not exist in the DB")
	// Bucket-side errors

)

const (
	ErrServiceCmdPrefix = "server error"
	ErrClientCmdPrefix  = "client error"
)
