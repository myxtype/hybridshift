package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// 判断字符串是否在数组中
func InStringSlice(v string, ss []string) bool {
	for _, n := range ss {
		if n == v {
			return true
		}
	}
	return false
}

// 获取文件名中的文件类型
func ResolveFileType(filename string) string {
	fileNames := strings.Split(filename, ".")
	if len(fileNames) < 2 {
		return ""
	}

	return strings.ToLower(fileNames[len(fileNames)-1])
}

// 判断是否是版本号字符串
func IsVersion(version string) bool {
	regex := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return regex.MatchString(version)
}

func NonEmptyString(s1, s2 string) string {
	if s2 != "" {
		return s2
	}
	return s1
}

func SnakeCase(s string) string {
	in := []rune(s)
	isLower := func(idx int) bool {
		return idx >= 0 && idx < len(in) && unicode.IsLower(in[idx])
	}

	out := make([]rune, 0, len(in)+len(in)/2)
	for i, r := range in {
		if unicode.IsUpper(r) {
			r = unicode.ToLower(r)
			if i > 0 && in[i-1] != '_' && (isLower(i-1) || isLower(i+1)) {
				out = append(out, '_')
			}
		}
		out = append(out, r)
	}

	return string(out)
}
