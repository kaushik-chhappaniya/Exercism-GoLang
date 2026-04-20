package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// layout := "2006-01-02 15:04:05 +0000 UTC"
    layout := "1/2/2006 15:04:05"
	t,_ := time.Parse(layout,date)
	
    // if err!=nil{
    //     return nil
    // }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    t, _ := time.Parse("January 2, 2006 15:04:05", date)
    // if err!=nil{
    //     return err
    // }
    // time.Now().After() returns true if the time t input is after the current time. Meaning in future will happen. 
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false // fail fast; invalid input is not afternoon
	}
	hour := t.Hour()
    return hour>=12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		return fmt.Sprintf("Err in parsing time:%v",err)
	}
	// Format into a readable output
	return "You have an appointment on " + t.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {    
	anniversary_date, _ := time.Parse("01-02", "09-15")
	month := anniversary_date.Month()
	day := anniversary_date.Day()
	return time.Date(time.Now().Year(), month, day, 0, 0, 0, 0, time.UTC)
	
}
