package room

import "context"

type Room struct{}

var rooms = map[string]struct{}{"econom": {}, "standart": {}, "lux": {}}

func New() *Room {
	return &Room{}
}

func (r Room) GetAll(ctx context.Context) (map[string]struct{}, error) {
	res := make(map[string]struct{}, len(rooms))

	for room := range rooms {
		res[room] = struct{}{}
	}

	return res, nil
}
