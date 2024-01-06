package sson

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectId(value string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(value)
}

func StringToObjectIdArray(values []string) ([]primitive.ObjectID, error) {
	objectIDs := make([]primitive.ObjectID, 0, len(values))
	for _, id := range values {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIDs = append(objectIDs, objectID)
	}
	return objectIDs, nil
}
