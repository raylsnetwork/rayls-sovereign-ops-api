package domain

import "github.com/google/uuid"

// UserSignupDetails is the profile/lead information the standalone email sign-up form
// collects alongside the email address (company, headcount, referral source, goals).
//
// Deliberately separate from User: these are not identity or credentials, they are only
// ever supplied by the email sign-up path (OAuth and SIWE never collect them), and they
// are optional for everyone else. One row per user — re-submitting the form updates it.
type UserSignupDetails struct {
	Model
	UserID uuid.UUID `json:"userId"`
	// Company is the free-text organisation name typed into the form.
	Company string `json:"company"`
	// Employees is the headcount bucket label as shown in the form ("1-50", "51-200", ...),
	// stored verbatim rather than parsed so re-labelling the UI options cannot corrupt rows.
	Employees string `json:"employees"`
	// HeardAbout is the referral-source option selected in the form.
	HeardAbout string `json:"heardAbout"`
	// Goals is the free-text answer describing what the user wants to build.
	Goals string `json:"goals"`
}
