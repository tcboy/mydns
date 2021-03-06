package utils

import (
	"time"
)

var (
	languageCodeMap = map[string]string{
		"zh":      "CN",
		"zh-Hans": "CN",
		"zh-Hant": "CN",
		"ko-Kore": "KR",
		"ja-Jpan": "JP",
		"ja":      "JP",
		"ko":      "KR",
		"en":      "EN",
		"pt":      "PT",
		"hi":      "IN",
		"in":      "ID",
		"ms":      "MY",
		"th":      "TH",
		"vi":      "VN",
		"ru":      "RU",
	}
	countryTz = map[string]string{
		"BR": "America/Sao_Paulo",
		"CN": "Asia/Shanghai",
		"EN": "Asia/Kuala_Lumpur",
		"ID": "Asia/Jakarta",
		"IN": "Asia/Kolkata",
		"JP": "Asia/Tokyo",
		"KR": "Asia/Seoul",
		"MY": "Asia/Kuala_Lumpur",
		"TH": "Asia/Bangkok",
		"TW": "Asia/Taipei",
		"VN": "Asia/Ho_Chi_Minh",
	}
	countryCodeMap = map[string]string{
		"中国":          "CN",
		"安道尔":         "AD",
		"阿联酋":         "AE",
		"阿富汗":         "AF",
		"安提瓜和巴布达":     "AG",
		"安圭拉":         "AI",
		"阿尔巴尼亚":       "AL",
		"亚美尼亚":        "AM",
		"安哥拉":         "AO",
		"南极洲":         "AQ",
		"阿根廷":         "AR",
		"美属萨摩亚":       "AS",
		"奥地利":         "AT",
		"澳大利亚":        "AU",
		"阿鲁巴":         "AW",
		"奥兰":          "AX",
		"阿塞拜疆":        "AZ",
		"波斯尼亚和黑塞哥维那":  "BA",
		"巴巴多斯":        "BB",
		"孟加拉国":        "BD",
		"比利时":         "BE",
		"布基纳法索":       "BF",
		"保加利亚":        "BG",
		"巴林":          "BH",
		"布隆迪":         "BI",
		"贝宁":          "BJ",
		"圣巴泰勒米":       "BL",
		"百慕大":         "BM",
		"文莱":          "BN",
		"玻利维亚":        "BO",
		"加勒比荷兰":       "BQ",
		"巴西":          "BR",
		"巴哈马":         "BS",
		"不丹":          "BT",
		"布韦岛":         "BV",
		"博茨瓦纳":        "BW",
		"白俄罗斯":        "BY",
		"伯利兹":         "BZ",
		"加拿大":         "CA",
		"科科斯（基林）群岛":   "CC",
		"刚果（金）":       "CD",
		"中非":          "CF",
		"刚果（布）":       "CG",
		"瑞士":          "CH",
		"科特迪瓦":        "CI",
		"库克群岛":        "CK",
		"智利":          "CL",
		"喀麦隆":         "CM",
		"中华人民共和国":     "CN",
		"哥伦比亚":        "CO",
		"哥斯达黎加":       "CR",
		"古巴":          "CU",
		"佛得角":         "CV",
		"库拉索":         "CW",
		"圣诞岛":         "CX",
		"塞浦路斯":        "CY",
		"捷克":          "CZ",
		"德国":          "DE",
		"吉布提":         "DJ",
		"丹麦":          "DK",
		"多米尼克":        "DM",
		"多米尼加":        "DO",
		"阿尔及利亚":       "DZ",
		"厄瓜多尔":        "EC",
		"爱沙尼亚":        "EE",
		"埃及":          "EG",
		"阿拉伯撒哈拉民主共和国": "EH",
		"厄立特里亚":       "ER",
		"西班牙":         "ES",
		"埃塞俄比亚":       "ET",
		"芬兰":          "FI",
		"斐济":          "FJ",
		"福克兰群岛":       "FK",
		"密克罗尼西亚联邦":    "FM",
		"法罗群岛":        "FO",
		"法国":          "FR",
		"加蓬":          "GA",
		"英国":          "GB",
		"格林纳达":        "GD",
		"格鲁吉亚":        "GE",
		"法属圭亚那":       "GF",
		"根西":          "GG",
		"加纳":          "GH",
		"直布罗陀":        "GI",
		"格陵兰":         "GL",
		"冈比亚":         "GM",
		"几内亚":         "GN",
		"瓜德罗普":        "GP",
		"赤道几内亚":       "GQ",
		"希腊":          "GR",
		"南乔治亚和南桑威奇群岛": "GS",
		"危地马拉":        "GT",
		"关岛":          "GU",
		"几内亚比绍":       "GW",
		"圭亚那":         "GY",
		"香港":          "HK",
		"赫德岛和麦克唐纳群岛":  "HM",
		"洪都拉斯":        "HN",
		"克罗地亚":        "HR",
		"海地":          "HT",
		"匈牙利":         "HU",
		"印尼":          "ID",
		"爱尔兰":         "IE",
		"以色列":         "IL",
		"马恩岛":         "IM",
		"印度":          "IN",
		"英属印度洋领地":     "IO",
		"伊拉克":         "IQ",
		"伊朗":          "IR",
		"冰岛":          "IS",
		"意大利":         "IT",
		"泽西":          "JE",
		"牙买加":         "JM",
		"约旦":          "JO",
		"日本":          "JP",
		"肯尼亚":         "KE",
		"吉尔吉斯斯坦":      "KG",
		"柬埔寨":         "KH",
		"基里巴斯":        "KI",
		"科摩罗":         "KM",
		"圣基茨和尼维斯":     "KN",
		"朝鲜":          "KP",
		"韩国":          "KR",
		"科威特":         "KW",
		"开曼群岛":        "KY",
		"哈萨克斯坦":       "KZ",
		"老挝":          "LA",
		"黎巴嫩":         "LB",
		"圣卢西亚":        "LC",
		"列支敦士登":       "LI",
		"斯里兰卡":        "LK",
		"利比里亚":        "LR",
		"莱索托":         "LS",
		"立陶宛":         "LT",
		"卢森堡":         "LU",
		"拉脱维亚":        "LV",
		"利比亚":         "LY",
		"摩洛哥":         "MA",
		"摩纳哥":         "MC",
		"摩尔多瓦":        "MD",
		"黑山":          "ME",
		"法属圣马丁":       "MF",
		"马达加斯加":       "MG",
		"马绍尔群岛":       "MH",
		"马其顿":         "MK",
		"马里":          "ML",
		"缅甸":          "MM",
		"蒙古国":         "MN",
		"澳门":          "MO",
		"北马里亚纳群岛":     "MP",
		"马提尼克":        "MQ",
		"毛里塔尼亚":       "MR",
		"蒙特塞拉特":       "MS",
		"马耳他":         "MT",
		"毛里求斯":        "MU",
		"马尔代夫":        "MV",
		"马拉维":         "MW",
		"墨西哥":         "MX",
		"马来西亚":        "MY",
		"莫桑比克":        "MZ",
		"纳米比亚":        "NA",
		"新喀里多尼亚":      "NC",
		"尼日尔":         "NE",
		"诺福克岛":        "NF",
		"尼日利亚":        "NG",
		"尼加拉瓜":        "NI",
		"荷兰":          "NL",
		"挪威":          "NO",
		"尼泊尔":         "NP",
		"瑙鲁":          "NR",
		"纽埃":          "NU",
		"新西兰":         "NZ",
		"阿曼":          "OM",
		"巴拿马":         "PA",
		"秘鲁":          "PE",
		"法属波利尼西亚":     "PF",
		"巴布亚新几内亚":     "PG",
		"菲律宾":         "PH",
		"巴基斯坦":        "PK",
		"波兰":          "PL",
		"圣皮埃尔和密克隆":    "PM",
		"皮特凯恩群岛":      "PN",
		"波多黎各":        "PR",
		"巴勒斯坦":        "PS",
		"葡萄牙":         "PT",
		"帕劳":          "PW",
		"巴拉圭":         "PY",
		"卡塔尔":         "QA",
		"留尼汪":         "RE",
		"罗马尼亚":        "RO",
		"塞尔维亚":        "RS",
		"俄罗斯":         "RU",
		"卢旺达":         "RW",
		"沙特阿拉伯":       "SA",
		"所罗门群岛":       "SB",
		"塞舌尔":         "SC",
		"苏丹":          "SD",
		"瑞典":          "SE",
		"新加坡":         "SG",
		"圣赫勒拿":        "SH",
		"斯洛文尼亚":       "SI",
		"斯瓦尔巴和扬马延":    "SJ",
		"斯洛伐克":        "SK",
		"塞拉利昂":        "SL",
		"圣马力诺":        "SM",
		"塞内加尔":        "SN",
		"索马里":         "SO",
		"苏里南":         "SR",
		"南苏丹":         "SS",
		"圣多美和普林西比":    "ST",
		"萨尔瓦多":        "SV",
		"荷属圣马丁":       "SX",
		"叙利亚":         "SY",
		"斯威士兰":        "SZ",
		"特克斯和凯科斯群岛":   "TC",
		"乍得":          "TD",
		"法属南方和南极洲领地":  "TF",
		"多哥":          "TG",
		"泰国":          "TH",
		"塔吉克斯坦":       "TJ",
		"托克劳":         "TK",
		"东帝汶":         "TL",
		"土库曼斯坦":       "TM",
		"突尼斯":         "TN",
		"汤加":          "TO",
		"土耳其":         "TR",
		"特立尼达和多巴哥":    "TT",
		"图瓦卢 ":        "TV",
		"坦桑尼亚":        "TZ",
		"乌克兰":         "UA",
		"乌干达":         "UG",
		"美国本土外小岛屿":    "UM",
		"美国":          "US",
		"乌拉圭":         "UY",
		"乌兹别克斯坦":      "UZ",
		"梵蒂冈":         "VA",
		"圣文森特和格林纳丁斯":  "VC",
		"委内瑞拉":        "VE",
		"英属维尔京群岛":     "VG",
		"美属维尔京群岛":     "VI",
		"越南":          "VN",
		"瓦努阿图":        "VU",
		"瓦利斯和富图纳":     "WF",
		"萨摩亚":         "WS",
		"也门":          "YE",
		"马约特":         "YT",
		"南非":          "ZA",
		"赞比亚":         "ZM",
		"津巴布韦":        "ZW",
	}
)

