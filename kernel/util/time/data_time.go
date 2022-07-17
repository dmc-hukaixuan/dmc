package time

import (
    model "dmc/kernel/model/ticket"
    "dmc/kernel/service/admin/ticket"
    "fmt"
    "strconv"
    "time"

    "gopkg.in/yaml.v3"
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

func SystemTimeString() string {
    // get user preference time zone
    // should be check system default time zone
    // return system time according to the time zone in the system configuration
    return strconv.FormatInt(time.Now().Unix(), 10)
}

func NowSystemTime() time.Time {
    return time.Now()
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

func StringToTime(timeString string) (time.Time, error) {
    loc, _ := time.LoadLocation("Local")
    theTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeString, loc)
    if err != nil {
        return theTime, err
    }
    return theTime, nil
}

// converts a given time stamp to local system time.
func TimeStamp2SystemTime(time_stamp string) int64 {
    loc, _ := time.ParseInLocation("2006-01-02 15:04:05", time_stamp, time.Local)
    return loc.Unix()
}

// converts a structured date array to system time of system timezone
func Date2SystemTime() {

}

/*
   方法说明：计算一段时间的开始时间和结束时间

   入参：
       StartTime ：开始时间,
       StopTime: 结束时间

   返回
       WorkingTime 工作时间的秒数
*/
func WorkingTime(StartTime, StopTime time.Time, calendarID int) (WorkingTime int) {
    if calendarID == 0 {
        //WorkingTime = 0
        return 0
    }
    // get system set time zone
    loc, _ := time.LoadLocation("Asia/Shanghai")
    //fmt.Println("calendarID ", calendarID)
    workingCalender := ticket.WorkingTimeGet(calendarID)
    var working1 model.WorkingCalender
    err1 := yaml.Unmarshal([]byte(workingCalender.WorkingTime), &working1)
    layout := "2006-01-02 15:04:05"
    // 检查开始时间、结束时间格式是否正确
    start_time, err := time.ParseInLocation(layout, StartTime.Format(layout), loc)
    if err != nil {
        fmt.Println("start time format err: %v\n ", err)
        return
    }
    end_time, err1 := time.ParseInLocation(layout, StopTime.Format(layout), loc)
    if err1 != nil {
        fmt.Println("stop time format err1: %v\n ", err1)
        return
    }

    // start time and end time object
    start_time_unix := start_time.Unix()
    start_time_unix_obj := start_time
    end_time_unix := end_time.Unix()

    // 判断开始时间和结束时间的大小
    if start_time_unix >= end_time_unix {
        fmt.Println("starttime : ", StartTime, StopTime)
        fmt.Println("start time greater then end time, return")
        return
    }

    LoopStartTime := time.Now().Unix()

    // 该日历额外的上班日
    VacationDays := make(map[string]string)
    for _, item := range working1.VacationDays {
        for k, v := range item {
            VacationDays[k] = v
        }
    }

    // 该日历的补班日（本来应该休假的日期，结果在补班）
    ExtraWorkingDay := make(map[string]string)
    for _, item := range working1.ExtraWorkingDay {
        for k, v := range item {
            ExtraWorkingDay[k] = v
        }
    }

    Workhours := make(map[string]map[string]string)
    // 将每天的工作时间算出来
    for k, v := range working1.WorkingHours {
        tmp := make(map[string]string)
        for _, value := range v {
            tmp[value] = value
        }
        Workhours[k] = tmp
    }

Loop:
    for start_time_unix < end_time_unix {
        // Fail if this loop takes longer than 5 seconds
        if time.Now().Unix()-LoopStartTime > 5 {
            return
        }
        // the seconds is start time and end time inerval
        var RemainingSeconds int64 = end_time_unix - start_time_unix

        Day := start_time_unix_obj.Day()
        DayName := start_time_unix_obj.Weekday().String()[0:3]
        Hour := start_time_unix_obj.Hour()
        Minute := start_time_unix_obj.Minute()
        Second := start_time_unix_obj.Second()
        YMD := start_time_unix_obj.Format("2006-01-02")
        var IsWorkingDay bool = false
        _, ExtraWorkingDayok := ExtraWorkingDay[YMD]
        worktime, wkhok := Workhours[DayName]

        _, vdok := VacationDays[YMD]
        // 开始时间在补班日期或者不在 休假日期中，则开始时间在工作日时间内
        if ExtraWorkingDayok && wkhok && len(worktime) > 0 {
            IsWorkingDay = true
        } else if !vdok && wkhok && len(worktime) > 0 {
            IsWorkingDay = true
        } else {
            IsWorkingDay = false
        }
        // 检查开始日期是否是 这一天后的 24 小时是否在工作日时间
        // 如果开始时间后的 24 小时在工作时间，working 直接加一天的时间，
        // 防止每半个小时处理消耗时间
        if Hour == 0 && Minute == 0 && Second == 0 {

            next_day_obj := start_time_unix_obj.AddDate(0, 0, 1)
            Epoch := start_time_unix_obj.Unix() + 60*60*24
            oneDaySecond := 60 * 60 * 24
            if RemainingSeconds > int64(oneDaySecond) && Day != next_day_obj.Day() {
                FullDayProcessed := true

                if IsWorkingDay {
                    // 当天的工作时长
                    WorkingHours := len(worktime)
                    WorkingSeconds := WorkingHours * 30 * 60
                    if RemainingSeconds > int64(WorkingSeconds) {
                        WorkingTime += WorkingSeconds
                    } else {
                        FullDayProcessed = false
                    }
                }
                if FullDayProcessed {
                    start_time_unix = Epoch
                    start_time_unix_obj = time.Unix(start_time_unix, 0)
                    continue Loop
                }
            }
        }

        SecondsOfCurrentHour := Minute*60 + Second
        var SecondsToAdd int
        // 把开始时间的分钟数计算成半点时间
        if SecondsOfCurrentHour >= 30*60 {
            SecondsToAdd = (60 * 60) - SecondsOfCurrentHour
        } else {
            SecondsToAdd = (30 * 60) - SecondsOfCurrentHour
        }
        currentTime := strconv.Itoa(Hour)
        // 刚好到半点，将时间格式成半点时间，用户判断该半点是否在工作日时间中
        if Minute == 30 {
            currentTime = currentTime + ":30"
        }

        // 开始时间所在周时间，和小时时间在 working time 中，并且当前时间是工作日
        if _, wkhok := Workhours[DayName][currentTime]; wkhok && IsWorkingDay {
            if int64(SecondsToAdd) > RemainingSeconds {
                SecondsToAdd = int(RemainingSeconds)
            }
            WorkingTime += SecondsToAdd
        }
        Epoch := start_time_unix + int64(SecondsToAdd)
        start_time_unix_obj = time.Unix(Epoch, 0)
        start_time_unix = Epoch
    }
    return
}
