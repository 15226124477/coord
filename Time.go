package coord

import (
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	UNIX    = 0
	UTC     = 1
	GPST    = 2
	GPSWeek = 3
	Local   = 4
)

// UnixTime 时间戳
type UnixTime struct {
	Stamp uint64
}

// UtcTime UTC时间
type UtcTime struct {
	UTC time.Time
}

// GpstTime GPS时间
type GpstTime struct {
	GPST time.Time
}

func (gps *GpstTime) GpstString() string {
	return gps.GPST.Format("2006-01-02 15:04:05.000")
}

// GpsWeekTime GPS周
type GpsWeekTime struct {
	GpstWeek  int
	GpsSecond float64
}

// LocalTime 本地时间
type LocalTime struct {
	Local time.Time
}

// MTime 时间
type MTime struct {
	*UnixTime
	*UtcTime
	*GpstTime
	*GpsWeekTime
	*LocalTime
	ConvertBefore int
	ConvertAfter  int
}

// UnixConvert Unix时间戳数据转换
func (t *MTime) UnixConvert() {
	ms := t.UnixTime.Stamp % 1000
	if t.ConvertAfter == Local {
		value := time.Unix(int64(t.UnixTime.Stamp)/1000, 2).Add(1000 * time.Duration(ms) * time.Microsecond)
		t.LocalTime.Local = value
	} else if t.ConvertAfter == UTC {
		value := time.Unix(int64(t.UnixTime.Stamp)/1000, 2).Add(-8 * time.Hour).Add(1000 * time.Duration(ms) * time.Microsecond)
		t.UtcTime.UTC = value
	} else if t.ConvertAfter == GPST {
		value := time.Unix(int64(t.UnixTime.Stamp)/1000, 2).Add(-8 * time.Hour).Add(-18 * time.Second).Add(1000 * time.Duration(ms) * time.Microsecond)
		t.GpstTime.GPST = value
	} else if t.ConvertAfter == GPSWeek {
		value := time.Unix(int64(t.UnixTime.Stamp)/1000, 2).Add(+18*time.Second).Add(1000*time.Duration(ms)*time.Microsecond).Unix() - 315964800
		t.GpsWeekTime.GpstWeek = int(value / 604800)
		t.GpsWeekTime.GpsSecond = float64(value % 604800)
	}
}

// Convert 时间综合转换
func (t *MTime) Convert() {
	switch t.ConvertBefore {
	case UNIX:
		t.UnixConvert()
	}

}

// Value 打印时间值
func (t *MTime) Value() {
	// 时间转换
	t.Convert()
	// 输出值
	switch t.ConvertAfter {
	case UNIX:
		log.Info(0, t.UnixTime.Stamp)
	case UTC:
		log.Info("UTC时:", t.UtcTime.UTC.Format("2006-01-02 15:04:05.000"))
	case GPST:
		log.Info("GPS时:", t.GpstTime.GPST.Format("2006-01-02 15:04:05.000"))
	case GPSWeek:
		log.Info("GPS周:", t.GpsWeekTime.GpstWeek, t.GpsWeekTime.GpsSecond)
	case Local:
		log.Info("本地时间:", t.LocalTime.Local.Format("2006-01-02 15:04:05.000"))
	}

}
