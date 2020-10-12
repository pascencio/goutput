package out

const levelErrorValue int = 3
const levelWarnValue int = 2
const levelInfoValue int = 1
const levelDebugValue int = 0

const levelErrorName string = "ERROR"
const levelWarnName string = "WARN"
const levelInfoName string = "INFO"
const levelDebugName string = "DEBUG"

const levelErrorArg string = "-e"
const levelWarnArg string = "-w"
const levelInfoArg string = "-i"
const levelDebugArg string = "-d"

// Level ...
type Level interface {
	GetName() string
	GetValue() int
}

// LevelDebug ...
type LevelDebug struct {
}

// GetName ...
func (i LevelDebug) GetName() string {
	return levelDebugName
}

// GetValue ...
func (i LevelDebug) GetValue() int {
	return levelDebugValue
}

// LevelInfo ...
type LevelInfo struct {
}

// GetName ...
func (i LevelInfo) GetName() string {
	return levelInfoName
}

// GetValue ...
func (i LevelInfo) GetValue() int {
	return levelInfoValue
}

// LevelWarning ...
type LevelWarning struct {
}

// GetName ...
func (i LevelWarning) GetName() string {
	return levelWarnName
}

// GetValue ...
func (i LevelWarning) GetValue() int {
	return levelWarnValue
}

// LevelError ...
type LevelError struct {
}

// GetName ...
func (i LevelError) GetName() string {
	return levelErrorName
}

// GetValue ...
func (i LevelError) GetValue() int {
	return levelErrorValue
}

// StringToLevel ...
func StringToLevel(v string) (l Level, e bool) {
	e = true
	switch v {
	case levelErrorName:
		l = LevelError{}
	case levelWarnName:
		l = LevelWarning{}
	case levelInfoName:
		l = LevelInfo{}
	case levelDebugName:
		l = LevelDebug{}
	default:
		e = false
	}
	return l, e
}

// ArgumentToLevel ...
func ArgumentToLevel(a map[string][]string) (l Level, e bool) {
	e = true
	_, ok := a[levelErrorArg]
	if ok {
		l = LevelError{}
		return l, e
	}
	_, ok = a[levelWarnArg]
	if ok {
		l = LevelWarning{}
		return l, e
	}

	_, ok = a[levelInfoArg]
	if ok {
		l = LevelInfo{}
		return l, e
	}

	_, ok = a[levelDebugArg]
	if ok {
		l = LevelDebug{}
		return l, e
	}
	e = false
	return l, e
}
