package number

import (
    "dmc/global"
    "fmt"
    "math/rand"
    "strconv"
    "time"
    "unsafe"
)

func TicketNumberCounterAdd() int64 {
    type numberCounter struct {
        ID         int64     `json:"id" gorm:"primaryKey;autoIncrement;"` // 工单的 ID
        Counter    int64     `json:"counter" gorm:"comment:api中文描述"`      // 工单的单号
        CounterUID string    `json:"counter_uid" gorm:"comment:api组"`     // 工单的标题
        CreateTime time.Time `json:"create_time" gorm:"comment:api组"`
    }
    // get uuid
    counterUID := strconv.FormatInt(time.Now().UnixNano(), 10) + GetUID(12)
    // get now time .Format("2022-01-02 15:04:05")
    currentTimeString := time.Now()

    // Insert new ticket counter into the database (with value 0)
    resultok := global.GVA_DB.Table("ticket_number_counter").Create(&numberCounter{
        Counter:    0,
        CounterUID: counterUID,
        CreateTime: currentTimeString,
    })
    if resultok.Error != nil {
        fmt.Println("add working error ", resultok.Error)
    }

    // It's strange, but this sleep seems to be needed to make sure that other database sessions also see this record.
    //   Without it, there were race conditions because the fillup of unset values below didn't find records that other
    //   sessions already inserted.
    time.Sleep(5 * time.Nanosecond)

    var nc numberCounter
    // Get the ID of the just inserted ticket counter.
    global.GVA_DB.Table("ticket_number_counter").Where("counter_uid = ?", counterUID).First(&nc)

    // Calculate the counter values for all records that don't have a generated value yet.
    //   This is safe even if multiple processes access the records at the same time.
    var unsetCounterIDs []numberCounter
    global.GVA_DB.Table("ticket_number_counter").Where("counter = 0 AND id <= ? ", nc.ID).Order("id ASC").Find(&unsetCounterIDs)

    var setOffset int
    for _, unsetID := range unsetCounterIDs {
        var tmpnc numberCounter
        // Get previous counter record value (tolerate gaps).
        previousCounter := 0
        global.GVA_DB.Table("ticket_number_counter").Where("id < ?", unsetID.ID).Order("id DESC").First(&tmpnc)
        previousCounter = int(tmpnc.Counter)

        // Offset must only be set once (following are consecutive)
        newCounter := previousCounter + 1
        if setOffset == 0 {
            newCounter = previousCounter + 1
            setOffset = 1
        }
        fmt.Println("unsetID :", unsetID, "previousCounter ", previousCounter)
        // Update the counter value, unless another process already did it.
        global.GVA_DB.Table("ticket_number_counter").Where("id = ? AND counter = 0", unsetID.ID).Update("counter", newCounter)
    }
    var ncc numberCounter
    // Get the just inserted ticket counter with the now computed value.
    global.GVA_DB.Table("ticket_number_counter").Where("counter_uid = ?", counterUID).First(&ncc)
    fmt.Println("ht ", ncc)
    //
    return ncc.Counter
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

/*
   @desc: Generates a unique identifier.
   @param:
   @return: a unique string
*/
func GetUID(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return *(*string)(unsafe.Pointer(&b))
}
