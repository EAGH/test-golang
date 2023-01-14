package models

type ArrayObj struct {
	Obj []string `json:"obj" bson:"obj"`
}
type ResponseRandom struct {
	//Categories []string `json:"categories" bson:"categories"`
	//CreatedAt string `json:"created_at" bson:"created_at"`
	//IcoUrl string `json:"icon_url" bson:"icon_url"`
	ID string `json:"id" bson:"id"`
	//UpdatedAt string `json:"updated_at" bson:"updated_at"`
	//Url string `json:"url" bson:"url"`
	//Value string `json:"value" bson:"value"`
}

type MessageResponseHTTP struct {
	Status  int         `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Content interface{} `json:"content" bson:"content"`
}
