package arraylist

import (
	"encoding/json"
)

func (a *ArrayList) ToJson() ([]byte, error) {
	return json.Marshal(a.elements[:a.size])
}

func (a *ArrayList) FromJson(oriJson []byte) error {
	return json.Unmarshal(oriJson, &a.elements)
}

/* impl std json interface */

func (a *ArrayList) MarshalJSON() ([]byte, error) {
	return a.ToJson()
}

func (a *ArrayList) UnmarshalJSON(orijson []byte) error {
	return a.FromJson(orijson)
}