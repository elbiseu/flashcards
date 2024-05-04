package types

type ContentType string

func (ct *ContentType) String() string {
	return string(*ct)
}
