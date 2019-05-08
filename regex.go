package regex
import (
    "regexp"
)
var regexPattern  = `(\w+\d+)+C` 

// MatchRegex matches input string against a pattern
func MatchRegex (input string) bool {
    match, _ := regexp.MatchString(regexPattern, input)
    return match
}

// MatchCompiledRegex using Compile
func MatchCompiledRegex (input string) bool {
    r, _ := regexp.Compile(regexPattern)
    return r.MatchString(input)
}

