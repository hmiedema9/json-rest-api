package main

import "fmt"

var currentId int

var languages Languages

// Give us some seed data
func init() {
	RepoCreateLanguage(Language{Name: "Write presentation"})
	RepoCreateLanguage(Language{Name: "Host meetup"})
}

func RepoFindLanguage(id int) Language {
	for _, t := range languages {
		if t.Id == id {
			return t
		}
	}
	// return empty Language if not found
	return Language{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateLanguage(t Language) Language {
	currentId += 1
	t.Id = currentId
	languages = append(languages, t)
	return t
}

func RepoDestroyLanguage(id int) error {
	for i, t := range languages {
		if t.Id == id {
			languages = append(languages[:i], languages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find language with id of %d to delete", id)
}
