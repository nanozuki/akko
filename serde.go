package ononoki

import (
	"encoding/json"
)

// serde.go
// serializing and deserializing Schema

type serdeSchema struct {
	Name        string       `json:"name"`
	Middlewares []Middleware `json:"middlewares"`
	Routes      []*Route     `json:"routes"`
}

func (s Schema) MarshalJSON() ([]byte, error) {
	ss := serdeSchema{
		Name:        s.name,
		Middlewares: s.middlewares,
	}
	for _, b := range s.builders {
		ss.Routes = append(ss.Routes, b.r)
	}
	return json.Marshal(ss)
}

func (s *Schema) UnmarshalJSON(data []byte) error {
	var ss serdeSchema
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	s.name, s.middlewares, s.routes = ss.Name, ss.Middlewares, ss.Routes
	for _, r := range s.routes {
		s.builders = append(s.builders, &RouteBuilder{r: r})
	}
	return nil
}
