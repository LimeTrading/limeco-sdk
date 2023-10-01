package limecosdk

import "reflect"

func isStructFieldsSet(data interface{}) string {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Struct {
		return "Input is not a struct"
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		zeroValue := reflect.Zero(field.Type())

		if reflect.DeepEqual(field.Interface(), zeroValue.Interface()) {
			return value.Type().Field(i).Name
		}
	}

	return "" // All fields are set
}
