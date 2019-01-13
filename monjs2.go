package main

import (
	"encoding/json"
	"fmt" //"reflect"
	"io/ioutil"
	"math/rand" //"regexp"
	"os"
	"strconv"
)

type Monsters struct {
	Monsters []Monster `json:"monsters"`
	//Monster  Monster    `json:"monster"`
}

type Monster struct {
	Name        string `json:"name"`
	Id          int    `json:"id"`
	Age         string `json:"HITPOINTS"`
	AC          int    `json:"AC"`
	CR          int    `json:"CR"`
	XP          string `json:"XP"`
	Description string `json:"Descriptions"`
	//squad
}

//type squad interface {

func randomencounters(i int) []int {
	argTwo, _ := strconv.Atoi(os.Args[2])
	s := make([]int, argTwo+1)
	for n := range s {
		s[n] = rand.Intn(i)
	}
	return s
}

func getXPandName(is []int, m Monsters) []string {
	total := 0
	//argTwo, _ := strconv.Atoi(os.Args[2])
	argTwo, _ := strconv.Atoi(os.Args[2])
	maxTwo := argTwo*2 + 1
	//total := one + two + three + four + five
	var intXp int
	namesAndXP := make([]string, maxTwo+1)
	for i := range is {
		namesAndXP[i] = m.Monsters[is[i]].XP
		namesAndXP[i+argTwo] = m.Monsters[is[i]].Name
		intXp, _ = strconv.Atoi(m.Monsters[is[i]].XP)

		total = total + intXp
	}
	namesAndXP[maxTwo] = strconv.Itoa(total)
	return namesAndXP
}
func findWinner(s []string) ([]string, bool) {
	var modify [16]float64
	modify[0] = 1
	modify[1] = 1
	modify[2] = 1.5
	modify[3] = 2
	modify[4] = 2
	modify[5] = 2
	modify[6] = 2
	modify[7] = 2.5
	modify[8] = 2.5
	modify[9] = 2.5
	modify[10] = 2.5
	modify[11] = 3
	modify[12] = 3
	modify[13] = 3
	modify[14] = 3
	modify[15] = 4
	argTwo, _ := strconv.Atoi(os.Args[2])
	argOne, _ := strconv.ParseFloat(os.Args[1], 64)

	total, _ := strconv.ParseFloat(s[argTwo*2+1], 64)
	adjTotal := total * modify[argTwo]
	if adjTotal < argOne {
		return s, true
	}
	return s, false
}
func (m Monster) N() string {
	return m.Name
}

func main() {

	jsonFile, err := os.Open("/Users/dustincarver/torrents/dun/DnD.5e/Rulebooks/Core/dummys.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened monsters.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var monsters Monsters

	json.Unmarshal(byteValue, &monsters)

	for {
		s := randomencounters(len(monsters.Monsters))
		final := getXPandName(s, monsters)
		msg, tf := findWinner(final)
		if tf {
			fmt.Println("success: ", msg)
		}
	}

	fmt.Println(monsters.Monsters)

}
