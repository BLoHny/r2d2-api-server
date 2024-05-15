package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/oneee-playground/r2d2-api-server/internal/domain"
)

func RegisterTaskStage(v *validator.Validate) error {
	err := v.RegisterValidation("task_stage", func(fl validator.FieldLevel) bool {
		return TaskStageValid(domain.TaskStage(fl.Field().String()))
	})
	if err != nil {
		return err
	}

	return nil
}
