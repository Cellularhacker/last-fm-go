package lastFm

import "errors"

var ErrNotInitialized = errors.New("not initialized")
var ErrNoMethodSpecified = errors.New("no method supplied for lastFm")
