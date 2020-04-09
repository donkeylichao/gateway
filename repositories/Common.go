package repositories

import (
	"reflect"
	"store/helper"
)

type CommonRepository struct {
}

// format return data many
func (c CommonRepository) GetSelectField(data []interface{}, fields ...string) (ml []interface{}) {

	fieldsLen := len(fields)
	// if fields length is 0 get the struct all fields
	if fieldsLen == 0 {
		val := reflect.ValueOf(data[0])
		// get fields by defined method FixFields
		fieldMethod := val.MethodByName("FixFields")
		if fieldMethod.IsValid() {
			fieldsByMethod := fieldMethod.Call(nil)
			if fieldsByMethod[0].Interface() != nil {
				fields = fieldsByMethod[0].Interface().([]string)
			} else {
				fields = []string{}
			}
		} else {
			fields = c.GetAllFields(data[0])
		}
	}

	for _, v := range data {
		m := make(map[string]interface{})
		val := reflect.ValueOf(v)
		typ := reflect.TypeOf(v)
		kname := ""

		for _, fname := range fields {
			if val.FieldByName(fname).IsValid() {
				// get kname
				fieldStruct, bool := typ.FieldByName(fname)
				if bool {
					kname = fieldStruct.Tag.Get("json")
				} else {
					kname = fname
				}
				// get format fields method
				method := val.MethodByName("Get" + fname)
				if method.IsValid() {
					va := method.Call(nil)
					if va[0].Interface() != nil {
						m[kname] = va[0].Interface()
					} else {
						m[kname] = val.FieldByName(fname).Interface()
					}
				} else {
					m[kname] = val.FieldByName(fname).Interface()
				}
			} else {
				kname = helper.UnMarshal(fname)
				// get format fields method
				method := val.MethodByName("Get" + fname)
				if method.IsValid() {
					va := method.Call(nil)
					if va[0].Interface() != nil {
						m[kname] = va[0].Interface()
					}
				}
			}
		}
		ml = append(ml, m)
	}
	return
}

// format return data one
func (c CommonRepository) GetSelectFieldOne(v interface{}, fields ...string) map[string]interface{} {
	fieldsLen := len(fields)
	// if fields length is 0 get the struct all fields
	if fieldsLen == 0 {
		val := reflect.ValueOf(v)
		// get fields by defined method FixFields
		fieldMethod := val.MethodByName("FixFields")
		if fieldMethod.IsValid() {
			fieldsByMethod := fieldMethod.Call(nil)
			if fieldsByMethod[0].Interface() != nil {
				fields = fieldsByMethod[0].Interface().([]string)
			} else {
				fields = []string{}
			}
		} else {
			fields = c.GetAllFields(v)
		}
	}

	m := map[string]interface{}{}

	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)
	kname := ""

	for _, fname := range fields {
		if val.FieldByName(fname).IsValid() {
			// get kname
			fieldStruct, bool := typ.FieldByName(fname)
			if bool {
				kname = fieldStruct.Tag.Get("json")
			} else {
				kname = fname
			}
			// get format fields method
			method := val.MethodByName("Get" + fname)
			if method.IsValid() {
				va := method.Call(nil)
				if va[0].Interface() != nil {
					m[kname] = va[0].Interface()
				} else {
					m[kname] = val.FieldByName(fname).Interface()
				}
			} else {
				m[kname] = val.FieldByName(fname).Interface()
			}
		} else {
			kname = helper.UnMarshal(fname)
			// get format fields method
			method := val.MethodByName("Get" + fname)
			if method.IsValid() {
				va := method.Call(nil)
				if va[0].Interface() != nil {
					m[kname] = va[0].Interface()
				}
			}
		}
	}

	return m
}

func (c CommonRepository) GetDbFields(model interface{}, fields ...string) (dbFields []string) {
	tye := reflect.TypeOf(model)
	for _, v := range fields {
		structField, ok := tye.FieldByName(v)
		if !ok {
			continue
		}
		columnName := structField.Tag.Get("orm")
		if columnName == "-" {
			continue
		}
		dbFields = append(dbFields, v)
	}
	return
}

func (c CommonRepository) GetAllFields(model interface{}) (allFields []string) {
	tye := reflect.TypeOf(model)
	fields := tye.NumField()

	for i := 0; i < fields; i++ {
		f := tye.Field(i)
		allFields = append(allFields, f.Name)
	}
	return
}