func Language2Code(language string) string {
	code, exists := languageCodeMap[language]
	if exists {
		return code
	}
	return "CN"
}

func Language2CodeV2(language string, defaultCode string) string {
	code, exists := languageCodeMap[language]
	if exists {
		return code
	}
	return defaultCode
}

func Country2Code(country string) string {
	code, exists := countryCodeMap[country]
	if exists {
		return code
	}

	return "CN"
}

func TimeIn(name string) time.Time {
	timezone, found := countryTz[name]
	if !found {
		timezone = "Local"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now()
	}

	return time.Now().In(loc)
}

func TimeFromString(country string, datetimeString string, toLocal bool) time.Time {
	timezone, found := countryTz[country]
	if !found {
		timezone = "Local"
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Now()
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", datetimeString, loc)
	if err != nil {
		return time.Now()
	}

	if toLocal {
		localLoc, err := time.LoadLocation("Local")
		if err != nil {
			return time.Now()
		}
		t = t.In(localLoc)
	}

	return t
}

func TimestampFromString(s string) int64 {
	var ret int64
	if t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", s); err == nil {
		ret = t.Unix()
	}
	return ret
}

func StringFromTimestamp(t int64, f string) string {
	if f == "" {
		f = "2006-01-02 15:04:05"
	}
	return time.Unix(t, 0).Format(f)
}
