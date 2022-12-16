package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"

	"github.com/jaaanko/advent-of-code-2022/imath"
	"github.com/jaaanko/advent-of-code-2022/stack"
)

type point struct {
	x, y int
}

type searchArea struct {
	top, bottom, left, right point
	radius                   int
}

type interval struct {
	start, end int
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	closest := parseInput(file)
	fmt.Println("Part 1:", solvePart1(closest))
	fmt.Println("Part 2:", solvePart2(closest))
}

func parseInput(r io.Reader) map[point]point {
	scanner := bufio.NewScanner(r)
	closest := map[point]point{}

	for scanner.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		closest[point{sensorX, sensorY}] = point{beaconX, beaconY}
	}
	return closest
}

func solvePart1(closest map[point]point) int {
	areas := []searchArea{}
	beacons := map[point]bool{}

	for sensor, beacon := range closest {
		beacons[beacon] = true
		dist := manhattanDist(sensor, beacon)
		areas = append(areas, searchArea{
			top:    point{sensor.x, sensor.y - dist},
			bottom: point{sensor.x, sensor.y + dist},
			left:   point{sensor.x - dist, sensor.y},
			right:  point{sensor.x + dist, sensor.y},
			radius: dist,
		})
	}

	y := 10
	intervals := scannedIntervals(areas, y)
	sort.Slice(intervals, func(i, j int) bool {
		return compare(intervals[i], intervals[j]) < 0
	})

	merged := stack.New[interval]()

	for _, curr := range intervals {
		last, _ := merged.Peek()
		if len(merged) > 0 && last.end >= curr.start {
			merged[len(merged)-1].end = imath.MaxInt(last.end, curr.end)
		} else {
			merged.Push(curr)
		}
	}

	total := 0
	for _, in := range merged {
		total += in.end - in.start + 1
	}

	for beacon, _ := range beacons {
		if beacon.y == y {
			for _, in := range merged {
				if in.start <= beacon.x && beacon.x <= in.end {
					total -= 1
				}
			}
		}
	}
	return total
}

func manhattanDist(p1 point, p2 point) int {
	return imath.AbsInt(p1.x-p2.x) + imath.AbsInt(p1.y-p2.y)
}

func compare(i1 interval, i2 interval) int {
	if i1.start != i2.start {
		return i1.start - i2.start
	}
	return i1.end - i2.end
}

func scannedIntervals(areas []searchArea, y int) []interval {
	intervals := []interval{}

	for _, area := range areas {
		if area.top.y <= y && y <= area.bottom.y {
			minY := imath.MinInt(y, area.left.y)
			maxY := imath.MaxInt(y, area.left.y)
			remaining := area.radius*2 + 1 - (maxY-minY)*2
			start := maxY - minY + area.left.x
			intervals = append(intervals, interval{start, start + remaining - 1})
		}
	}
	return intervals
}

func solvePart2(closest map[point]point) int {
	areas := []searchArea{}
	beacons := map[point]bool{}

	for sensor, beacon := range closest {
		beacons[beacon] = true
		dist := manhattanDist(sensor, beacon)
		areas = append(areas, searchArea{
			top:    point{sensor.x, sensor.y - dist},
			bottom: point{sensor.x, sensor.y + dist},
			left:   point{sensor.x - dist, sensor.y},
			right:  point{sensor.x + dist, sensor.y},
			radius: dist,
		})
	}

	for y := 0; y <= 20; y++ {
		intervals := scannedIntervals(areas, y)
		sort.Slice(intervals, func(i, j int) bool {
			return compare(intervals[i], intervals[j]) < 0
		})

		merged := stack.New[interval]()

		for _, curr := range intervals {
			last, _ := merged.Peek()
			if len(merged) > 0 && (last.end >= curr.start || curr.start-last.end == 1) {
				if last.end >= curr.start {
					merged[len(merged)-1].end = imath.MaxInt(last.end, curr.end)
				} else {
					merged[len(merged)-1].end = curr.end
				}
			} else {
				merged.Push(curr)
			}
		}

		if len(merged) == 2 {
			return (merged[0].end+1)*4000000 + y
		}
	}
	return -1
}
