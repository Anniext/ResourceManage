package utils

func LimitParameter(s1 string , s2 string) (s string) {
    if s1 == "" && s2 != "" {
        return s2
    } else if s1 != "" && s2 == "" {
        return s1
    } else if s1 != "" && s2 != "" {
        return ""
    }
    return ""
}
