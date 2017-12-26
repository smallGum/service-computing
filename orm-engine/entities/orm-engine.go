package entities

import (
	"errors"
	"reflect"
	"service-computing/orm-engine/sqlt"
	"strings"
)

// ORMEngine definition
type ORMEngine struct {
	sqlt.SQLTemplate
}

// Insert insert new data entry into the table
func (e *ORMEngine) Insert(o interface{}) (int, error) {
	insertQuery, err := genInsertStmt(o)
	if err != nil {
		return 0, err
	}
	_, args, err := getTableField(o)
	if err != nil {
		return 0, err
	}

	_, err = e.Exec(insertQuery, args...)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

// Find query all the entries of the table
func (e *ORMEngine) Find(o interface{}) error {
	sliceValue := reflect.Indirect(reflect.ValueOf(o))
	if sliceValue.Kind() != reflect.Slice {
		return errors.New("needs a pointer to a slice")
	}

	sliceElementType := sliceValue.Type().Elem()
	data := sliceElementType.Elem()
	queryString, err := genQueryStmt(data.Name())
	if err != nil {
		return err
	}

	rows, _ := e.Query(queryString)

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	newSlice := reflect.MakeSlice(sliceValue.Type(), 0, 4)

	for rows.Next() {
		for i, _ := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			pv := reflect.New(data)
			element := pv.Elem()

			if ok {
				v = string(b)
			} else {
				v = val
			}

			for j := 0; j < data.NumField(); j++ {
				if strings.ToLower(data.Field(j).Name) == col {
					element.Field(j).Set(reflect.ValueOf(v))
					break
				}
			}

			newSlice = reflect.Append(newSlice, pv)
		}
	}

	s := reflect.ValueOf(o).Elem()
	s.Set(newSlice)

	return nil
}

// ---------------------
// helpful function
// ---------------------

// get database table name
func getTableName(o interface{}) (string, error) {
	t := reflect.TypeOf(o)
	if t.Name() == "" {
		return "", errors.New("non-exist interface type")
	}
	return strings.ToLower(t.Name()), nil
}

// get table field's name and value
func getTableField(o interface{}) ([]string, []interface{}, error) {
	fieldNames := make([]string, 0)
	fieldValues := make([]interface{}, 0)

	s := reflect.ValueOf(o)
	typeOfO := s.Type()
	if typeOfO.Kind() != reflect.Struct {
		return []string{}, []interface{}{}, errors.New("no struct type")
	}
	for i := 0; i < s.NumField(); i++ {
		fieldNames = append(fieldNames, strings.ToLower(typeOfO.Field(i).Name))
		fieldValues = append(fieldValues, s.Field(i).Interface())
	}

	return fieldNames, fieldValues, nil
}

// generate insert statement
func genInsertStmt(o interface{}) (string, error) {
	tableName, err := getTableName(o)
	if err != nil {
		return "", err
	}
	stmt := "INSERT " + tableName + " SET "
	fields, _, err := getTableField(o)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(fields)-1; i++ {
		stmt += fields[i] + "=?,"
	}
	stmt += fields[len(fields)-1] + "=?"

	return stmt, nil
}

// generate query statement
func genQueryStmt(tableName string) (string, error) {
	if tableName == "" {
		return "", errors.New("non-exist interface type")
	}

	return "SELECT * FROM " + tableName, nil
}
