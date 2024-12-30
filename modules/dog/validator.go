package modules_dog

import (
	core_dog_entity "go-microservice-boilerplate-api/core/dog/entity"
	core_dog_usecase "go-microservice-boilerplate-api/core/dog/use-case"
	"go-microservice-boilerplate-api/utils"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListAdatper struct{}
type UpdateAdatper struct{}
type GetByIDAdatper struct{}
type DeleteAdatper struct{}
type CreateAdatper struct{}

func (c *ListAdatper) DogListExecute(input utils.PaginationType) ([]core_dog_entity.DogEntity, *utils.AppException) {
	return core_dog_usecase.DogListUsecase(dogRepository)(input)
}

func (c *UpdateAdatper) Validate(input *core_dog_entity.DogEntity) utils.Nullable[string] {
	objectIdError := utils.IsObjectID(input.ID, "ID")

	errorList := []string{}

	if objectIdError != nil {
		errorList = append(errorList, objectIdError.Error())
	}

	nameEror := utils.ValidateSchema(input.Name, "Name", validation.Required, is.LowerCase, is.Alpha, validation.Length(3, 50))

	if nameEror != nil {
		errorList = append(errorList, nameEror.Error())
	}

	return strings.Join(errorList, ",")
}

func (c *UpdateAdatper) DogUpdateExecute(dog *core_dog_entity.DogEntity) (*core_dog_entity.DogEntity, *utils.AppException) {
	err := c.Validate(dog)

	if err != nil {
		return nil, utils.ApiBadRequestException(err.(string))
	}
	return core_dog_usecase.DogUpdateUsecase(dogRepository)(dog)
}

func (c *GetByIDAdatper) Validate(input primitive.ObjectID) error {
	return utils.IsObjectID(input, "ID")
}

func (c *GetByIDAdatper) DogGetByIDExecute(id string) (*core_dog_entity.DogEntity, *utils.AppException) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := c.Validate(objectID)

	if err != nil {
		return nil, utils.ApiBadRequestException(err.Error())
	}

	return core_dog_usecase.DogGetByIDUsecase(dogRepository)(id)
}

func (c *DeleteAdatper) Validate(input primitive.ObjectID) error {
	return utils.IsObjectID(input, "ID")
}

func (c *DeleteAdatper) DogDeleteExecute(id string) *utils.AppException {
	objectID, _ := primitive.ObjectIDFromHex(id)
	err := c.Validate(objectID)

	if err != nil {
		return utils.ApiBadRequestException(err.Error())
	}

	return core_dog_usecase.DogDeleteUsecase(dogRepository)(id)
}

func (c *CreateAdatper) Validate(dog *core_dog_entity.DogEntity) error {
	return utils.ValidateSchema(dog.Name, "Name", validation.Required, validation.Length(3, 100))
}

func (c *CreateAdatper) DogCreateExecute(dog *core_dog_entity.DogEntity) (utils.Nullable[string], *utils.AppException) {
	err := c.Validate(dog)
	if err != nil {
		return nil, utils.ApiBadRequestException(err.Error())
	}
	return core_dog_usecase.DogCreateUsecase(dogRepository)(dog)
}
