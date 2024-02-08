package config

import "fmt"

type Attribute struct {
	Name  string
	Emoji string
}

func GenerateSelfIntroduction(config Config) string {
	var attributes = []Attribute{
		{"Name", "👤"},
		{"Birthday", "🎂"},
		{"Location", "🌍"},
		{"Workplace", "🏢"},
		{"Role", "👔"},
		{"Github", "🐙"},
		{"Email", "📧"},
		{"Programming", "👨‍💻"},
		{"Hobby", "🎮"},
	}

	// Construct the self-introduction card
	card := "┏━━━━━━━━━━━━━━━━━━━━┓\n"
	for _, attribute := range attributes {
		card += fmt.Sprintf("┃ %s %s \n", attribute.Emoji, config.Get(attribute.Name))
	}
	card += "┗━━━━━━━━━━━━━━━━━━━━┛"

	return card
}
