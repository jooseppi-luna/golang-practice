package battleship

import "fmt"

type Pair struct {
    x, y interface{}
}

type Ship struct {
	length    int
	shotsLeft int
}

type BattleBoard map[Pair]*Ship

func Shoot(battleBoard *BattleBoard, coords Pair) bool {
	// if coord is in map, decrement ship count
	_, shipHit := (*battleBoard)[coords]
	if shipHit {
		(*battleBoard)[coords].shotsLeft--
		return true
	}

	return false
}

func ShipDown(battleBoard *BattleBoard, coords Pair) (bool, error) {
	// if coord is not in map, return error
	_, isShip := (*battleBoard)[coords]
	if isShip {
		if (*battleBoard)[coords].shotsLeft == 0 {
			return true, nil
		}
		return false, nil
	}
	
	return false, fmt.Errorf("No ship at %+v", coords)
}

