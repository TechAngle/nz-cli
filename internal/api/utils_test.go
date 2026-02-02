package api

import "testing"

func TestPayloadDatesValidation(t *testing.T) {
	// default seq
	validStartDate := "2026-01-01"
	validEndDate := "2026-01-02"

	// reversed
	validStartDate2 := "2026-01-02"
	validEndDate2 := "2026-01-01"

	// invalid date format
	invalidStardDate := "01-01-2025"
	invalidEndDate := "02-02-2025"

	// function returns first and second dates, but they could be reversed if needed
	f, s, err := ValidatePayloadDates(validStartDate, validEndDate)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if f != validStartDate {
		t.Fatalf("Returned invalid start date: %s", f)
	}

	if s != validEndDate {
		t.Fatalf("Returned invalid end date: %s", f)
	}

	// function returns first and second dates, but they could be reversed if needed
	f2, s2, err := ValidatePayloadDates(validStartDate2, validEndDate2)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if f2 != validEndDate2 {
		t.Fatalf("Returned invalid reversed start date: %s", f)
	}

	if s2 != validStartDate2 {
		t.Fatalf("Returned invalid reversed end date: %s", f)
	}

	// function returns first and second dates, but they could be reversed if needed
	_, _, err = ValidatePayloadDates(invalidStardDate, invalidEndDate)
	if err == nil {
		t.Fatalf("Error NOT occurred but it should be")
	}
}
