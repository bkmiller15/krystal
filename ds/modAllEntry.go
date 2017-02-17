package ds

import (
	"sort"
	"regexp"
	"strings"
)

func Delete_Entry(i int) {
	AllEntries = append(AllEntries[:i], AllEntries[i+1:]...)
}


func Set_Toggle_Enabled_Entry(index int) {
	if AllEntries[index].Enabled {
		AllEntries[index].Enabled = false
	} else {
		AllEntries[index].Enabled = true
	}
}


func Set_Title_Entry(index int, title string) {
	AllEntries[index].Title = title
}

func Set_AOF_Entry(index int, aof string) {
	AllEntries[index].AOF = aof
}


func Set_Category_Entry(index int, category string) {
	AllEntries[index].Category = category
}

func Set_Topic_Entry(index int, topic string) {
	AllEntries[index].Topic = topic
}

func Set_Duration_Entry(index int, dur int) {
	AllEntries[index].Duration = dur
}

func Set_Content_Entry(index int, content string) {
	AllEntries[index].Content = content

	AllEntries[index].People = nil
	
	regexFile := regexp.MustCompile(`(^@\w+|(?: )@\w+|(?:\n)@\w+)`)
	peopleA := regexFile.FindAllString(content, -1)
	
	for i := range peopleA {
		peopleA[i] = strings.Replace(peopleA[i], "\n", "", -1)
		peopleA[i] = strings.Replace(peopleA[i], " ", "", -1)
	}


	AllEntries[index].People = peopleA
}



func Add_Attachment_Entry(index int, attachments string) {
	AllEntries[index].Attachments = append(AllEntries[index].Attachments, attachments)
}





func Set_Date_Entry(index int, date string) {
	var year, month, day string

	year = date[6:len(date)-1]
	month = date[0:2]
	day = date[3:5]

	AllEntries[index].Number = Get_Next_Number(year, month, day)

	AllEntries[index].Year = year
	AllEntries[index].Month = month
	AllEntries[index].Day = day



}





func Get_Entry_Index(year, month, day string, number int) int {
	for i := range AllEntries {
		if AllEntries[i].Year == year && AllEntries[i].Month == month && AllEntries[i].Day == day && AllEntries[i].Number == number {
			return i
		}
	}
	return 0
}


func Get_FilteredEntry_Info(i int) (year, month, day string, number int) {
	return FilteredEntries[i].Year, FilteredEntries[i].Month, FilteredEntries[i].Day, FilteredEntries[i].Number
}


func Add_New_Entry(year, month, day, title string) {
	var newEntry Entry
	
	newEntry.Enabled = true
	newEntry.Title = title
	newEntry.Year  = year
	newEntry.Month = month
	newEntry.Day   = day
	newEntry.Number = Get_Next_Number(year, month, day)


	AllEntries = append(AllEntries, newEntry)
	sort.Sort(AllEntries)

}



func Get_Next_Number(year, month, day string) int {
	var number int = 0

	for i := range AllEntries {
		if AllEntries[i].Year == year && AllEntries[i].Month == month && AllEntries[i].Day == day {
			if AllEntries[i].Number > number {
				number = AllEntries[i].Number
			}
			
		}
	}
	number++
	return number
}





