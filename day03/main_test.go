package main

import "testing"

func TestItemToPriority_WithA_ShouldReturn27(t *testing.T) {
	item := "A"
	want := 27

	itemPriority, err := ItemToPriority(item)

	if itemPriority != want || err != nil {
		t.Fatalf(`ItemToPriority("%s") = %d, %v, want equal for %d, nil`, item, itemPriority, err, want)
	}
}

func TestItemToPriority_Withp_ShouldReturn16(t *testing.T) {
	item := "p"
	want := 16

	itemPriority, err := ItemToPriority(item)

	if itemPriority != want || err != nil {
		t.Fatalf(`ItemToPriority("%s") = %d, %v, want equal for %d, nil`, item, itemPriority, err, want)
	}
}

func TestItemToPriority_WithL_ShouldReturn38(t *testing.T) {
	item := "L"
	want := 38

	itemPriority, err := ItemToPriority(item)

	if itemPriority != want || err != nil {
		t.Fatalf(`ItemToPriority("%s") = %d, %v, want equal for %d, nil`, item, itemPriority, err, want)
	}
}

func TestItemToPriority_WithP_ShouldReturn42(t *testing.T) {
	item := "P"
	want := 42

	itemPriority, err := ItemToPriority(item)

	if itemPriority != want || err != nil {
		t.Fatalf(`ItemToPriority("%s") = %d, %v, want equal for %d, nil`, item, itemPriority, err, want)
	}
}

func TestItemsStringToArray(t *testing.T) {
	items := "vJrwpWtwJgWrhcsFMMfFFhFp"
	want := []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r", "h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"}

	itemsArray, err := ItemsStringToArray(items)

	if len(itemsArray) != len(want) || err != nil {
		t.Fatalf(`ItemsStringToArray("%s") = %#q, %v, want equal for %#q, nil`, items, itemsArray, err, want)
	}
}

func TestRucksackToCompartments(t *testing.T) {
	rucksack := []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r", "h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"}
	want1 := []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r"}
	want2 := []string{"h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"}

	firstCompartment, secondCompartment, err := RucksackToCompartments(rucksack)

	if !compareStringArrays(firstCompartment, want1) || !compareStringArrays(secondCompartment, want2) || err != nil {
		t.Fatalf(`RucksackToCompartments("%#q") = %#q, %#q, %v, want equal for %#q, %#q, nil`, rucksack, firstCompartment, secondCompartment, err, want1, want2)
	}
}

func TestRucksackToCompartments_WithOddNumberOfItems(t *testing.T) {
	rucksack := []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r", "h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p", "j"}

	_, _, err := RucksackToCompartments(rucksack)

	if err == nil {
		t.Fatalf(`RucksackToCompartments("%#q") with odd numbers of items should return error %v`, rucksack, err)
	}
}

func TestGetItemsInBothCompartments(t *testing.T) {
	rucksack := []string{"v", "J", "r", "w", "p", "W", "t", "w", "J", "g", "W", "r", "h", "c", "s", "F", "M", "M", "f", "F", "F", "h", "F", "p"}
	want := []string{"p"}

	firstCompartment, secondCompartment, _ := RucksackToCompartments(rucksack)

	items := GetItemsInBothCompartment(firstCompartment, secondCompartment)

	if !compareStringArrays(want, items) {
		t.Fatalf(`GetItemsInBothCompartment("%#q") = %#q, want %#q`, rucksack, items, want)
	}

}

func compareStringArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
