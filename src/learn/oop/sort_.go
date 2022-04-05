package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012},
	{"Go", "Moby", "Moby", 1992},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011},
}

type byArtist []*Track

/**
排序接口 Len Less Swap
*/
func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
}

func printTracks(tracks []*Track) {
	// %v  相应值的默认格式
	// %+v 打印结构体时，会添加字段名
	const FORMAT = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, FORMAT, "Title", "Artist", "Album", "Year")
	fmt.Fprintf(tw, FORMAT, "-----", "------", "-----", "----")

	for _, t := range tracks {
		fmt.Fprintf(tw, FORMAT, t.Title, t.Artist, t.Album, t.Year)
	}
	tw.Flush() // calculate column widths and print table
}
