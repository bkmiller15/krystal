package timers

import (
	"time"	
	"fmt"
	"../printout"
	"../ds"
	"github.com/jroimartin/gocui"
)



func Init_Timers(g *gocui.Gui) {

	go func(g *gocui.Gui) {
		ticker := time.NewTicker(time.Second)

	    	for {
			if ds.PomTimer.State {
				if ds.PomTimer.Seconds > 0 {
					ds.PomTimer.Seconds--
				} else {
					ds.PomTimer.Seconds = 59
					if ds.PomTimer.Minutes > 0 {
						ds.PomTimer.Minutes--
					} else {
						fmt.Print("\a")
						if ds.PomTimer.Break == false {
							ds.PomTimer.Break = true
							ds.PomTimer.Minutes = 5 
							ds.PomTimer.Seconds = 0
						} else {
							ds.PomTimer.Intervals++
							ds.PomTimer.Reset()
						}
					}
				}
			}



			if ds.TaskTimer.State {
				ds.TaskTimer.Seconds++
				if ds.TaskTimer.Seconds > 59 {
					ds.TaskTimer.Minutes++
					ds.TaskTimer.Seconds = 0
				}
			}

			g.Execute(func(g *gocui.Gui) error {
				return printout.Timers(g)
			})

			<- ticker.C 
		}
		
	   
	}(g)

}


