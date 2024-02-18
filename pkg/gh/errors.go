package gh

type GHErrorCode int
type GHError struct {
	message string
	status  GHErrorCode
}

func (gh GHError) Error() string {
	return gh.message
}

const (
	INVALID_HTML = 0
)
