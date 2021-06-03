package jsonutil

// Useful JSON raw messages.
var (
	EmptyArray  = RawMessage(`[]`)
	EmptyObject = RawMessage(`{}`)
	True        = RawMessage(`true`)
	False       = RawMessage(`false`)
	Null        = RawMessage(`null`)
)
