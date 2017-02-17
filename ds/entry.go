package ds

import (
	//"time"
)

var AllEntries EntryList

var FilteredEntries []Entry

var UserFilterString string

var ProgFilterString string

type EntryList []Entry


type Entry struct {
	Enabled		bool		`json:"Enable"`
	Number		int
	Title		string		`json:"Title"`
	AOF		string		`json:"AOF"`
	Category	string		`json:"Category"`
	Topic		string
	People		[]string	`json:"People"`
	Month		string
	Day		string
	Year		string
	EstDuration	int		//`json:"Duration"`
	Duration	int		`json:"Duration"`
	Content		string		`json:"Content"`
	Attachments	[]string	`json:"Attachments"`
}





func (e EntryList) Len() int {
	return len(e)
}

func (e EntryList) Less(i, j int) bool {
	if e[i].Year > e[j].Year {
		return true
	} else if e[i].Year < e[j].Year {
		return false
	} else {
		if e[i].Month > e[j].Month {
			return true
		} else if e[i].Month < e[j].Month {
			return false
		} else {
			if e[i].Day > e[j].Day {
				return true
			} else if e[i].Day < e[j].Day {
				return false
			} else {
				if e[i].Number > e[j].Number {
					return true
				} else if e[i].Number < e[j].Number {
					return false
				} else {
					return false
				}
			}
		}
	}
}

func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}


// FormattedCreatedAt formats the CreatedAt field into a human readable format.
//func (entry *Entry) Update_CreatedAt_Time() {
//	entry.CreatedAt.Format(time.RFC822)
//}

// FormattedUpdatedAt formats the UpdatedAt field into a human readable format.
//func (entry *Entry) Update_UpdatedAt_Time() {
//	entry.UpdatedAt.Format(time.RFC822)
//}












func Get_AllEntries_On_Date(year, month, day string) []Entry {

	var subEntries []Entry

	for i := range AllEntries {
		if AllEntries[i].Year == year && AllEntries[i].Month == month && AllEntries[i].Day == day {
			subEntries = append(subEntries, AllEntries[i])
		}
	}

	return subEntries

}


