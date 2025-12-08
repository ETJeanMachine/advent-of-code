package day7

import (
	"errors"
	"log"
)

// Struct representing a tachyon beam within the manifold; with
// it's current end being stored.
type TachyonBeam struct {
	row, col int
}

// Struct representing the manifold.
type TachyonManifold struct {
	width, height int
	splitters     [][]bool
	beams         map[TachyonBeam]bool
}

// Creates a new manifold with an inital tachyon beam and the array containing
// where the splitter are/are not.
func NewManifold(init TachyonBeam, splitters [][]bool) *TachyonManifold {
	beams := make(map[TachyonBeam]bool)
	beams[init] = true
	return &TachyonManifold{
		width:     len(splitters[0]),
		height:    len(splitters),
		splitters: splitters,
		beams:     beams,
	}
}

// Progresses the tachyon beam forward in time two ticks - the amount needed to hit the next
// splitter. Returns the new beams possible from this point.
func (tm *TachyonManifold) progressBeam(beam TachyonBeam) ([]TachyonBeam, error) {
	newBeams := []TachyonBeam{}
	if beam.row < tm.height-2 {
		newBeams = append(newBeams, TachyonBeam{beam.row + 2, beam.col})
		if tm.splitters[beam.row+2][beam.col] {
			newBeams[0].col -= 1
			newBeams = append(newBeams, TachyonBeam{beam.row + 2, beam.col + 1})
		}
		return newBeams, nil
	}
	return newBeams, errors.New("Beam progressed beyond the height of the manifold!")
}

// Ticks the tachyon manifold forward in time, counting the number of times the tachyon
// beam split.
func (tm *TachyonManifold) tick() int {
	timesSplit := 0
	newBeams := []TachyonBeam{}
	for beam := range tm.beams {
		beams, err := tm.progressBeam(beam)
		if err != nil {
			log.Fatalf("Error when ticking: %s", err.Error())
		}
		if len(beams) > 1 {
			timesSplit += 1
		}
		newBeams = append(newBeams, beams...)
	}
	for tb := range tm.beams {
		delete(tm.beams, tb)
	}
	for _, beam := range newBeams {
		tm.beams[beam] = true
	}
	return timesSplit
}

// Returns the count of splits that occur without timeline shenangians.
func (tm *TachyonManifold) CountSplits() int {
	totalSplit := 0
	for range tm.height/2 - 1 {
		totalSplit += tm.tick()
	}
	return totalSplit
}

// Calculates the number of possible end timelines for a tachyon manifold.
func (tm *TachyonManifold) CountTimelines() int {
	// We memoize the number of future timelines that can span out from
	// each beam to save a great deal of processing time.
	memoizedTime := make(map[TachyonBeam]int)
	var timelineHelper func(beam TachyonBeam) int
	// DFS helper that skips out on previously iterated paths via memoization.
	timelineHelper = func(beam TachyonBeam) int {
		var newBeams []TachyonBeam
		if timelineCount, ok := memoizedTime[beam]; ok {
			return timelineCount
		} else {
			beams, err := tm.progressBeam(beam)
			newBeams = beams
			if err != nil {
				memoizedTime[beam] = 1
				return 1
			}
		}
		timelineCount := 0
		for _, newBeam := range newBeams {
			timelineCount += timelineHelper(newBeam)
		}
		memoizedTime[beam] = timelineCount
		return timelineCount
	}
	// Since I'm using a map to store beams from previous problems;
	// i just iterate over breaking from the first to grab the
	// "first" beam.
	var initBeam TachyonBeam
	for beam := range tm.beams {
		initBeam = beam
		break
	}
	return timelineHelper(initBeam)
}
