package response

import (
	"encoding/json"
	"encoding/xml"
	"time"
)

// DateTime is a time.Time that marshals to/from RFC3339 in both XML attributes and JSON.
type DateTime time.Time

func (d DateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: time.Time(d).Format(time.RFC3339)}, nil
}

func (d *DateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	t, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(time.RFC3339))
}

func (d *DateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}
