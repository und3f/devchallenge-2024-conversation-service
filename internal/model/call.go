package model

type Call struct {
	id             string
	name           string
	location       string
	emotional_tone string
	text           string
	categories     []Category
}
