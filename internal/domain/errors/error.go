package errors

// ABFGuardError .
type ABFGuardError string

// Error .
func (abfg ABFGuardError) Error() string {
	return string(abfg)
}

/** Client-side errors */
const (
	// ErrCLIFlagsAreNotSet .
	ErrCLIFlagsAreNotSet = ABFGuardError("CLI flags are not set")
	// ErrCorruptConfigFileExtension .
	ErrCorruptConfigFileExtension = ABFGuardError("configuration file's extension is not supported")
	// ErrAuthorisationFailed .
	ErrAuthorisationFailed = ABFGuardError("the authorisation request was declined")
	// ErrFlushBucketsFailed .
	ErrFlushBucketsFailed = ABFGuardError("flush buckets request failed")
	// ErrPurgeBucketFailed .
	ErrPurgeBucketFailed = ABFGuardError("the requested bucket was not removed")
)

/** Server-side errors */
const (
	// ErrMissingServerParameters .
	ErrMissingServerParameters = ABFGuardError("missing mandatory parameters to the server")
	// ErrBadDBConfiguration .
	ErrBadDBConfiguration = ABFGuardError("incorrect configuration is received")
	// ErrBadRequest .
	ErrBadRequest = ABFGuardError("bad request")
	// ErrEmptyIPList .
	ErrEmptyIPList = ABFGuardError("empty IP list")
	// ErrIPAddFailure .
	ErrAddIPFailure = ABFGuardError("ip was not added to the list")
	// ErrDeleteIPFailure .
	ErrDeleteIPFailure = ABFGuardError("ip was not deleted from the list")
)

/** DB-side errors */
const (
	// ErrEmptyIP .
	ErrEmptyIP = ABFGuardError("empty IP is received")
	// ErrEmptyIP .
	ErrEmptyLogin = ABFGuardError("empty login is received")
	// ErrEmptyIP .
	ErrEmptyPWD = ABFGuardError("empty password is received")
	// ErrAlreadyStored .
	ErrAlreadyStored = ABFGuardError("provided IP is already stored")
	// ErrDoesNotExist .
	ErrDoesNotExist = ABFGuardError("provided IP does not exist in the DB")
	// ErrIsInTheBlacklist .
	ErrIsInTheBlacklist = ABFGuardError("the ip is in the blacklist")
)

/** Bucket-side errors */
const (
	// ErrNoBucketFound .
	ErrNoBucketFound = ABFGuardError("no bucket found in store")
	// ErrDeleteMissingBucket .
	ErrDeleteMissingBucket = ABFGuardError("no bucket found in store for deletion")
	// ErrBucketFull .
	ErrBucketFull = ABFGuardError("some of the buckets returned false")
	// ErrEmptyBucketName .
	ErrEmptyBucketName = ABFGuardError("empty bucket name received")
)

/** Configuration errors */
const (
	// ErrEmptyCfgLogin .
	ErrIncorrectCfgLogin = ABFGuardError("incorrect login limit field in configuration")
	// ErrEmptyCfgPWD .
	ErrIncorrectCfgPWD = ABFGuardError("incorrect password limit field in configuration")
	// ErrEmptyCfgIP .
	ErrIncorrectCfgIP = ABFGuardError("incorrect ip limit field in configuration")
	// ErrEmptyCfgDuration
	ErrEmptyCfgDuration = ABFGuardError("empty expire field in configuration")
)

/** Manager errors */
const ()

const (
	// ErrServiceCmdPrefix .
	ErrServiceCmdPrefix = "server error"
	// ErrClientCmdPrefix .
	ErrClientCmdPrefix = "client error"
	// ErrBucketManagerPrefix .
	ErrBucketManagerPrefix = "bucket manager error"
	// ErrBucketStoragePrefix .
	ErrBucketStoragePrefix = "bucket manager error"
)
