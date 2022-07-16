package translation

import (
	"github.com/mostafasolati/t2/english"
	"github.com/mostafasolati/t2/persian"
)

func TranslateWeekDay2(lang string, day int) string {
	switch lang {
	case "fa":
		return persian.TranslateWeekDay(day)
	case "en":
		return english.TranslateWeekDay(day)
	default:
		return "نامشخصی"
	}
}
