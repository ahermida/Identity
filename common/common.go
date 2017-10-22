/*
        Package common contains functions and types common
        throughout this project.

        A bunch of unlabeled byte arrays will get confusing.
*/
package common

// Hashes should be 32 bytes
type Hash [32]byte

// Get the string representation of the underlying hash
func (h Hash) String() string   {
    return string(h[:])
}

// Explicitly get the internal byte slice
func (h Hash) Bytes() []byte {
    return h[:]
}

// Addresses should be length less than 34, but it's variable
type Address []byte

// Get string of Address
func (a Address) String() string {
    return string(a[:])
}
