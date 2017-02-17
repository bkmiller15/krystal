package printout

import (
	"fmt"
	//"errors"
	"../ds"
	"github.com/jroimartin/gocui"
)

func Timers(g *gocui.Gui) error {
	var err error
	var timerView *gocui.View

	if timerView, err = g.View("timersview"); err != nil {
		return err
	}

	timerView.Clear()

	// Pomodoro Timer
	fmt.Fprintln(timerView, "")
	fmt.Fprintf(timerView, " \x1b[0;34mPomodoro: \x1b[0;37m")
	
	if ds.PomTimer.Break {
		fmt.Fprintf(timerView, "   \x1b[0;34mBreak !!!  \x1b[0;37m\n")
	} else {
		fmt.Fprintf(timerView, "   \x1b[0;34mWork Time \x1b[0;37m\n")
	}

	fmt.Fprintf(timerView, " ")
	print_Rotator(timerView, ds.PomTimer.Rotator)
	fmt.Fprintf(timerView, " ")
	fmt.Fprintf(timerView, " ")

	if ds.PomTimer.State {
		if ds.PomTimer.Rotator == 0 { ds.PomTimer.Rotator = 10 }
		ds.PomTimer.Rotator--
		

		fmt.Fprintf(timerView, "Running ")
	} else {
		fmt.Fprintf(timerView, "Paused  ")
	}

	fmt.Fprintf(timerView, " ")
	fmt.Fprintf(timerView, " [ %0.2d:%0.2d ]\n", ds.PomTimer.Minutes, ds.PomTimer.Seconds)	
	fmt.Fprintln(timerView, "")
	fmt.Fprintf(timerView, " Total Intervals: %d\n", ds.PomTimer.Intervals)	
	fmt.Fprintf(timerView, " Total Work Mins: %d\n", (ds.PomTimer.Intervals*25))	
	fmt.Fprintf(timerView, " Total Break Mins: %d\n", (ds.PomTimer.Intervals*5))	
	fmt.Fprintln(timerView, " ----------------------")

	//Task Timer
	fmt.Fprintf(timerView, " \x1b[0;34mTask Timer: \x1b[0;37m\n")

	fmt.Fprintf(timerView, " ")
	print_Rotator(timerView, ds.TaskTimer.Rotator)
	fmt.Fprintf(timerView, " ")
	fmt.Fprintf(timerView, " ")

	if ds.TaskTimer.State {
		ds.TaskTimer.Rotator++
		if ds.TaskTimer.Rotator >= 10 { ds.TaskTimer.Rotator = 0 }

		fmt.Fprintf(timerView, "Running ")
	} else {
		fmt.Fprintf(timerView, "Paused  ")
	}

	fmt.Fprintf(timerView, " ")
	fmt.Fprintf(timerView, " [ %0.2d:%0.2d ]\n", ds.TaskTimer.Minutes, ds.TaskTimer.Seconds)	
	fmt.Fprintln(timerView, "")
	fmt.Fprintf(timerView, " Same Task Mins: %d\n", ds.TaskTimer.TotalTaskTime)

	timeTotal := 0
	for index := range ds.FilteredEntries {
		timeTotal +=  ds.FilteredEntries[index].Duration
	}
	
	fmt.Fprintf(timerView, " Overview Mins: %d\n", timeTotal)
	


	return nil
}


func print_Rotator(v *gocui.View,rot int) {

	switch rot {
	case 0:
		fmt.Fprintf(v, "〇")
	case 1:
		fmt.Fprintf(v, "一")
	case 2:
		fmt.Fprintf(v, "二")
	case 3:
	    	fmt.Fprintf(v, "三")
	case 4:
	    	fmt.Fprintf(v, "四")
	case 5:
	    	fmt.Fprintf(v, "五")
	case 6:
	    	fmt.Fprintf(v, "六")
	case 7:
	    	fmt.Fprintf(v, "七")
	case 8:
	    	fmt.Fprintf(v, "八")
	case 9:
	    	fmt.Fprintf(v, "九")
	case 10:
	    	fmt.Fprintf(v, "十")
	}
}



func timerview_Title(g *gocui.Gui) error {
	var err error
	var v *gocui.View

	if v, err = g.View("timersview"); err != nil {
		return err
	}

	v.Title = "[ Timers ]"

	return nil
}

