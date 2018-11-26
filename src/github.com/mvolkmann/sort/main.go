package main

import (
	"container/list"
	"fmt"
	"sort"
)

// ListAt returns a pointer to the List Element at a given index.
// This is not an efficient operation for a linked list.
func ListAt(l *list.List, index int) *list.Element {
	len := l.Len()
	if index >= len {
		return nil
	}

	element := l.Front()
	for i := 1; i < index; i++ {
		element = element.Next()
	}
	return element
}

// ListPrint prints the elements of a linked list.
func ListPrint(l *list.List) {
	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

// Team describes a sports team.
type Team struct {
	Name   string
	Wins   int
	Losses int
}

// ByName is used to sort a linked list of Team object by team name.
type ByName struct {
	teams *list.List
}

// Len returns the length of the linked list of Teams.
func (b ByName) Len() int {
	return b.teams.Len()
}

// Less determines if the Team at one index belongs
// before the Team at another index in the sort order.
func (b ByName) Less(i, j int) bool {
	teams := b.teams
	team1 := ListAt(teams, i).Value.(Team)
	team2 := ListAt(teams, j).Value.(Team)
	return team1.Name < team2.Name
}

// Swap swaps the Teams at the given indexes.
func (b ByName) Swap(i, j int) {
	if i > j {
		i, j = j, i // swap indexes so i < j
	}

	teams := b.teams
	el1 := ListAt(teams, i)
	el2 := ListAt(teams, j)
	el2Next := el2.Next()

	teams.MoveAfter(el2, el1)

	if el2Next == nil {
		teams.MoveToBack(el1)
	} else {
		teams.MoveBefore(el1, el2Next)
	}
}

func main() {
	// Create an empty, doubly-linked list.
	teams := list.New()

	// Add Teams to the list.
	firstPlace := teams.PushFront(Team{"Chiefs", 9, 2})
	lastPlace := teams.PushBack(Team{"Raiders", 2, 8})
	teams.InsertBefore(Team{"Broncos", 4, 6}, lastPlace)
	teams.InsertAfter(Team{"Chargers", 7, 3}, firstPlace)

	sort.Sort(ByName{teams})
	ListPrint(teams) // Broncos, Chargers, Chiefs, Raiders
}
