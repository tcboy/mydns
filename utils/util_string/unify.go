package util_string

import (
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

var tagInvalidReplace = ""
var maxTagLength = 60
var tagReplacement = map[string]string{
	"\u0130": "i\u0307",
	"\u0526": "\u0527",
	"\ua660": "\ua661",
}
var regexAllInvalidTagChar = regexp.MustCompile(
	"[^0-9\uff10-\uff19_a-zA-Z\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u00ff\u0100-\u024f\u0253\u0254\u0256\u0257\u0259\u025b\u0263\u0268\u026f\u0272\u0289\u028b\u02bb\u0300-\u036f\u1e00-\u1eff\u0400-\u04ff\u0500-\u0527\u2de0-\u2dff\ua640-\ua69f\u0591-\u05bf\u05c1-\u05c2\u05c4-\u05c5\u05c7\u05d0-\u05ea\u05f0-\u05f4\ufb1d-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufb4f\u0610-\u061a\u0620-\u065f\u066e-\u06d3\u06d5-\u06dc\u06de-\u06e8\u06ea-\u06ef\u06fa-\u06fc\u06ff\u0750-\u077f\u08a0\u08a2-\u08ac\u08e4-\u08fe\ufb50-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe70-\ufe74\ufe76-\ufefc\u200c\u0e01-\u0e3a\u0e40-\u0e4e\u1100-\u11ff\u3130-\u3185\uA960-\uA97F\uAC00-\uD7AF\uD7B0-\uD7FF\u3003\u3005\u303b\uff21-\uff3a\uff41-\uff5a\uff66-\uff9f\uffa1-\uffdc\u4E00-\u9FFF\uF900-\uFAFF\u3041-\u3096\u3099-\u30fe\u30fc\u30ff]+")

func UnifyHashtag(tag string) string {
	if len(tag) == 0 {
		return ""
	}

	tag = strings.TrimSpace(tag)
	tag = regexAllInvalidTagChar.ReplaceAllString(tag, tagInvalidReplace)

	rs := []rune(tag)
	if len(rs) > maxTagLength {
		rs = rs[:maxTagLength]
	}

	tag = string(rs)
	tag = norm.NFKC.String(tag)

	for older, newer := range tagReplacement {
		tag = strings.Replace(tag, older, newer, -1)
	}

	return strings.ToLower(tag)
}
