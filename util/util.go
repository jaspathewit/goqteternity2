package util

import (
	"sort"
)

// Type encapsulates a map and the keys
type sortedMap struct {
	Map  map[string]int
	Keys []string
}

// The length of the sorted map
func (sortedMap *sortedMap) Len() int {
	return len(sortedMap.Map)
}

// test if one element is less than anothe element
func (sortedMap *sortedMap) Less(i, j int) bool {
	return sortedMap.Map[sortedMap.Keys[i]] < sortedMap.Map[sortedMap.Keys[j]]
}

// swap elements
func (sortedMap *sortedMap) Swap(i, j int) {
	sortedMap.Keys[i], sortedMap.Keys[j] = sortedMap.Keys[j], sortedMap.Keys[i]
}

// Mothod takes a map[string]int and returns array of map keys sorted by the value
// in the given map
func SortedKeys(mapToSort map[string]int) []string {
	sortedMap := new(sortedMap)
	sortedMap.Map = mapToSort
	sortedMap.Keys = make([]string, len(mapToSort))
	i := 0
	for key, _ := range mapToSort {
		sortedMap.Keys[i] = key
		i++
	}
	sort.Sort(sortedMap)
	return sortedMap.Keys
}

// Method returns true if the given slice of strings contains the given element
func Contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
