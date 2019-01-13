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

func randomfiveencounters(i int) []int {
	s := make([]int, 5)
	s[0] = rand.Intn(i)
	s[1] = rand.Intn(i)
	s[2] = rand.Intn(i)
	s[3] = rand.Intn(i)
	s[4] = rand.Intn(i)
	return s
}

func getXPandName(is []int, m Monsters) []string {

	one, _ := strconv.Atoi(m.Monsters[is[0]].XP)
	two, _ := strconv.Atoi(m.Monsters[is[1]].XP)
	three, _ := strconv.Atoi(m.Monsters[is[2]].XP)
	four, _ := strconv.Atoi(m.Monsters[is[3]].XP)
	five, _ := strconv.Atoi(m.Monsters[is[4]].XP)
	total := one + two + three + four + five
	namesAndXP := make([]string, 11)
	for i := range is {
		namesAndXP[i] = m.Monsters[is[i]].XP
		namesAndXP[i+5] = m.Monsters[is[i]].Name

	}
	namesAndXP[10] = strconv.Itoa(total)
	return namesAndXP
}
func findWinner(s []string) ([]string, bool) {
	argOne, _ := strconv.Atoi(os.Args[1])

	total, _ := strconv.Atoi(s[10])
	adjTotal := total * 2
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

	fmt.Println(len(monsters.Monsters))

	for {
		s := randomfiveencounters(len(monsters.Monsters))
		final := getXPandName(s, monsters)
		msg, tf := findWinner(final)
		if tf {
			fmt.Println("success: ", msg)
		}
	}

	fmt.Println(monsters.Monsters)

}
