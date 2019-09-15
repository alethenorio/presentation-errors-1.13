package main

// START UNWRAP
func Unwrap() error // HL

// END UNWRAP

// START ISDEF OMIT
// An error is considered to match a target if it is equal to
// that target or it it implements a method Is(error) bool // HL
// such that Is(target) returns true. // HL
// END ISDEF OMIT
// START IS OMIT
func Is(err error) bool // HL
// END IS OMIT

// START ASDEF OMIT
// An error matches target if the error's concrete value is
// assignable to the value pointed to by target, or if the error
// has a method As(interface{}) bool such that As(target)
// returns true. In the latter case, the As method is responsible
// for setting target.
// END ASDEF OMIT
// START AS OMIT
func As(t interface{}) bool // HL
// END AS OMIT
