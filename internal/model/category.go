package model

type Point struct {
	Title string `json:title`
}

type Category struct {
	Id     string  `json:id,omitifempty`
	Title  string  `json:title`
	Points []Point `json:points,omitifempty`
}
