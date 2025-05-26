package stats

type (
	Rows []Value

	Value struct {
		Timestamp        int64  `db:"created_at"`
		Params           string `db:"params"`
		Headers          string `db:"headers"`
		Method           string `db:"method_name"`
		Success          bool   `db:"success"`
		ErrorDescription string `db:"error_description"`
		Hash             string `db:"hash"`
		BodyOnError      string `db:"resp_body_on_error"`
	}
)
