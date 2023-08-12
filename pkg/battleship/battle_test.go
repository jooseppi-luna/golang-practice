package battleship

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func initBabyBoard() *BattleBoard {
	twoShip := Ship{length: 2, shotsLeft: 1}
	threeShip1 := Ship{length: 3, shotsLeft: 3}
	threeShip2 := Ship{length: 3, shotsLeft: 0}
	
	battleBoard := BattleBoard{
		Pair{3, 5}: &twoShip,
		Pair{3, 6}: &twoShip,
		Pair{2, 1}: &threeShip1,
		Pair{3, 1}: &threeShip1,
		Pair{4, 1}: &threeShip1,
		Pair{9, 4}: &threeShip2,
		Pair{9, 5}: &threeShip2,
		Pair{9, 6}: &threeShip2,
	}

	return &battleBoard
}

func initBattleBoard() *BattleBoard {
	twoShip := Ship{length: 2, shotsLeft: 2}
	threeShip1 := Ship{length: 3, shotsLeft: 3}
	threeShip2 := Ship{length: 3, shotsLeft: 3}
	fourShip := Ship{length: 4, shotsLeft: 4}
	fiveShip := Ship{length: 5, shotsLeft: 5}
	
	battleBoard := BattleBoard{
		Pair{3, 5}: &twoShip,
		Pair{3, 6}: &twoShip,
		Pair{2, 1}: &threeShip1,
		Pair{3, 1}: &threeShip1,
		Pair{4, 1}: &threeShip1,
		Pair{9, 4}: &threeShip2,
		Pair{9, 5}: &threeShip2,
		Pair{9, 6}: &threeShip2,
		Pair{5, 6}: &fourShip,
		Pair{5, 7}: &fourShip,
		Pair{5, 8}: &fourShip,
		Pair{5, 9}: &fourShip,
		Pair{2, 3}: &fiveShip,
		Pair{3, 3}: &fiveShip,
		Pair{4, 3}: &fiveShip,
		Pair{5, 3}: &fiveShip,
		Pair{6, 3}: &fiveShip,
	}

	return &battleBoard
}

//babyBoard := initBabyBoard()

func TestShoot(t *testing.T) {
	battleBoard := initBabyBoard()
        tests := []struct{
		name     string
                board    *BattleBoard
		coords   Pair
		hit      bool
		hitsLeft int
        }{
		{
			name:     "missed shot",
			board:    battleBoard,
			coords:   Pair{4, 7},
			hit:      false,
			hitsLeft: -1,
		},
		{
			name:  "good shot",
			board:  battleBoard,
			coords: Pair{4, 1},
			hit:    true,
			hitsLeft: 2,
		},
	}

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        hit := Shoot(tt.board, tt.coords)
                        assert.Equal(t, tt.hit, hit)
                        //assert.Equal(t, (*tt.board)[tt.coords].shotsLeft, tt.hitsLeft)
                })
        }
}

func TestShipDown(t *testing.T) {
	battleBoard := initBabyBoard()
        tests := []struct{
		name   string
                board  *BattleBoard
		coords Pair
		isDown bool
        }{
		{
			name:  "no ship at coords",
			board:  battleBoard,
			coords: Pair{4, 7},
			isDown: false,
		},
		{
			name:  "ship down",
			board:  battleBoard,
			coords: Pair{3, 1},
			isDown: false,
		},
		{
			name:  "ship not down",
			board:  battleBoard,
			coords: Pair{9, 5},
			isDown: true,
		},
	}

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        isDown, _ := ShipDown(tt.board, tt.coords)
                        assert.Equal(t, tt.isDown, isDown)
                })
        }
}

