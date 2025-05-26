package baseModel

// swagger:response baseModel
type BaseModel struct {
	// Example: true, false
	OK bool `json:"ok"`
	// Example: ok, invalid params, no access token, internal error, no data
	Description string `json:"description"`
	// Example: ok, bad_request, unauthorized
	// Code of error: f7a45bf8-946b-4932-b56b-ceb8db881f61
	InternalCode string `json:"internal_code"`
}

var (
	BadRequest = BaseModel{
		OK:           false,
		Description:  "invalid params",
		InternalCode: "bad_request",
	}

	Unauthorized = BaseModel{
		OK:           false,
		Description:  "no access token",
		InternalCode: "unauthorized",
	}

	OK = BaseModel{
		OK:           true,
		Description:  "ok",
		InternalCode: "ok",
	}
)
