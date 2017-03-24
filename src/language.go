package main

// Language - A programming language type 
type Language struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Experience int    `json:"experience"`
	Difficulty int    `json:"difficulty"`
}

// Languages - An array to hold all available language types
type Languages []Language
