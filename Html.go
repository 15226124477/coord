package coord

type LostInterval struct {
	Key       int
	StartTime string
	EndTime   string
	LostCount float64
}

type FixInterval struct {
	Key     int
	GPST    string
	N       string
	E       string
	Z       string
	DeltaN  string
	DeltaE  string
	DeltaZ  string
	DeltaNE string
	IsErr   bool
}

type LostHtmlFormat struct {
	Name       string
	OutTime    string
	Start      string
	End        string
	Epoch      int64
	AllEpoch   float64
	LostEpoch  float64
	Intergrity float64
	Duration   float64
	Items      []LostInterval
}

type FixHtmlFormat struct {
	Name       string
	OutTime    string
	Fix        int64
	FixErr     int64
	ErrV       float64
	ErrH       float64
	FixErrRate string
	Mode       string
	RefN       string
	RefE       string
	RefZ       string
	Items      []FixInterval
}
