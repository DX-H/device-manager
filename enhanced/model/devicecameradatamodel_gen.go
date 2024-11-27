// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type (
	deviceCameraDataModel interface {
		Insert(ctx context.Context, data *DeviceCameraData) (*mongo.InsertOneResult, error)
		FindOne(ctx context.Context, id int64) (*DeviceCameraData, error)
		FindOneByDeviceSnTimestamp(ctx context.Context, deviceSn string, timestamp int64) (*DeviceCameraData, error)
		FindByDeviceSnTimeRange(ctx context.Context, deviceSn string, startTimestamp int64, endTimestamp int64) ([]*DeviceCameraData, error)
		Update(ctx context.Context, data *DeviceCameraData) (*mongo.UpdateResult, error)
		Upsert(ctx context.Context, data []*DeviceCameraData)  (*BatchResult, error)
		Delete(ctx context.Context, id int64)  error
	}

	defaultDeviceCameraDataModel struct {
		conn  *mongo.Client
		table string
	}

	DeviceCameraData struct {
		Id           int64   `bson:"id"`
		DeviceSn     string  `bson:"device_sn"`
		Timestamp    int64   `bson:"timestamp"`
		IsFixed      int64   `bson:"is_fixed"`
		BatteryLevel uint64  `bson:"battery_level"`
		RotationX    float64 `bson:"rotation_x"`
		RotationY    float64 `bson:"rotation_y"`
		RotationZ    float64 `bson:"rotation_z"`
	}
)

func newDeviceCameraDataModel(conn *mongo.Client) *defaultDeviceCameraDataModel {
	return &defaultDeviceCameraDataModel{
		conn:  conn,
		table: "device_camera_data",
	}
}

func (m *defaultDeviceCameraDataModel) Delete(ctx context.Context, id int64) error {
	collection := m.conn.Database("test").Collection(m.table)
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (m *defaultDeviceCameraDataModel) FindOne(ctx context.Context, id int64) (*DeviceCameraData, error) {
	collection := m.conn.Database("test").Collection(m.table)
	filter := bson.D{{"id", id}}
	var result DeviceCameraData
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *defaultDeviceCameraDataModel) FindOneByDeviceSnTimestamp(ctx context.Context, deviceSn string, timestamp int64) (*DeviceCameraData, error) {
	collection := m.conn.Database("test").Collection(m.table)
	filter := bson.D{{"device_sn", deviceSn}, {"timestamp", timestamp}}
	var result DeviceCameraData
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *defaultDeviceCameraDataModel) FindByDeviceSnTimeRange(ctx context.Context, deviceSn string, startTimestamp int64, endTimestamp int64) ([]*DeviceCameraData, error) {
	collection := m.conn.Database("test").Collection(m.table)
	filter := bson.D{{"device_sn", deviceSn}, {"timestamp", bson.D{{"$gte", startTimestamp}, {"$lte", endTimestamp}}}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []*DeviceCameraData
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *defaultDeviceCameraDataModel) Insert(ctx context.Context, data *DeviceCameraData) (*mongo.InsertOneResult, error) {
	collection := m.conn.Database("test").Collection(m.table)
	result, err := collection.InsertOne(ctx, data)
	return result, err
}

func (m *defaultDeviceCameraDataModel) Update(ctx context.Context, newData *DeviceCameraData) (*mongo.UpdateResult, error) {
	collection := m.conn.Database("test").Collection(m.table)
	filter := bson.D{{"id", newData.Id}}
	update := bson.D{{"$set", bson.D{{"device_sn", newData.DeviceSn}, {"timestamp", newData.Timestamp}, {"is_fixed", newData.IsFixed}, {"battery_level", newData.BatteryLevel}, {"rotation_x", newData.RotationX}, {"rotation_y", newData.RotationY}, {"rotation_z", newData.RotationZ}}}}
	result, err := collection.UpdateOne(ctx, filter, update)
	return result, err
}

func (m *defaultDeviceCameraDataModel) Upsert(ctx context.Context, data []*DeviceCameraData)  (*BatchResult, error) {
	if len(data) == 0 {
		return &BatchResult{}, nil
	}

	result := &BatchResult{}
	collection := m.conn.Database("test").Collection(m.table)
	for _, item := range data {
		filter := bson.D{{"id", item.Id}}
		update := bson.D{{"$set", bson.D{{"device_sn", item.DeviceSn}, {"timestamp", item.Timestamp}, {"is_fixed", item.IsFixed}, {"battery_level", item.BatteryLevel}, {"rotation_x", item.RotationX}, {"rotation_y", item.RotationY}, {"rotation_z", item.RotationZ}}}}
		upsert := true
		options := options.Update().SetUpsert(upsert)
		_, err := collection.UpdateOne(ctx, filter, update, options)
                if err != nil {
		    //session.AbortTransaction(ctx)
		    //result.FailedBatch = append(result.FailedBatch, i/BatchSize)
                    result.Err = err
                    return result, err
		}
	}
	/*
        err = session.CommitTransaction(ctx)
        if err != nil {
            return nil, err
        }
	*/

        result.SuccessCount  = 1
	return result, nil
}

func (m *defaultDeviceCameraDataModel) tableName() string {
	return m.table
}