package models

type UserRoom struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	CreatedAt    int    `bson:"created_at"`
	UpdatedAt    int    `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}
