package logger

const (
	ErrLevelInfo    = "info"
	ErrLevelWarning = "warning"
	ErrLevelError   = "error"
	ErrLevelFatal   = "fatal"

	TableName     = "public.logs"
	Id            = "id"
	ClientID      = "client_id"
	ErrLevel      = "error_level"
	SessionOwner  = "session_owner"
	RequestMethod = "request_method"
	RequestPath   = "request_path"
	StatusCode    = "status_code"
	ErrCode       = "error_code"
	Message       = "message"
	CreationDate  = "creation_date"
)
