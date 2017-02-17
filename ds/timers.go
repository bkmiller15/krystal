package ds



var PomTimer pTimer

var TaskTimer tTimer


type pTimer struct {
	State		bool
	Break		bool
	Seconds		int
	Minutes		int
	Intervals	int
	Rotator		int
}

type tTimer struct {
	State			bool
	Seconds			int
	Minutes			int
	TotalTaskTime		int
	TotalOverviewTime	int
	Rotator			int
}


func (p *pTimer) Reset() {
	p.Minutes = 25
	p.Seconds = 0
	p.Break = false
	p.Rotator = 0
}

func (p *pTimer) ResetIntervals() {
	p.Intervals = 0
}

func (p *pTimer) ToggleState() {
	p.State = !(p.State)
}


func (p *pTimer) Stop() {
	p.State = false
}

func (p *pTimer) Start() {
	p.State = false
}





func (t *tTimer) Reset() {
	t.Minutes = 0
	t.Seconds = 0
	t.Rotator = 0
}

func (t *tTimer) LoadMin(index int) {
	t.Seconds = 0
	t.Minutes = AllEntries[index].Duration
}

func (t tTimer) SaveMin(index int) {
	AllEntries[index].Duration = t.Minutes
}


func (t *tTimer) ToggleState() {
	t.State = !(t.State)
}


func (t *tTimer) Stop() {
	t.State = false
}

func (t *tTimer) Start() {
	t.State = false
}


