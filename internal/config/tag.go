package config

import (
	"errors"
	"reflect"
)

type Tag struct {
	Name string
	Default *string
	Required bool
}

func parseTag(tag reflect.StructTag) (*Tag, error) {
	errString := ""

	nameTag, ok := tag.Lookup("env")
	if !ok {
		errString += "config: field is missing env tag\n"
	}

	defaultTag, ok := tag.Lookup("default")

	var requiredTag bool = false
	temp, ok := tag.Lookup("required")
	if ok {
		switch temp {
		case "true":
			requiredTag = true
		case "false":
			requiredTag = false
		default:
			errString += "config: invalid value for required tag"
		}
	}

	if errString != "" {
		return nil, errors.New(errString)
	}

	return &Tag{
		Name: nameTag,
		Default: &defaultTag,
		Required: requiredTag,
	}, nil
}