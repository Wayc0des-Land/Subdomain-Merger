package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	firstFile := "first.txt"
	secondFile := "second.txt"
	mergedFile := "merged.txt"

	// Read subdomains from first file
	firstSubdomains, err := readSubdomainsFromFile(firstFile)
	if err != nil {
		fmt.Printf("Error reading from first file: %s\n", err.Error())
		return
	}

	// Read subdomains from second file
	secondSubdomains, err := readSubdomainsFromFile(secondFile)
	if err != nil {
		fmt.Printf("Error reading from second file: %s\n", err.Error())
		return
	}

	// Remove "http://" and "https://" prefixes from subdomains
	firstSubdomains = removePrefixFromSubdomains(firstSubdomains, "http://")
	firstSubdomains = removePrefixFromSubdomains(firstSubdomains, "https://")
	secondSubdomains = removePrefixFromSubdomains(secondSubdomains, "http://")
	secondSubdomains = removePrefixFromSubdomains(secondSubdomains, "https://")

	// Merge subdomains
	mergedSubdomains := mergeSubdomains(firstSubdomains, secondSubdomains)

	// Remove duplicates from merged subdomains
	mergedSubdomains = removeDuplicates(mergedSubdomains)

	// Sort merged subdomains
	sort.Strings(mergedSubdomains)

	// Write merged subdomains to new file
	err = writeSubdomainsToFile(mergedFile, mergedSubdomains)
	if err != nil {
		fmt.Printf("Error writing to merged file: %s\n", err.Error())
		return
	}

	fmt.Println("Subdomains merged and sorted successfully.")
}

// Read subdomains from a file and return them as a slice
func readSubdomainsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	subdomains := strings.Split(string(content), "\n")
	return subdomains, nil
}

// Remove a prefix from each subdomain in the slice
func removePrefixFromSubdomains(subdomains []string, prefix string) []string {
	for i, subdomain := range subdomains {
		subdomains[i] = strings.TrimPrefix(subdomain, prefix)
	}
	return subdomains
}

// Merge two slices of subdomains into a single slice
func mergeSubdomains(subdomains1, subdomains2 []string) []string {
	mergedSubdomains := append(subdomains1, subdomains2...)
	return mergedSubdomains
}

// Remove duplicates from the slice
func removeDuplicates(subdomains []string) []string {
	uniqueSubdomains := make(map[string]bool)
	result := []string{}

	for _, subdomain := range subdomains {
		if !uniqueSubdomains[subdomain] {
			uniqueSubdomains[subdomain] = true
			result = append(result, subdomain)
		}
	}

	return result
}

// Write subdomains to a file
func writeSubdomainsToFile(filename string, subdomains []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, subdomain := range subdomains {
		_, err := file.WriteString(subdomain + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
