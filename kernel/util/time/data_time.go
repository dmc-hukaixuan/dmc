package time

import (
	"time"
)

// type datatime struct{}

// var DataTime = new(datatime)

//returns the number of non-leap seconds since what ever time the
//system considers to be the epoch (that's 00:00:00, January 1, 1904
//for Mac OS, and 00:00:00 UTC, January 1, 1970 for most other systems).
func SystemTime() int64 {
	// get user preference time zone
	// should be check system default time zone
	// return system time according to the time zone in the system configuration
	return time.Now().Unix()
}

// returns a time stamp for a given system time in C<yyyy-mm-dd 23:59:59> format.
//
//
//
// If you need the short format "23:59:59" for dates that are "today",
// pass the Type parameter like this:
func SystemTime2TimeStamp(system_time int64) string {
	t := time.Unix(system_time, 0)
	return t.Format("2006-01-02 15:04:05")
}

// returns a time stamp of the local system time (see L<SystemTime()>)
// in C<yyyy-mm-dd 23:59:59> format.
func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// converts a system time to a structured date array.
//
func SystemTime2Date() {

}

// converts a given time stamp to local system time.
func TimeStamp2SystemTime(time_stamp string) int64 {
	loc, _ := time.ParseInLocation("2006-01-02 15:04:05", time_stamp, time.Local)
	return loc.Unix()
}

// converts a structured date array to system time of system timezone
func Date2SystemTime() {

}
