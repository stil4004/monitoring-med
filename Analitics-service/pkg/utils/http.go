package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Use a single instance of Validate, it caches struct info
var validate *validator.Validate = validator.New()

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

func ParseBody(r *http.Request, input any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	if err = json.Unmarshal(body, input); err != nil {
		return err
	}
	return validate.StructCtx(r.Context(), input)
}
