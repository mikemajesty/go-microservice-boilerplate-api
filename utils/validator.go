package utils

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateSchema(value interface{}, property string, rules ...validation.Rule) error {
	err := validation.Validate(value, rules...)

	if err == nil {
		return nil
	}

	return errors.New(property + ": " + err.Error())
}

func IsObjectID(value primitive.ObjectID, property string) error {
	var v, e = primitive.ObjectIDFromHex(value.Hex())

	if e != nil {
		return e
	}

	if v.Hex() == "000000000000000000000000" {
		return errors.New(property + ": is not valid objectID")
	}

	return nil
}
