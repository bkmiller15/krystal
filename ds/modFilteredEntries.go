package ds

import (
	"strings"
	"regexp"
)


func Update_FilteredEntries() int {
	var count int

	FilteredEntries = nil


	for i := range AllEntries {
		if prog_Filter_Test(i) && user_Filter_Test(i) {
			FilteredEntries = append(FilteredEntries, AllEntries[i])
			count++
		}
	}


	return count
}


func prog_Filter_Test(index int) bool {
	if filter_Category(ProgFilterString, index) == false { return false }
	if filter_AOF(ProgFilterString, index) == false { return false }
	return true
}


func user_Filter_Test(index int) bool {

	if AllEntries[index].Enabled == false && KConfig.ShowDisabled == false { return false }

	if filter_People(UserFilterString, index) == false { return false }
	if filter_Category(UserFilterString, index) == false { return false }
	if filter_AOF(UserFilterString, index) == false { return false }
	if filter_Topic(UserFilterString, index) == false { return false }

	if filter_Words(UserFilterString, index) == false { return false }	

	return true
}





func filter_People(filter string, index int) bool {

	regexPeople := regexp.MustCompile(`(^@\w+|(?: )@\w+|(?:\n)@\w+)`)
	peopleS := regexPeople.FindString(filter)
	
	peopleS = strings.Replace(peopleS, "\n", "", -1)
	peopleS = strings.Replace(peopleS, " ", "", -1)


	if peopleS == "" { return true }

	for i := range AllEntries[index].People {
		if AllEntries[index].People[i] == peopleS { return true }
	}

	return false
}


func filter_Category(filter string, index int) bool {

	regexCat := regexp.MustCompile(`(^\*\w+|(?: )\*\w+|(?:\n)\*\w+)`)
	categoryS := regexCat.FindString(filter)
	
	categoryS = strings.Replace(categoryS, "\n", "", -1)
	categoryS = strings.Replace(categoryS, " ", "", -1)
	categoryS = strings.Replace(categoryS, "*", "", -1)


	if categoryS == "" { return true }

	if AllEntries[index].Category == categoryS { return true }

	return false
}



func filter_AOF(filter string, index int) bool {

	regexAOF := regexp.MustCompile(`(^\^\w+|(?: )\^\w+|(?:\n)\^\w+)`)
	AOFS := regexAOF.FindString(filter)
	
	AOFS = strings.Replace(AOFS, "\n", "", -1)
	AOFS = strings.Replace(AOFS, " ", "", -1)
	AOFS = strings.Replace(AOFS, "^", "", -1)

	if AOFS == "" { return true }

	if AllEntries[index].AOF == AOFS { return true }

	return false
}

func filter_Topic(filter string, index int) bool {

	regexTopic := regexp.MustCompile(`(^#\w+|(?: )#\w+|(?:\n)#\w+)`)
	topic := regexTopic.FindString(filter)
	
	topic = strings.Replace(topic, "\n", "", -1)
	topic = strings.Replace(topic, " ", "", -1)
	topic = strings.Replace(topic, "#", "", -1)

	if topic == "" { return true }

	if AllEntries[index].Topic == topic { return true }

	return false
}







func filter_Words(filter string, index int) bool {
	filterString := filter

	// Matching Strings
	regexEmpty := regexp.MustCompile(`^\s+$`)
	regexWords := regexp.MustCompile(`[a-zA-Z0-9]+`)
	regexAOF := regexp.MustCompile(`(^\^\w+|(?: )\^\w+|(?:\n)\^\w+)`)
	regexCat := regexp.MustCompile(`(^\*\w+|(?: )\*\w+|(?:\n)\*\w+)`)
	regexPeople := regexp.MustCompile(`(^@\w+|(?: )@\w+|(?:\n)@\w+)`)
	regexTopic := regexp.MustCompile(`(^#\w+|(?: )#\w+|(?:\n)#\w+)`)

	// Remove all strings
	filterString = regexAOF.ReplaceAllString(filterString, "")
	filterString = regexCat.ReplaceAllString(filterString, "")
	filterString = regexPeople.ReplaceAllString(filterString, "")
	filterString = regexTopic.ReplaceAllString(filterString, "")

	if filterString == "" || regexEmpty.MatchString(filterString) {	return true }

	
	allWords := regexWords.FindAllString(filterString, -1)

	for i := range allWords {
		if strings.Contains(AllEntries[index].Content, allWords[i]) {
			return true
		} 
		if strings.Contains(AllEntries[index].Title, allWords[i]) {
			return true
		} 
	}
	

	return false
}



