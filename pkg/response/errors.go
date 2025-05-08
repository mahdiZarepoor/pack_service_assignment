package response

type Error interface {
	GetMessage() string
	GetAttributes() map[string]interface{}
}

type ServiceError struct {
	message    string
	attributes map[string]interface{}
}

func NewServiceError(msg string, attrs ...map[string]interface{}) *ServiceError {

	var attributes map[string]interface{}

	if len(attrs) > 0 {
		attributes = attrs[0]
	}

	return &ServiceError{
		message:    msg,
		attributes: attributes,
	}
}

func (e ServiceError) GetMessage() string {
	return e.message
}

func (e ServiceError) GetAttributes() map[string]interface{} {
	return e.attributes
}
