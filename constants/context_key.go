package constants

type ContextKey string

const (
	RequestContextKey ContextKey = "RequestContextKey"
)

func (c ContextKey) String() string {
	return string(c)
}
