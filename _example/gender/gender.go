//go:generate enumeration

package gender

//enumeration: gender
//enumeration-lookup: name, display, abbreviation
//enumeration-encoding: name
//enumeration-string: display
var genderData = []struct {
	name         string
	display      string
	abbreviation string
	comment      string
}{
	{"NIL", "None", "N", "Nil (uninitialied) gender value"},
	{"FEMALE", "Female", "F", "A gender with two X chromosomes"},
	{"MALE", "Male", "M", "A gender with one X and one Y chromosome"},
	{"OTHER", "Other", "O", "A gender which chooses not to identify with either of the above"},
}
