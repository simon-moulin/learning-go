package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}

	return found, res, occ
}
func main() {
	new := "Python"
	old := "Go"
	occ, lines, err := FindReplaceFile("wikigo.txt", old, new)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("== Summary ==")
	fmt.Printf("Number of occurences of %s: %d\n", old, occ)
	fmt.Println("Number of lines :", len(lines))
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")
	fmt.Println("== End of Summary ==")
}

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	var total string

	fileToRead, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer fileToRead.Close()

	old = old + " "
	new = new + " "
	lineIdx := 1
	scanner := bufio.NewScanner(fileToRead)
	for scanner.Scan() {
		found, curRes, curOcc := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += curOcc
			lines = append(lines, lineIdx)
		}
		total += curRes + "\n"
		lineIdx++
	}
	fmt.Println(total)
	return occ, lines, nil
}
