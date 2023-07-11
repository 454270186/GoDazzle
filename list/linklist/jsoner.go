package linklist

import "encoding/json"

func (l *Linklist) ToJson() ([]byte, error) {
	return json.Marshal(l.Values())
}

func (l *Linklist) FromJson(oriJson []byte) error {
	vals := []interface{}{}
	if err := json.Unmarshal(oriJson, &vals); err != nil {
		return err
	}

	l.Clear()
	l.Add(vals...)
	return nil
}

/* impl std json interface */

func (l *Linklist) MarshalJSON() ([]byte, error) {
	return l.ToJson()
}

func (l *Linklist) UnmarshalJSON(orijson []byte) error {
	return l.FromJson(orijson)
}