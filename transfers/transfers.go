package transfers

import "encoding/json"

type Verb struct {
	BaseForm       string
	PastParticiple string
	PastSimple     string
	Type           string
	Value          string
}

func (v *Verb) JSON() ([]byte, error) {
	data := map[string]interface{}{
		"past_participle": v.PastParticiple,
		"base_form":       v.BaseForm,
		"past_simple":     v.PastSimple,
		"type":            v.Type,
		"value":           v.Value,
	}
	return json.Marshal(data)
}
