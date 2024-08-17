package utils

import (
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
)

const (
	BoardSize int64 = 24
	CellSize  int64 = 20
)

type Point struct {
	X int64
	Y int64
}

type Cell struct {
	Alive     bool
	NextState bool
	Neighbour int
}

type GameOfLife struct {
	Board [BoardSize][BoardSize]Cell
}

func NewGameOfLife() GameOfLife {
	return GameOfLife{}
}

func (b *GameOfLife) Init() {
	for i := int64(0); i < BoardSize; i++ {
		for j := int64(0); j < BoardSize; j++ {
			b.Board[i][j] = Cell{
				Alive: rand.Float64() < 0.2,
			}
		}
	}
}

func (b *GameOfLife) CheckNeighbour() {
	for i := int64(0); i < BoardSize; i++ {
		for j := int64(0); j < BoardSize; j++ {
			aliveNeighbours := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x >= 0 && y >= 0 && x < BoardSize && y < BoardSize && !(x == i && y == j) {
						if b.Board[x][y].Alive {
							aliveNeighbours++
						}
					}
				}
			}
			b.Board[i][j].Neighbour = aliveNeighbours
		}
	}
}

func (b *GameOfLife) TickStep() {
	b.CheckNeighbour()
	for i := int64(0); i < BoardSize; i++ {
		for j := int64(0); j < BoardSize; j++ {
			cell := &b.Board[i][j]
			if cell.Alive && (cell.Neighbour < 2 || cell.Neighbour > 3) {
				cell.NextState = false
			} else if !cell.Alive && cell.Neighbour == 3 {
				cell.NextState = true
			} else {
				cell.NextState = cell.Alive
			}
		}
	}

	for i := int64(0); i < BoardSize; i++ {
		for j := int64(0); j < BoardSize; j++ {
			b.Board[i][j].Alive = b.Board[i][j].NextState
		}
	}
}

func (b *GameOfLife) DrawToImage() *image.Paletted {
	h, w := int(BoardSize*CellSize), int(BoardSize*CellSize)
	palette := []color.Color{color.White, color.Black}
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette)

	for i := int64(0); i < BoardSize; i++ {
		for j := int64(0); j < BoardSize; j++ {
			var colIndex uint8
			if b.Board[i][j].Alive {
				colIndex = 1
			} else {
				colIndex = 0
			}

			for y := i * CellSize; y < (i+1)*CellSize; y++ {
				for x := j * CellSize; x < (j+1)*CellSize; x++ {
					img.SetColorIndex(int(x), int(y), colIndex)
				}
			}
		}
	}

	return img
}

func (b *GameOfLife) GenerateGIF(filename string, steps int, delay int) error {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < steps; i++ {
		images = append(images, b.DrawToImage())
		delays = append(delays, delay)
		b.TickStep()
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
