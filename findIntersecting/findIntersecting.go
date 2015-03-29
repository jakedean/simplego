/*
  This program will take in the top left corner and bottom
  right corner (x,y) of 2 rectanlges and print them on a plane
  to the screen.  We will find the intersection and paint it
  with the + sign.
 */

package findIntersecting

import (
  	"fmt"
  	"regexp"
  	"strings"
  	"strconv"
  	"bufio"
  	"os"
  	"math"
)

// This struct will hold two points, the top left point
// and the bottom right point.  This will be enough to draw
// the rectangle to the screen.
type Rect struct {
	TopLeft, BotRight Point
}

// This struct has the x and y coordinates for the point.
type Point struct {
	X, Y int64
}

// This struct will have the max and min x and y we willneed to
// put on the screen to draw the rectangles.
type XYPlane struct {
	MaxX, MinX, MaxY, MinY int
}

// A struct to hold slice of Rect structs
type Rects struct {
	R []*Rect
}

func (rects Rects) Show() {
	for i, rectPt := range rects.R {
		fmt.Printf("Here is rectangle number %d\n", (i + 1))
		rectPt.Draw()
	}
}


// A convienent way to just see what the rectangle looks like.
func (rec *Rect) Print() {
	fmt.Println("This is the rectangle!")
	fmt.Printf("Upper Left -> x: %d, y: %d\n", rec.TopLeft.X, rec.TopLeft.Y)
	fmt.Printf("Lower Right -> x: %d , y: %d\n", rec.BotRight.X, rec.BotRight.Y)
}

// A function to draw a rectangle on your console on an x,y plane.
func (rec *Rect) Draw() {
	// First we need to create our XYPlane
	plane := XYPlane{
		int(rec.BotRight.X) + 3,
		int(rec.TopLeft.X) -3,
		int(rec.TopLeft.Y) + 3,
		int(rec.BotRight.Y) -3,
	}
	for y := plane.MaxY; y >= plane.MinY; y-- {
		row := ""
		for x := plane.MinX; x <= plane.MaxX; x++ {
			if y <= int(rec.TopLeft.Y) && y >= int(rec.BotRight.Y) && x >= int(rec.TopLeft.X) && x <= int(rec.BotRight.X) {
			   	row += "@"
			} else if y == 0 || x == 0 {
			    row += "="
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
	fmt.Println()
}

// The main func for this package to start us off.
func Start() {
	giveInstructions()
}

// Here we will give the instructions to the user about the game.
func giveInstructions() {
	fmt.Println("Welcome to the findIntersection program!")
	fmt.Println("Please give the coordinates for the first rectangle.")
	fmt.Println("You will give them like x,y x,y where the first pair is the top left, second is bottom right.")
	// Get the rects ready.
	theRects := Rects{}
	for i:=0;i<2;i++ {
		rectPt := callForCoords()
		theRects.R = append(theRects.R, rectPt)
	}
	theRects.Show()
	theRects.findIntersection()
}

func callForCoords() *Rect {
	fmt.Println("Please give the coordinates for the first rectangle now...")
	reader := bufio.NewReader(os.Stdin)
	// Read in the bytes from stdIn and then strip out the new line.
	coords, _ := reader.ReadBytes('\n')
	rectPt := parseCoords(string(coords[0:len(coords)-1]))
	return rectPt
}

func parseCoords(coordStr string) *Rect {
	// We must match this pattern 3,4 5,6
	match, _ := regexp.MatchString("-*[0-9]+,-*[0-9]+ -*[0-9]+,-*[0-9]+", coordStr)
	if match {
		// We have some valid coords so we will split the string
		// into the parts.
		leftAndRight := strings.Split(coordStr, " ")
		left := strings.Split(leftAndRight[0], ",")
		right := strings.Split(leftAndRight[1], ",")
		lx, _ := strconv.ParseInt(left[0], 10, 8)
		ly, _ := strconv.ParseInt(left[1], 10, 8)
		rx, _ := strconv.ParseInt(right[0], 10, 8)
		ry, _ := strconv.ParseInt(right[1], 10, 8)
		rect := &Rect{Point{lx, ly}, Point{rx, ry}}
		return rect
	} else {
		fmt.Println("Make sure you have the right format, lets try again.")
		return callForCoords()
	}
}

// This will get the intersection of the rects.
func (rects Rects) findIntersection() {
	lx := math.Max(float64(rects.R[0].TopLeft.X), float64(rects.R[1].TopLeft.X))
	ly := math.Min(float64(rects.R[0].TopLeft.Y), float64(rects.R[1].TopLeft.Y))
	rx := math.Min(float64(rects.R[0].BotRight.X), float64(rects.R[1].BotRight.X))
	ry := math.Max(float64(rects.R[0].BotRight.Y), float64(rects.R[1].BotRight.Y))
	intersec := &Rect{Point{int64(lx), int64(ly)}, Point{int64(rx), int64(ry)}}
	fmt.Println("This is the intersection rectangle.")
	intersec.Draw()
}