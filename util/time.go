package util

import "time"

func Parse(tStr string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", tStr)
}

func TryParse(tStr string) time.Time {
	if t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr); err != nil {
		t, _ := time.Parse("2006-01-02 15:04:05", tStr)
		return t
	} else {
		return t
	}
}

func MustParse(tStr string) time.Time {
	if t, err := time.Parse("2006-01-02T15:04:05.000Z", tStr); err != nil {
		t, err := time.Parse("2006-01-02 15:04:05", tStr)
		if err != nil {
			panic(err)
		} else {
			return t
		}
	} else {
		return t
	}
}

// Deprecated: Format
func FormatVal(t time.Time) string {

	if t.IsZero() {
		return ""
	}

	return t.Format("2006-01-02T15:04:05.000Z")
}

// Deprecated: Format
func Format1(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02T15:04:05.000Z")
}

// Deprecated: Format
func FormatV0(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

// Format 此时直接出入 nil 会报错
func Format[T *time.Time | time.Time](t T, format ...string) string {
	var f = "2006-01-02T15:04:05.000Z"
	if len(format) > 0 {
		f = format[0]
	}
	switch v := any(t).(type) {
	case time.Time:
		if v.IsZero() {
			return ""
		}
		return v.Format(f)
	case *time.Time:
		if v == nil || v.IsZero() {
			return ""
		}
		return v.Format(f)
	default:
		return ""
	}
}
