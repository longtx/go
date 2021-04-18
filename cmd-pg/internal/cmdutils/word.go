package cmdutils

import (
	"strings"
	"unicode"
)

/**
creator:longtuxing
description: 单词全部转为小写/大写  把下划线替换成驼峰式
 */

func ToUpper (s string) string  {
	return strings.ToUpper(s)
}

func ToLower(s string) string  {
	return strings.ToLower(s)
}

//把下划线替换成大写驼峰式
func UnderscoreToUpperCamelCase(s string) string  {
	s = strings.Replace(s, "_", " ",1)
	s = strings.Title(s)
	return strings.Replace(s, " " ,"",-1)
}

//把下划线替换成小写驼峰式
func UnderscoreToLowerCamelCase(s string) string  {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

//把驼峰式单词转下划线单词
func CamelCaseToUnderscore(s string) string  {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}