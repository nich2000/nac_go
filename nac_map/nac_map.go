//==============================================================================
//==============================================================================
package nac_map

const MAP_HEIGHT = 100
const MAP_WIDTH = 100
const MAP_NEIGHBORS_COUNT = 8

type cell_t struct {
  x    int
  y    int
  kind int
}

type terrain_t [MAP_HEIGHT][MAP_WIDTH]cell_t

type neighbors_t [MAP_NEIGHBORS_COUNT]cell_t

var terrain terrain_t

var neighbors_cells neighbors_t = neighbors_t{
  &cell_t{-1, -1, 0},
  &cell_t{-1, 0, 0},
  &cell_t{-1, 1, 0},
  &cell_t{0, -1, 0},
  &cell_t{0, 1, 0},
  &cell_t{1, -1, 0},
  &cell_t{1, 0, 0},
  &cell_t{1, 1, 0},
}

func map_init(terrain terrain_t, file string) {
}

func map_load(terrain terrain_t, file string) {
}

func map_save(terrain terrain_t, file string) {
}

func map_neighbors(terrain terrain_t, cell cell_t) neighbors_t {
  var res neighbors_t

  for i := 0; i < MAP_NEIGHBORS_COUNT; i++ {
  }
}
