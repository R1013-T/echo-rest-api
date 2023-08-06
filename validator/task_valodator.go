package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go-rest-api/model"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("Title is required"),
			validation.RuneLength(1, 10).Error("Title must be between 1 and 10 characters"),
		),
	)
}
