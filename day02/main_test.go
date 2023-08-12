package main

import (
	"testing"
)

// TestShapeToString_ShapeIsPaper calls main.ShapeToString with 'A'
// which must return 'rock'.
func TestShapeToString_WithShapeA_ShouldReturnRock(t *testing.T) {
	shape := "A"
	want := "rock"

	shapeString, err := ShapeToString(shape)

	if want != shapeString || err != nil {
		t.Fatalf(`ShapeToString("A") = %q, %v, want equal for %#q, nil`, shapeString, err, want)
	}
}

// TestShapeToString_ShapeIsPaper calls main.ShapeToString with 'B'
// which must return 'paper'.
func TestShapeToString_WithShapeB_ShouldReturnPaper(t *testing.T) {
	shape := "B"
	want := "paper"

	shapeString, err := ShapeToString(shape)

	if want != shapeString || err != nil {
		t.Fatalf(`ShapeToString("B") = %q, %v, want equal for %#q, nil`, shapeString, err, want)
	}
}

// TestShapeToString_ShapeIsPaper calls main.ShapeToString with 'C'
// which must return 'scissors'.
func TestShapeToString_WithShapeC_ShouldReturnScissors(t *testing.T) {
	shape := "C"
	want := "scissors"

	shapeString, err := ShapeToString(shape)

	if want != shapeString || err != nil {
		t.Fatalf(`ShapeToString("C") = %q, %v, want equal for %#q, nil`, shapeString, err, want)
	}
}

func TestShapeToString_WithUnknownShape_ShouldReturnError(t *testing.T) {
	shape := "T"
	want := ""

	shapeString, err := ShapeToString(shape)

	if want != shapeString || err == nil {
		t.Fatalf(`ShapeToString("C") = %q, %v, want equal for %#q, nil`, shapeString, err, want)
	}
}

func TestShapeToCoor_WithShapeRock_ShouldReturn0(t *testing.T) {
	shape := "rock"
	want := 0

	shapeCoor, err := ShapeToCoor(shape)

	if want != shapeCoor || err != nil {
		t.Fatalf(`ShapeToCoor("rock") = %q, %v, want equal for %#q, nil`, shapeCoor, err, want)
	}
}

func TestShapeToCoor_WithShapePaper_ShouldReturn1(t *testing.T) {
	shape := "paper"
	want := 1

	shapeCoor, err := ShapeToCoor(shape)

	if want != shapeCoor || err != nil {
		t.Fatalf(`ShapeToCoor("paper") = %q, %v, want equal for %#q, nil`, shapeCoor, err, want)
	}
}

func TestShapeToCoor_WithShapeScissors_ShouldReturn2(t *testing.T) {
	shape := "scissors"
	want := 2

	shapeCoor, err := ShapeToCoor(shape)

	if want != shapeCoor || err != nil {
		t.Fatalf(`ShapeToCoor("scissors") = %q, %v, want equal for %#q, nil`, shapeCoor, err, want)
	}
}

func TestShapeToCoor_WithUnknowShape_ShouldReturnError(t *testing.T) {
	shape := "fire"
	want := -1

	shapeCoor, err := ShapeToCoor(shape)

	if want != shapeCoor || err == nil {
		t.Fatalf(`ShapeToCoor("fire") = %q, %v, want equal for %#q, nil`, shapeCoor, err, want)
	}
}

func TestGetPlayScore_WithAY_ShouldReturn8(t *testing.T) {
	play := []string{"A", "Y"}
	want := 8

	playScore := GetPlayScore(play)

	if want != playScore {
		t.Fatalf(`GetPlayScore(["A", "Y"]) = %v, want equal for %#v`, playScore, want)
	}
}

func TestGetPlayScore_WithBX_ShouldReturn1(t *testing.T) {
	play := []string{"B", "X"}
	want := 1

	playScore := GetPlayScore(play)

	if want != playScore {
		t.Fatalf(`GetPlayScore(["B", "X"]) = %v, want equal for %#v`, playScore, want)
	}
}

func TestGetPlayScore_WithCZ_ShouldReturn6(t *testing.T) {
	play := []string{"C", "Z"}
	want := 6

	playScore := GetPlayScore(play)

	if want != playScore {
		t.Fatalf(`GetPlayScore(["C", "Z"]) = %v, want equal for %#v`, playScore, want)
	}
}

func TestPlayToScore(t *testing.T) {
	play := []string{"A", "Y"}
	want := 4

	res := PlayToScore(play)

	if want != res {
		t.Fatalf(`PlayToScore(["A", "Y"]) = %v, want equal for %#v`, res, want)
	}
}

func TestPlayToScore_WithBX_ShouldReturn1(t *testing.T) {
	play := []string{"B", "X"}
	want := 1

	res := PlayToScore(play)

	if want != res {
		t.Fatalf(`PlayToScore(["B", "X"]) = %v, want equal for %#v`, res, want)
	}
}
