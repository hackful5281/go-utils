package utilx

import (
	"errors"
	"regexp"
	"time"
)

// GetConstellation 根据日期获取星座
func GetConstellation(date string) (string, error) {
	const fmtDate = "20060102"
	if res := checkDate(date); !res {
		return "", errors.New("日期格式异常")
	}

	t, err := time.Parse(fmtDate, date)
	if err != nil {
		//fmt.Println("日期格式解析错误")
		return "", err
	}
	arr1 := []int{21, 20, 21, 21, 22, 22, 23, 24, 24, 24, 23, 22}
	arr2 := []string{"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	if t.Day() < arr1[t.Month()-1] {
		return arr2[t.Month()-1], nil
	} else {
		return arr2[t.Month()], nil
	}
}

func checkDate(date string) bool {
	pattern := `\d{8}`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(date)
}
