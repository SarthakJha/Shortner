package shortner

import "go.mongodb.org/mongo-driver/bson/primitive"

// CreateURL ...
type CreateURL struct {
	// ID      string `bson:"_id, omitempty" json:"id"`
	URLMain string `json:"url_main" bson:"url_main"`
	ShortID string `json:"short_id" bson:"short_id"`
}

// SendURL ...
type SendURL struct {
	ID      string `bson:"_id, omitempty" json:"id"`
	URLMain string `json:"url_main" bson:"url_main"`
	ShortID string `json:"short_id" bson:"short_id"`
}

// URLRequest ...
type URLRequest struct {
	URLMain string `json:"url_main" bson:"url_main"`
	Event   string `json:"event" bson:"event"`
}

// DecodeRequest ...
type DecodeRequest struct {
	ID      primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	URLMain string             `json:"url_main" bson:"url_main"`
	ShortID string             `json:"short_id" bson:"short_id"`
}
