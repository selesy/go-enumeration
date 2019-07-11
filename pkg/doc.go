/*
Package enumeration is a Go code generator that creates complex enumerated
types.  Inspiration is taken from Java's ability to include constant
"cargo" data as values that are related to each enumeration member.  The
Java tutorial includes an example of such an enumeration for the solar
system's planets (https://docs.oracle.com/javase/tutorial/java/javaOO/enum.html)

Enumeration Directives:

Enumeration Directives are special comments that provide a specification
that is used to control the generation of the enumeration's code.  It is
possible to generate a basic enumeration with cargo data without using
these directives at all.  Each of the directives is described below.  See
the examples for additional information on how these directives can be
used.

For directives that target the fields of the enumeration's data struct,
there is an alternate format that uses the field's tag.  The results are
the same but field tags are preferred over comment group directives as
there's no chance of specifying a field name that's not part of the
enumeration's struct.  The following two struct's are equivalent (note
the inversion between the keys and the values):

  // Preferred way to specify an enumeration field directive
  type a struct {
	  name string `enumeration:"string"`
  }

  // Alternate way to specify an enumeration field directive
  //enumeration-string: name
  type b struct {
	  name string
  }

enumeration: The "enumeration" directive can be used to specify the Type
name of the generated enumeration.  The type will be defined

enumeration-name: The "enumeration-name" directive is used to indication
which field will be used to generate each enumeration value's constant
name.  Only a single field may be selected as the name.  The default is
to use the "name" field.  An error is thrown and the code generation will
fail if there is no implicit or explicit name field selected.

enumeration-ordinal:

enumeration-zero

enumeration-display:

enumeration-lookup: The "enumeration-lookup" directive is used to indicate
that one or more fields are alternate keys.  As such, fields that can be
used to look up an enumeration's value must have unique values.  More than
one field may be selected as lookup values.  The default behavior is that
only the name field (see enumeration-name above) is included for lookup.

enumeration-string:

enumeration-abbreviation:

enumeration-encoding:

enumeration-encoding-strategy:
*/
package enumeration
