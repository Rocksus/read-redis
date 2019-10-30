package userdata

import (
	"time"
)

// Gets the adjusted date of birth to work around leap year differences.
func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthDay >= 60 {
		return birthDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthDay + 1
	}
	return birthDay
}

// Works out if a time.Time is in a leap year.
func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// AgeAt gets the age of an entity at a certain time.
func AgeAt(birthDate time.Time, now time.Time) int {
	// Get the year number change since the player's birth.
	years := now.Year() - birthDate.Year()

	// If the date is before the date of birth, then not that many years have elapsed.
	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years--
	}

	return years
}

// GetUsers returns the list of users
func (h *Handler) GetUsers() []User {
	var users []User
	rows, _ := h.DB.Raw("SELECT user_id, full_name, msisdn, user_email, birth_date FROM ws_user WHERE birth_date IS NOT NULL LIMIT 500").Rows()
	defer rows.Close()
	for rows.Next() {
		var user User
		rows.Scan(&user.UserID, &user.Name, &user.MSISDN, &user.Email, &user.BirthDate)
		user.UserAge = AgeAt(user.BirthDate, time.Now())
		if user.UserAge < 150 {
			users = append(users, user)
		}
	}
	return users
}
