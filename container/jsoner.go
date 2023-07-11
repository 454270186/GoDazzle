package container

type JsonCoder interface {
	ToJson() ([]byte, error)
	MarshalJSON() ([]byte, error) // impl json.Marshaler interface
}

type JsonDecoder interface {
	FromJson([]byte) error
	UnmarshalJSON([]byte) error // impl json.Unmarshaler interface
}