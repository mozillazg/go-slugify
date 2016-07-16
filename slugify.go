package slugify

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mozillazg/go-unidecode"
)

// Separator separator between words
var Separator = "-"
var separatorForRe = regexp.QuoteMeta(Separator)

// ReInValidChar match invalid slug character
var ReInValidChar = regexp.MustCompile(fmt.Sprintf("[^%sa-zA-Z0-9]", separatorForRe))
var reDupSeparatorChar = regexp.MustCompile(fmt.Sprintf("%s{2,}", separatorForRe))

// Version return version
func Version() string {
	return "0.1.0"
}

// Slugify implements make a pretty slug from the given text.
// e.g. Slugify("kožušček hello world") => "kozuscek-hello-world"
func Slugify(s string) string {
	return slugify(s)
}

func slugify(s string) string {
	s = unidecode.Unidecode(s)
	s = replaceInValidCharacter(s, Separator)
	s = removeDumpSeparator(s)
	s = strings.Trim(s, Separator)
	s = strings.ToLower(s)
	return s
}

func replaceInValidCharacter(s, repl string) string {
	s = ReInValidChar.ReplaceAllString(s, repl)
	return s
}

func removeDumpSeparator(s string) string {
	s = reDupSeparatorChar.ReplaceAllString(s, Separator)
	return s
}
