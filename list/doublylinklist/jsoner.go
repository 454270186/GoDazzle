package doublylinklist

import "encoding/json"

func (d *DoublyLinkList) ToJson() ([]byte, error) {
	return json.Marshal(d.Values())
}

func (d *DoublyLinkList) FromJson(oriJson []byte) error {
	vals := []interface{}{}
	err := json.Unmarshal(oriJson, &vals)
	if err != nil {
		return err
	}

	d.Clear()
	d.Add(vals...)
	return nil
}

/* impl std json interface */

func (d *DoublyLinkList) MarshalJSON() ([]byte, error) {
	return d.ToJson()
}

func (d *DoublyLinkList) UnmarshalJSON(orijson []byte) error {
	return d.FromJson(orijson)
}