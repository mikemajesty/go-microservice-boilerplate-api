package modules_cat

import (
	core_cat "go-microservice-boilerplate-api/core/cat/entity"
	core_usecase_cat "go-microservice-boilerplate-api/core/cat/use-case"
	"go-microservice-boilerplate-api/utils"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListAdatper struct{}
type UpdateAdatper struct{}
type DeleteAdatper struct{}
type CreateAdatper struct{}
type GetByIDAdatper struct{}

func (d *GetByIDAdatper) Validate(input primitive.ObjectID) error {
	return utils.IsObjectID(input, "ID")
}

func (c *CreateAdatper) Validate(cat *core_cat.CatEntity) error {
	return utils.ValidateSchema(cat.Name, "Name", validation.Required, validation.Length(3, 100))
}

func (d *DeleteAdatper) Validate(input primitive.ObjectID) error {
	return utils.IsObjectID(input, "ID")
}

func (d *UpdateAdatper) Validate(input *core_cat.CatEntity) utils.Nullable[string] {
	objectID, err := primitive.ObjectIDFromHex(input.ID)

	errorList := []string{}

	if err != nil {
		errorList = append(errorList, err.Error())
	}

	objectIdError := utils.IsObjectID(objectID, "ID")

	if objectIdError != nil {
		errorList = append(errorList, objectIdError.Error())
	}

	nameError := utils.ValidateSchema(input.Name, "Name", validation.Required, is.LowerCase, is.Alpha, validation.Length(3, 50))

	if nameError != nil {
		errorList = append(errorList, nameError.Error())
	}

	if len(errorList) == 0 {
		return nil
	}

	return strings.Join(errorList, ", ")

}

func (d *DeleteAdatper) CatDeleteExecute(input string) *utils.AppException {
	objectID, _ := primitive.ObjectIDFromHex(input)
	err := d.Validate(objectID)

	if err != nil {
		return utils.ApiBadRequestException(err.Error())
	}

	return core_usecase_cat.CatDeleteUsecase(catRepository)(input)
}

func (d *GetByIDAdatper) CatGetByIDExecute(input string) (*core_cat.CatEntity, *utils.AppException) {
	objectID, _ := primitive.ObjectIDFromHex(input)
	err := d.Validate(objectID)

	if err != nil {
		return nil, utils.ApiBadRequestException(err.Error())
	}

	return core_usecase_cat.CatGetByIDUsecase(catRepository)(input)
}

func (d *ListAdatper) CatListExecute(input utils.PostgresListInput) ([]core_cat.CatEntity, *utils.AppException) {
	return core_usecase_cat.CatListUsecase(catRepository)(input)
}

func (d *UpdateAdatper) CatUpdateExecute(input *core_cat.CatEntity) (*core_cat.CatEntity, *utils.AppException) {
	err := d.Validate(input)

	if err != nil {
		return nil, utils.ApiBadRequestException(err.(string))
	}

	return core_usecase_cat.CatUpdateUsecase(catRepository)(input)
}

func (c *CreateAdatper) CatCreateExecute(cat *core_cat.CatEntity) (utils.Nullable[string], *utils.AppException) {
	err := c.Validate(cat)

	if err != nil {
		return nil, utils.ApiBadRequestException(err.Error())
	}

	return core_usecase_cat.CatCreateUsecase(catRepository, loggerService)(cat)
}
