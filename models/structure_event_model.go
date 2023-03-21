package models

type StructureEvent struct {
	ObjectId string `json:"_id" bson:"_id"`
	DeviceSn string `json:"device_sn" bson:"device_sn"`
}
