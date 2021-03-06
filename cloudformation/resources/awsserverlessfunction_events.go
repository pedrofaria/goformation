package resources

import (
	"sort"

	"encoding/json"
)

// AWSServerlessFunction_Events is a helper struct that can hold either a String or String value
type AWSServerlessFunction_Events struct {
	String *string

	StringArray *[]string
}

func (r AWSServerlessFunction_Events) value() interface{} {
	ret := []interface{}{}

	if r.String != nil {
		ret = append(ret, r.String)
	}

	if r.StringArray != nil {
		ret = append(ret, r.StringArray)
	}

	sort.Sort(byJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r AWSServerlessFunction_Events) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *AWSServerlessFunction_Events) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case string:
		r.String = &val

	case []string:
		r.StringArray = &val

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

	case []interface{}:

		json.Unmarshal(b, &r.StringArray)

	}

	return nil
}
