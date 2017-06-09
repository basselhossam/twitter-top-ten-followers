// gets top 10 most followed on twitter dataset.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sort"
	"os"
	"strings"
	"bufio"
	"log"
)

// Implementing sort interface.
type SortUser struct {
	Followers map[int]int
	Keys      []int
}

// Gets length of map.
func (su *SortUser) Len() int {
	// TODO: Implement Len function.
	return len(su.Followers)
}

// Condition for sorting to compare between values of keys.
func (su *SortUser) Less(i, j int) bool {
	// TODO: Implement Less function.
	return su.Followers[su.Keys[i]] > su.Followers[su.Keys[j]]
}

// Swaps two keys in keys array.
func (su *SortUser) Swap(i, j int) {
	// TODO: Implement Swap function.
	su.Keys[i],su.Keys[j] = su.Keys[j],su.Keys[i]
}

// Sorts Keys based on number of followers in descending order.
func sortKeys(m map[int]int) []int {
	// TODO: Implement sortKeys function.
	var d SortUser
	d.Followers = m
	for k := range m {
	    d.Keys = append(d.Keys, k)
	}
	sort.Sort(&d)
	return d.Keys
}

// Calculates top 10 most followed for input file
// and returns array of user id (int) for top 10.
func topTen(dataInput string) []int {
	// TODO: Implement topTen function.
	file, err := os.Open(dataInput)
	if err != nil {
	    log.Fatal(err)
	}

	defer file.Close() // What does this line do?
	scanner := bufio.NewScanner(file)
	data := make(map[int]int)
	for scanner.Scan() {
        line := scanner.Text()
        s := strings.Fields(line)
        index,err := strconv.Atoi(s[1])
        if(err != nil){
        	return nil
        }
        data[index] = data[index] + 1
	}
	return sortKeys(data)[0:10]
}

// Connects to remote service through internet to convert user id
// to username.
func getUsername(userId string) string {
	response, err := http.PostForm("https://tweeterid.com/ajax.php",
		url.Values{"input": {userId}})
	if err != nil {
		fmt.Println("Error getting username("+userId+"): ", err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in response("+userId+"): ", err)
		return ""
	}
	if string(body) == "error" {
		return ""
	}
	return string(body)
}

func main() {
	fmt.Println("Calculating top 10 most followed...")
	topId := topTen("input.txt")
	fmt.Println("topTen length: ", len(topId))

	fmt.Println("Getting and printing screen name for top 10...")
	d1 := ""
	var username string
	for i := 0; i < 10 && i < len(topId); i++ {
		username = getUsername(strconv.Itoa(topId[i]))
		fmt.Printf("%-15d%s\n", topId[i], username)
		d1 = d1 + strconv.Itoa(topId[i]) + " " + username + "\n"
	}
	out := []byte(d1)
	err := ioutil.WriteFile("result.txt", out,0644)
	if err != nil {
    	panic(err)
	}
}
