package utils

import "strings"

type TagHelper struct {
	DefaultTags map[string]string
	IgnoreTags  []string
}

func (t *TagHelper) GetCombinedTags(rt []string) []string {
	resourceTagMap := tagStringToMap(rt)
	combinedTagMap := mergeMaps(t.DefaultTags, resourceTagMap)
	return tagMapToString(combinedTagMap)
}

func (t *TagHelper) GetTagsToCompare(rt []string) []string {
	tagsToClean := tagStringToMap(rt)
	for _, tag := range t.IgnoreTags {
		delete(tagsToClean, tag)
	}
	return tagMapToString(tagsToClean)
}

func mergeMaps(maps ...map[string]string) map[string]string {
	o := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			o[k] = v
		}
	}
	return o
}

func tagStringToMap(tagSlice []string) map[string]string {
	m := make(map[string]string)

	for _, s := range tagSlice {
		splitTag := strings.Split(s, ":")
		if splitTag[0] == "" {
			// The tag is like ":value", so we use it as the key, but we keep the : on the front
			m[strings.Join(splitTag, `:`)] = ""
		} else {
			// A more normal "tag:value"
			m[splitTag[0]] = strings.Join(splitTag[1:], `:`)
		}
	}

	return m
}

func tagMapToString(tagMap map[string]string) []string {
	s := make([]string, len(tagMap))

	i := 0
	for k, v := range tagMap {
		if v == "" && strings.HasPrefix(v, ":") {
			// The key is like ":value", and has no value, so we write it as the key only
			s[i] = k
		} else {
			s[i] = k + ":" + v
		}
		i++
	}
	return s
}
