package models

type RoomBasic struct {
	Number       string `bson:"number"`
	Name         string `bson:"name"`
	Info         string `bson:"info"`
	UserIdentity string `bson:"user_identity"`
	CreatedAt    int    `bson:"created_at"`
	UpdatedAt    int    `bson:"updated_at"`
}

func (RoomBasic) CollectionName() string {
	return "room_basic"
}
