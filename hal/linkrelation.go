// go2hal v0.1.0
// Copyright (c) 2016 Patrick Moule
// License: MIT

package hal

import "errors"

// A link relation provides a descriptive name to add a meaning to
// links. To create a more discoverable API, a link relation
// can optionally be prefixed with a CURIE name.
type LinkRelation interface {
	Name() string
	FullName() string
	SetCurieLink(curieLink *LinkObject)
	CurieLink() LinkObject
}

// An unexported implementation of the LinkRelation interface.
type linkRelation struct {
	name      string
	curieLink *LinkObject
}

// Initializer for a valid link relation.
func NewLinkRelation(name string) (LinkRelation, error) {
	if name == "" {
		return nil, errors.New("LinkRelation requires a name value.")
	}

	return &linkRelation{name: name}, nil
}

// Returns the assigned name.
func (lr *linkRelation) Name() string {
	return lr.name
}

// Returns the assigned name. In case of preceding CURIE link assignment
// the returned name is prefixed with the CURIE's name.
func (lr *linkRelation) FullName() string {
	if lr.curieLink == nil {
		return lr.Name()
	}

	return lr.curieLink.Name + ":" + lr.Name()
}

// Use CURIES to create a more discoverable API by assigning
// a CURIE link.
func (lr *linkRelation) SetCurieLink(curieLink *LinkObject) {
	lr.curieLink = curieLink
}

// Returns the assigned CURIE link.
func (lr *linkRelation) CurieLink() LinkObject {
	return *lr.curieLink
}