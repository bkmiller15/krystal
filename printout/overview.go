package printout

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"../ds"
	"github.com/jroimartin/gocui"
)

func Overview(g *gocui.Gui) error {
	var yearO, monthO, dayO string
	var year, month, day string
	var yearN, monthN, dayN int
	var maxX int
	var err error
	var v *gocui.View
	var index int
	var minX int = 136

	maxX, _ = g.Size()

	// Print Overview Title Also
	overview_Title(g)

	//Update Array before starting
	ds.Update_FilteredEntries()


	if v, err = g.View("overview"); err != nil {
		return err
	}

	v.Clear()

	if ds.FilteredEntries == nil {
		fmt.Fprintf(v, " No Entries :( ")
	} else {

		for index = range ds.FilteredEntries {

			year = ds.FilteredEntries[index].Year
			month = ds.FilteredEntries[index].Month
			day = ds.FilteredEntries[index].Day

			yearN, _ = strconv.Atoi(year)
			monthN, _ = strconv.Atoi(month)
			dayN, _ = strconv.Atoi(day)

			t := time.Date(yearN, time.Month(monthN), dayN, 0, 0, 0, 0, time.Local)

			if year == yearO && month == monthO && day == dayO {
				if maxX < minX {
					fmt.Fprintf(v, strings.Repeat(" ", 8))
				} else {
					fmt.Fprintf(v, strings.Repeat(" ", 16))
				}
				if is_Next_Date_The_Same(index, year, month, day) {
					fmt.Fprintf(v, " \x1b[0;34m├─ \x1b[0;37m")
				} else {
				
					fmt.Fprintf(v, " \x1b[0;34m└─ \x1b[0;37m")
				}
				if ds.FilteredEntries[index].Enabled == false {
					fmt.Fprintf(v, "[H]")
				}	
			} else {

				fmt.Fprintf(v, "\x1b[0;34m%0.15s\x1b[0;37m", t.Format(" Mon: ") )
				if maxX < minX {
					
					fmt.Fprintf(v, "%0.2s", t.Format("02") )

				} else {
					fmt.Fprintf(v, "%0.10s", t.Format("01/02/2006") + strings.Repeat(" ", 10))
				}


				if is_Next_Date_The_Same(index, year, month, day) {
					fmt.Fprintf(v, " \x1b[0;34m┬─ \x1b[0;37m")
				} else {
				
					fmt.Fprintf(v, " \x1b[0;34m── \x1b[0;37m")
				}
				if ds.FilteredEntries[index].Enabled == false {
					fmt.Fprintf(v, "[H]")
				}

				yearO = year
				monthO = month
				dayO = day
			}


			fmt.Fprintf(v, "%0.55s", ds.FilteredEntries[index].Title + strings.Repeat(" ", 50))
			fmt.Fprintf(v, "  ")
			if ds.KConfig.Mode == 0 {
				fmt.Fprintf(v, "%0.15s", ds.FilteredEntries[index].AOF + strings.Repeat(" ", 50))
				fmt.Fprintf(v, "  ")
			}
			fmt.Fprintf(v, "%0.25s", ds.FilteredEntries[index].Category + strings.Repeat(" ", 25))

			fmt.Fprintf(v, "%0.25s", ds.FilteredEntries[index].Topic + strings.Repeat(" ", 25))
			
			if ds.FilteredEntries[index].Duration > 0 {
				fmt.Fprintf(v, "%dmin  ", ds.FilteredEntries[index].Duration)
			}

			fmt.Fprintf(v, strings.Repeat(" ", maxX))
			fmt.Fprintf(v, "\n")
			
		}
	}

	return nil
}



func is_Next_Date_The_Same(index int, year, month, day string) bool {
	if index < len(ds.FilteredEntries)-1 {
		if year == ds.FilteredEntries[index+1].Year && month == ds.FilteredEntries[index+1].Month && day == ds.FilteredEntries[index+1].Day {
			return true
		} else {
			return false
		}	
	} else {
		return false
	}
}


func overview_Title(g *gocui.Gui) error {
	var err error
	var v *gocui.View

	if v, err = g.View("overview"); err != nil {
		return err
	}

	if ds.UserFilterString == "" {
		v.Title = "[ Overview ]"
	} else {
		v.Title = "[ Overview (Filter: '" + ds.UserFilterString + "') ]"
	}

	return nil
}


