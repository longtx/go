package timer

import "time"

const (
	Nanosecond time.Duration    =   1
	Microsecond                 = 1000 * Nanosecond
	Millisecond                 = 1000 * Microsecond
	Second                      = 1000 * Millisecond
	Minute                      = 60 * Second
	Hour                        = 60 * Minute
)

//用于返回当前本地时间的Time对象，此处的封装主要是为了便于后续对Time对象做进一步处理.
func GetNowTime() time.Time {
	return time.Now()
}

//推算时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

func GetCalculateTime1(currentTimer time.Time, d string) (time.Time, error) {
	return currentTimer.Add(time.Second * 60), nil

}