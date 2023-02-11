package utils

type (
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}

	successJSON struct {
		Meta Meta        `json:"meta"`
		Data interface{} `json:"data"`
	}

	errorJSON struct {
		Meta Meta        `json:"meta"`
		Data interface{} `json:"data"`
	}
)

func SuccessResponse(message string, code int, data interface{}) successJSON {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  "success",
	}

	res := successJSON{
		Meta: meta,
		Data: data,
	}
	return res
}

func ErrorResponse(message string, code int, data interface{}) errorJSON {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  "failed",
	}

	res := errorJSON{
		Meta: meta,
		Data: data,
	}

	return res
}
