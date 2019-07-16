//go:generate go-enumeration

package gender

//enumeration: Gender
//enumeration-lookup: name
//enumeration-lookup: display
//enumeration-lookup: abbreviation
//enumeration-encoding: name
//enumeration-string: display
var data = []struct {
	name         string `enumeration:"name,lookup,encoding"`
	display      string `enumeration:"lookup,string"`
	abbreviation string `enumeration:"lookup"`
	comment      string
}{
	{"NIL", "None", "N", "Nil (uninitialized) gender value"},
	{"FEMALE", "Female", "F", "A gender with two X chromosomes"},
	{"MALE", "Male", "M", "A gender with one X and one Y chromosome"},
	{"OTHER", "Other", "O", "A gender which chooses not to identify with either of the above"},
}
