package error

import "encoding/json"

type ErrorModel interface {
	error
	WithError(e error) ErrorModel
	WithId(id string) ErrorModel
	WithSource(source string) (err ErrorModel)

	WithMessage(message string) ErrorModel
	Message() string

	WithDetail(message string) ErrorModel
	WithType(errorType ErrorType) ErrorModel
	WithProperty(key string, value interface{}) ErrorModel
	WithProperties(properties map[string]interface{}) ErrorModel
	WithCode(code int) ErrorModel
	WithExtraInfo(map[string]interface{}) ErrorModel

	GetError() error
	Id() string

	Source() string
	Type() ErrorType
	Detail() string
	Properties() map[string]interface{}
	ExtraInfo() map[string]interface{}
	GetCode() int

	Is(errType ErrorType) bool

	Clone() ErrorModel
	IsMsgDefault() bool
	IsIdDefault() bool
}

type err struct {
	error
	ErrType          ErrorType `json:"-"`
	ID               string    `json:"-"`
	defaultId        bool
	ErrSource        string `json:"source,omitempty"`
	ErrMessage       string `json:"message,omitempty"`
	defaultMsg       bool
	Code             int                    `json:"code,omitempty"`
	ExtraInformation map[string]interface{} `json:"extraInfo,omitempty"`
	properties       map[string]interface{}
	detail           string
	// one is the content of the message for the CLDR plural form "one".
	one string
	// pluralCount determines which plural form of the message is used.
	pluralCount interface{}
}

func New(e ...error) (r ErrorModel) {
	r = new(err)
	if e != nil && len(e) > 0 && e[0] != nil {
		r = r.WithError(e[0])
	}
	return r
}

func (error *err) Clone() ErrorModel {
	e := *error
	return &e
}

func Parse(bytes []byte) (ErrorModel, error) {
	var e *err
	if err := json.Unmarshal(bytes, &e); err != nil {
		return nil, err
	}
	return e, nil
}

func (error *err) WithError(e error) ErrorModel {
	if e == nil {
		return error
	}
	error.error = e
	return error.WithDetail(e.Error())
}

func (error *err) WithCode(code int) ErrorModel {
	error.Code = code
	return error
}

func (error *err) GetCode() int {
	return error.Code
}

func (error *err) GetError() error {
	return error.error
}

func (error *err) setDefaults() {
	if error.ID == "" {
		switch error.ErrType {
		case TypeValidation:
			error.ID = "ValidationError"
			error.defaultId = true
			break
		case TypeNotFound:
			error.ID = "NotFoundError"
			error.defaultId = true
			break
		case TypeForbidden:
			error.ID = "ForbiddenError"
			error.defaultId = true
			break
		case TypeUnAuthorized:
			error.ID = "UnAuthorizedError"
			error.defaultId = true
			break
		case TypeInternal:
			error.ID = "InternalError"
			error.defaultId = true
			break
		}
	}
	if error.ErrMessage == "" {
		switch error.ErrType {
		case TypeValidation:
			error.ErrMessage = "Invalid request information"
			error.defaultMsg = true
			break
		case TypeNotFound:
			error.ErrMessage = "Requested information not found"
			error.defaultMsg = true
			break
		case TypeForbidden:
			error.ErrMessage = "Access to this section has been restricted"
			error.defaultMsg = true
			break
		case TypeUnAuthorized:
			error.ErrMessage = "You are not authorized"
			error.defaultMsg = true
			break
		case TypeInternal:
			error.ErrMessage = "Internal error occurred, Please contact administrator"
			error.defaultMsg = true
			break
		}
	}
}

func (error *err) WithId(id string) ErrorModel {
	error.ID = id
	error.defaultId = false
	return error
}

func (error *err) WithMessage(message string) ErrorModel {
	error.ErrMessage = message
	error.defaultMsg = false
	return error
}

func (error *err) Message() string {
	return error.ErrMessage
}

func (error *err) WithSource(source string) ErrorModel {
	error.ErrSource = source
	return error
}

func (error *err) WithDetail(d string) ErrorModel {
	error.detail = d
	return error
}

func (error *err) WithType(t ErrorType) ErrorModel {
	error.ErrType = t
	error.setDefaults()
	return error
}

func (error *err) WithProperty(key string, value interface{}) ErrorModel {
	if error.properties == nil {
		error.properties = map[string]interface{}{}
	}
	error.properties[key] = value
	return error
}

func (error *err) WithProperties(properties map[string]interface{}) ErrorModel {
	error.properties = properties
	return error
}

// WithExtraInfo adds extra information for client in error..multiple calls of this function adds to the map
// but it replaces the keys with same name
func (error *err) WithExtraInfo(extraInfo map[string]interface{}) ErrorModel {
	if error.ExtraInformation != nil {
		for k, v := range extraInfo {
			error.ExtraInformation[k] = v
		}
	} else {
		error.ExtraInformation = extraInfo
	}
	return error
}

func (error *err) Id() string {
	return error.ID
}

func (error *err) Source() string {
	return error.ErrSource
}

func (error *err) Error() string {
	var r string
	if error.ErrSource != "" {
		r += "source: " + error.ErrSource + ", "
	}
	if error.ID != "" {
		r += "id: " + error.ID + ", "
	}
	if error.ErrMessage != "" {
		r += "message: " + error.ErrMessage + ", "
	}
	if error.error != nil {
		r += "err: " + error.error.Error() + ", "
	}
	if error.detail != "" {
		r += "detail: " + error.detail
	}
	return r
}

func (error *err) Type() ErrorType {
	return error.ErrType
}

func (error *err) Detail() string {
	return error.detail
}

func (error *err) Properties() map[string]interface{} {
	return error.properties
}

func (error *err) ExtraInfo() map[string]interface{} {
	return error.ExtraInformation
}

func (error *err) IsMsgDefault() bool {
	return error.defaultMsg
}

func (error *err) IsIdDefault() bool {
	return error.defaultId
}
