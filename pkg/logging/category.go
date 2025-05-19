package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	App             Category = "App"
	RequestResponse Category = "RequestResponse"
)

const (
	InternalInfo  SubCategory = "InternalInfo"
	InternalError SubCategory = "InternalError"
	API           SubCategory = "API"
	Bootstrapping SubCategory = "BootStrapping"
)

const (
	ClientIp     ExtraKey = "ClientIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Headers      ExtraKey = "Headers"
	ErrorMessage ExtraKey = "ErrorMessage"
)
