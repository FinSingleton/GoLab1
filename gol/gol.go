package main

func createNewWorld(p golParams) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for i := 0; i < p.imageHeight; i++ {
		newWorld[i] = make([]byte, p.imageWidth)
	}
	return newWorld
}

func calculateAliveNeighbours(p golParams, x int, y int, world [][]byte) int {
	sum := 0

	/*for j := -1; j <= 1; j++ {
		for l := -1; l <= 1; l++ {
			ny := y + j
			nx := x + l

			if j == 0 && l == 0 {
				continue
			}

			if nx < 0 {
				nx = p.imageWidth - 1
			}
			if nx >= p.imageHeight {
				nx = 0
			}

			if ny < 0 {
				ny = p.imageHeight - 1
			}
			if ny >= p.imageHeight {
				ny = 0
			}

			if world[ny][nx] == 255 {
				sum++
			}
		}
	}*/

	nx := x
	ny := y

	for j := -1; j <= 1; j++ {
		for l := -1; l <= 1; l++ {
			if !(j == 0 && l == 0) {

				if (y+j < 0) || (y+j > 15) || (x+l < 0) || (x+l > 15) {
					if y+j < 0 {
						ny = 15
					} else if y+j > 15 {
						ny = 0
					}

					if x+l < 0 {
						nx = 15
					} else if x+l > 15 {
						nx = 0
					}

					if nx != x && ny != y {
						if world[ny][nx] == 255 {
							sum += 1
						}
					}
					if nx != x {
						if world[ny][nx+l] == 255 {
							sum += 1
						}
					}
					if ny != y {
						if world[ny+j][nx] == 255 {
							sum += 1
						}
					}
					nx = x
					ny = y

				} else if world[y+j][x+l] == 255 {
					sum += 1
				}
			}
		}
	}

	return sum
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	//create nextWorld, new world entity to put the new cells into
	newWorld := createNewWorld(p)

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {

				if calculateAliveNeighbours(p, x, y, world) == 2 || calculateAliveNeighbours(p, x, y, world) == 3 {
					newWorld[y][x] = 255
				}
			} else {
				if calculateAliveNeighbours(p, x, y, world) == 3 {
					newWorld[y][x] = 255
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	alive := []cell{}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				alive = append(alive, cell{x, y})
			}
		}
	}

	return alive
}
