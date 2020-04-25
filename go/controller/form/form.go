package form

import "main/model"

// Form ...
type Form interface {
	ParseToDto() model.Dto
}
