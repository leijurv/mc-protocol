// Code generated by "stringer -type=State"; DO NOT EDIT

package protocol

import "fmt"

const _State_name = "HandshakingPlayStatusLogin"

var _State_index = [...]uint8{0, 11, 15, 21, 26}

func (i State) String() string {
	if i < 0 || i >= State(len(_State_index)-1) {
		return fmt.Sprintf("State(%d)", i)
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}
