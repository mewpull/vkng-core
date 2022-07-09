package common

import "strings"

type flags interface {
	~int32 | ~uint32
}

// FlagStringMapping is used as a base type for many bitflag enums in vkngwrapper.
// It has the capability to Register flags with a descriptive string in an init() method.
// Once that has done, the flag type can be stringified into a pipe-separated list of
// flags.
type FlagStringMapping[T flags] struct {
	stringValues map[T]string
}

func NewFlagStringMapping[T flags]() FlagStringMapping[T] {
	return FlagStringMapping[T]{make(map[T]string)}
}

func (m FlagStringMapping[T]) Register(value T, str string) {
	m.stringValues[value] = str
}

func (m FlagStringMapping[T]) FlagsToString(value T) string {
	if value == 0 {
		return "None"
	}

	hasOne := false
	var sb strings.Builder

	for i := 0; i < 32; i++ {
		shiftedBit := T(1 << i)
		if value&shiftedBit != 0 {
			strVal, exists := m.stringValues[shiftedBit]
			if exists {
				if hasOne {
					sb.WriteString("|")
				}
				sb.WriteString(strVal)
				hasOne = true
			}
		}
	}

	return sb.String()
}
