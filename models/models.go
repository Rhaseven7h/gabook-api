package gabookApiModels

type (
	// Author models book authors
	Author struct {
		Name      string `json:"name"`
		Biography string `json:"biography"`
	}
)
