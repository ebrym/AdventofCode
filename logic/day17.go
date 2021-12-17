package logic

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile17 = flag.String("inputFile17", "assets/day17.txt", "Relative file path to use as input.")

func Day17Task1() int {
	flag.Parse()

	// input, err := general.ReadLines("assets/day17.txt")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }

	bytes, err := ioutil.ReadFile(*inputFile17)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	contents := string(bytes)

	split := strings.Split(contents, "\n")
	fmt.Println(split)
	//split = split[:len(split)-1]

	parts := strings.Split(split[0], " ")
	partsX := strings.Split(parts[2], "..")
	partsY := strings.Split(parts[3], "..")
	targetXMin, _ := strconv.Atoi(partsX[0][2:])
	targetXMax, _ := strconv.Atoi(partsX[1][:len(partsX[1])-1])
	targetYMin, _ := strconv.Atoi(partsY[0][2:])
	targetYMax, _ := strconv.Atoi(partsY[1])

	possibleXVelMin := 0
	for x := 0; x < targetXMin; x += possibleXVelMin {
		possibleXVelMin++
	}

	highestY := 0
	hits := 0
	for yVel := targetYMin; yVel <= 0-targetYMin; yVel++ {
		for xVel := possibleXVelMin; xVel <= targetXMax; xVel++ {
			p := Probe{0, 0, xVel, yVel}
			highest := 0
			for p.CanStillHitTarget(targetXMin, targetXMax, targetYMin, targetYMax) {
				p.Tick()
				if p.VelY == 0 {
					highest = p.CoordY
				}
				if p.HitTarget(targetXMin, targetXMax, targetYMin, targetYMax) {
					hits++
					if highest > highestY {
						highestY = highest
					}
					break
				}
			}
		}
	}
	fmt.Println(highestY)
	fmt.Println(hits)
	return 0
}

type Probe struct {
	CoordX, CoordY int
	VelX, VelY     int
}

func (p *Probe) Tick() {
	p.CoordX += p.VelX
	p.CoordY += p.VelY
	if p.VelX > 0 {
		p.VelX--
	}
	p.VelY--
}

func (p *Probe) CanStillHitTarget(xMin, xMax, yMin, yMax int) bool {
	return p.CoordX <= xMax && p.CoordY >= yMin
}

func (p *Probe) HitTarget(xMin, xMax, yMin, yMax int) bool {
	return p.CoordX <= xMax && p.CoordX >= xMin && p.CoordY <= yMax && p.CoordY >= yMin
}
