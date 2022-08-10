package models

type MessageBasic struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreatedAt    int    `bson:"created_at"`
	UpdatedAt    int    `bson:"updated_at"`
}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}
