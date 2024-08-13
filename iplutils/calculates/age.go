package calculates

import (
	"fmt"
	"time"
)

func calculateAge(birthdate string) (int8, error) {
	// Parse the birthdate string to a time.Time object
	layout := "02-01-2006" //"1994-07-07" //02-01-2006
	birthTime, err := time.Parse(layout, birthdate)
	if err != nil {
		return 0, err
	}

	// Get the current time
	now := time.Now()
	// now = now.AddDate(10, 24, 40)
	// Calculate the age
	age := now.Year() - birthTime.Year()

	// Check if the birthday has occurred this year
	if now.Month() < birthTime.Month() || (now.Month() == birthTime.Month() && now.Day() < birthTime.Day()) {
		age--
	}

	return int8(age), nil
}

func AgeCaculate(birthdate string) (int8, error) {
	age, err := calculateAge(birthdate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return 0, err
	}
	fmt.Printf("Age for birthdate %s is %d years.\n", birthdate, age)
	return age, err
}
