package stringers

import "fmt"

type IPAddr [4]byte

func (v IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", v[0], v[1], v[2], v[3])
}
