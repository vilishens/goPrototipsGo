package utils

import (
	"time"
)

/*
func MansWTicker(d time.Duration) (tick *time.Ticker) {
	if d <= 0 {
		return nil
	}

	return time.NewTicker(d)
}
*/

func TimeNow(format string) (str string) {
	return time.Now().Format(format)
}

/*
func DurationOf3PartMinimizeStr(str string) (min string) {

	parts := strings.Split(str, ":")

	min = ""
	for k, v := range parts {
		if "00" == v {
			min += ""
		} else if '0' == v[0] {
			min += string(v[1])
		} else {
			min += v
		}

		if k < 2 {
			min += ":"
		}
	}

	return
}

func DurationOf3PartStr(str string) (t time.Duration, err error) {

	parts := strings.Split(str, ":")

	if 3 != len(parts) {
		err = ErrFuncLine(fmt.Errorf("Thera are not 3 parts in string '%s' but %d", str, len(parts)))
		return
	}

	timeStr := ""
	vals := []string{"h", "m", "s"}

	for i, v := range parts {
		timeStr += v
		if "" != v {
			timeStr += vals[i]
		}
	}

	if t, err = time.ParseDuration(timeStr); nil != err {
		err = ErrFuncLine(fmt.Errorf("Thera are not 3 parts in string '%s' but %d", str, len(parts)))
		return
	}

	t = t.Round(time.Second)

	return
}

func DurationTo3PartStr(dur time.Duration, fillZero bool) (str string) {

	form := "%s:%s:%s"
	if fillZero {
		form = "%02s:%02s:%02s"
	}

	t := dur.Round(time.Second).String()

	seq := []string{"h", "m", "s"}
	vals := []string{"", "", ""}

	var spl []string
	for i, v := range seq {
		spl = strings.Split(t, v)

		//		fmt.Printf("STR %s ARR %v\n", t, spl)

		if 2 == len(spl) {
			vals[i] = spl[0]
			t = spl[1]
		}
	}
	vals[2] = spl[0]

	str = fmt.Sprintf(form, vals[0], vals[1], vals[2])

	return
}
*/
