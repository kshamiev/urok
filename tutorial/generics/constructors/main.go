package main

import "fmt"

func main() {
	fmt.Printf("%T\n", NewHasT[Unique]())
	fmt.Printf("%T\n", NewHasT[UniqueName]())
}

// HasID is a structural constraint satisfied by structs with a single field
// called "ID" of type "string".
type HasID interface {
	~struct {
		ID string
	}
	GetID() string
}

// CanGetID is an interface constraint satisfied by a type that has a function
// with the signature "GetID() string".
type CanGetID interface {
	GetID() string
}

// CanSetID is an interface constraint satisfied by a type that has a function
// with the signature "SetID(string)".
type CanSetID interface {
	SetID(string)
}

// Unique satisfies the structural constraint "HasID" *and* the interface
// constraint "CanGetID."
type Unique struct {
	ID string
}

func (u Unique) GetID() string {
	return u.ID
}

// UniqueName does *not* satisfiy the structural constraint "HasID," because
// while UniqueName has the field "ID string," the type also contains the field
// "Name string."
//
// Structural constraints must match *exactly*.
//
// UniqueName *does* satisfy the interface constraint "CanGetID."
type UniqueName struct {
	// Unique
	ID string
}

func (u UniqueName) GetID() string {
	return u.ID
}

func (u *UniqueName) SetID(s string) {
	u.ID = s
}

// //

// NewHasT returns a new instance of T.
func NewHasT[T HasID]() T {
	// Declare a new instance of T on the stack.
	var t T

	// Return the new T by value.
	return t
}

// NewT returns a new instance of T...
//
// ...or it would if this function were not invalid. Composite constraints
// cannot contain unions of concrete types such as Go primitive or struct
// types and interface types.
// func NewT[T HasID | CanGetID]() T {
// 	var t T
// 	return t
// }

// NewCanGetT returns a new instance of T.
func NewCanGetT[T CanGetID]() T {
	// Declare a new instance of T on the stack.
	var t T

	// Return the new T by value.
	return t
}

// NewCanSetT returns a new instance of T.
func NewCanSetT[T CanSetID]() T {
	// Declare a new instance of T. Because T is constrained to be a
	// concrete type, it can easily be declared on the stack.
	var t T

	// Return the new T by value.
	return t
}
