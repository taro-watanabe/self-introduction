package config

import "strings"

type Config struct {
	Name        string   `json:"name"`
	Birthday    string   `json:"birthday"`
	Location    string   `json:"location"`
	Workplace   string   `json:"workplace"`
	Role        string   `json:"role"`
	Github      string   `json:"gitHub"`
	Email       string   `json:"email"`
	Programming []string `json:"programming"`
	Hobby       []string `json:"hobby"`
}

func (c Config) Get(attribute string) string {
	attributes := map[string]string{
		"Name":        c.Name,
		"Birthday":    c.Birthday,
		"Location":    c.Location,
		"Workplace":   c.Workplace,
		"Role":        c.Role,
		"Github":      c.Github,
		"Email":       c.Email,
		"Programming": strings.Join(c.Programming, ", "),
		"Hobby":       strings.Join(c.Hobby, ", "),
	}

	value, ok := attributes[attribute]
	if !ok {
		return "Attribute not found"
	}

	return value
}
