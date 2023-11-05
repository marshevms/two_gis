package room

import "context"

type Room struct{}

var availableRooms = map[string]struct{}{"econom": {}, "standart": {}, "lux": {}}

func New() *Room {
	return &Room{}
}

func (r Room) GetAvailable(ctx context.Context) (map[string]struct{}, error) {
	res := make(map[string]struct{}, len(availableRooms))

	for room := range availableRooms {
		res[room] = struct{}{}
	}

	return res, nil
}
