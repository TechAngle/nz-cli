package api

import "testing"

func TestInvalidPayloadDates(t *testing.T) {
	// invalid date format
	invalidStardDate := "01-01-2025"
	invalidEndDate := "02-02-2025"

	_, _, err := validatePayloadDates(invalidStardDate, invalidEndDate)
	if err == nil {
		t.Fatalf("Error NOT occurred but it should be")
	}
}

func TestReversedPayloadDates(t *testing.T) {
	// reversed
	validStartDate := "2026-01-02"
	validEndDate := "2026-01-01"

	// function returns first and second dates, but they could be reversed if needed
	startDate, endDate, err := validatePayloadDates(validStartDate, validEndDate)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if startDate != validEndDate {
		t.Fatalf("Returned invalid reversed start date: %s", startDate)
	}

	if endDate != validStartDate {
		t.Fatalf("Returned invalid reversed end date: %s", endDate)
	}
}

func TestPayloadDatesValidation(t *testing.T) {
	// default seq
	validStartDate := "2026-01-01"
	validEndDate := "2026-01-02"

	// function returns first and second dates, but they could be reversed if needed
	startDate, endDate, err := validatePayloadDates(validStartDate, validEndDate)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if startDate != validStartDate {
		t.Fatalf("Returned invalid start date: %s", startDate)
	}

	if endDate != validEndDate {
		t.Fatalf("Returned invalid end date: %s", endDate)
	}
}
