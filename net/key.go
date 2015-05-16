package net

import "github.com/rthornton128/goncurses"

type Key struct {
	K goncurses.Key
	// additional information right here, later on
}

/* Gob transmits type information before it transmits the value, so think of it
   as a "stream" in that each NewEncoder needs to be matched with exactly one
   NewDecoder so that both the encode and decode side can have the correct
   internal state
*/
