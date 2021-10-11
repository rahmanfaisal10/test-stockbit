package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Print("Silahkan masukan jumlah kata yang akan di Anagram kan: ")
	var inputAmount int
	fmt.Scanln(&inputAmount)

	kata := make([]string, inputAmount)
	fmt.Println("Silahkan masukan kata yang akan di masukan lalu enter, lakukan sebanyak jumlah kata yang anda masukan sebelumnya: ")
	for i := 0; i < inputAmount; i++ {
		fmt.Scanln(&kata[i])
	}

	fmt.Println("hasilnya adalah: ")
	Anagram(kata)
}

func Anagram(req []string) {
	list := make(map[string][]string)
	sortList := make(map[int]string)

	for _, word := range req {
		key := sortStr(word)
		list[key] = append(list[key], word)
		sortList[len(list[key])] = key
	}

	fmt.Println(sortList)
	for _, words := range list {
		wordJson, _ := json.Marshal(words)
		fmt.Println(string(wordJson))
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
