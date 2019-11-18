package errors

// ABFGuardError .
type ABFGuardError string

// Error .
func (abfg ABFGuardError) Error() string {
	return string(abfg)
}

var (
	/**
	* Client-side errors
	 */

	// ErrCLIFlagsAreNotSet .
	ErrCLIFlagsAreNotSet = ABFGuardError("CLI flags are not set")
	// ErrCorruptConfigFileExtension .
	ErrCorruptConfigFileExtension = ABFGuardError("configuration file's extension is not supported")
	/**
	* Server-side errors
	 */

	// ErrBadDBConfiguration .
	ErrBadDBConfiguration = ABFGuardError("incorrect configuration is received")
	// ErrBadRequest .
	ErrBadRequest = ABFGuardError("bad request")
	// ErrEmptyIPList .
	ErrEmptyIPList = ABFGuardError("empty IP list")
	/**
	* DB-side errors
	 */

	// ErrEmptyIP .
	ErrEmptyIP = ABFGuardError("empty IP is received")
	// ErrAlreadyStored .
	ErrAlreadyStored = ABFGuardError("provided IP is already stored")
	// ErrDoesNotExist .
	ErrDoesNotExist = ABFGuardError("provided IP does not exist in the DB")
	/**
	* Bucket-side errors
	 */
)

const (
	// ErrServiceCmdPrefix .
	ErrServiceCmdPrefix = "server error"
	// ErrClientCmdPrefix .
	ErrClientCmdPrefix = "client error"
)
