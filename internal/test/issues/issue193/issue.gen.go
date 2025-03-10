// Package issue52 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/d3code/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package issue52

import (
	"encoding/json"
	"fmt"
)

// Person defines model for Person.
type Person struct {
	Age                  *float32               `json:"age,omitempty"`
	Metadata             string                 `json:"metadata"`
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Getter for additional properties for Person. Returns the specified
// element and whether it was found
func (a Person) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Person
func (a *Person) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Person to handle AdditionalProperties
func (a *Person) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["age"]; found {
		err = json.Unmarshal(raw, &a.Age)
		if err != nil {
			return fmt.Errorf("error reading 'age': %w", err)
		}
		delete(object, "age")
	}

	if raw, found := object["metadata"]; found {
		err = json.Unmarshal(raw, &a.Metadata)
		if err != nil {
			return fmt.Errorf("error reading 'metadata': %w", err)
		}
		delete(object, "metadata")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Person to handle AdditionalProperties
func (a Person) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Age != nil {
		object["age"], err = json.Marshal(a.Age)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'age': %w", err)
		}
	}

	object["metadata"], err = json.Marshal(a.Metadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'metadata': %w", err)
	}

	if a.Name != nil {
		object["name"], err = json.Marshal(a.Name)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'name': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
