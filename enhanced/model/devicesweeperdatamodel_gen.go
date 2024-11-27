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
	deviceSweeperDataModel interface {
		Insert(ctx context.Context, data *DeviceSweeperData) (*mongo.InsertOneResult, error)
		FindOne(ctx context.Context, id int64) (*DeviceSweeperData, error)
		FindOneByDeviceSnTimestamp(ctx context.Context, deviceSn string, timestamp int64) (*DeviceSweeperData, error)
		FindByDeviceSnTimeRange(ctx context.Context, deviceSn string, startTimestamp int64, endTimestamp int64) ([]*DeviceSweeperData, error)
		Update(ctx context.Context, data *DeviceSweeperData) (*mongo.UpdateResult, error)
		Upsert(ctx context.Context, data []*DeviceSweeperData) (*BatchResult, error)
		Delete(ctx context.Context, id int64)  error
	}

	defaultDeviceSweeperDataModel struct {
		conn  *mongo.Client
		table string
	}

	DeviceSweeperData struct {
		Id           int64   `bson:"id"`
		DeviceSn     string  `bson:"device_sn"`
		Timestamp    int64   `bson:"timestamp"`
		IsCharging   int64   `bson:"is_charging"`
		BatteryLevel uint64  `bson:"battery_level"`
		PositionX    float64 `bson:"position_x"`
		PositionY    float64 `bson:"position_y"`
	}
)

func newDeviceSweeperDataModel(conn *mongo.Client) *defaultDeviceSweeperDataModel {
	return &defaultDeviceSweeperDataModel{
		conn:  conn,
		table: "device_sweeper_data",
	}
}

func (m *defaultDeviceSweeperDataModel) Delete(ctx context.Context, id int64)  error {
	filter := bson.M{"id": id}
	_, err := m.conn.Database("test").Collection(m.table).DeleteOne(ctx, filter)
	return  err
}

func (m *defaultDeviceSweeperDataModel) FindOne(ctx context.Context, id int64) (*DeviceSweeperData, error) {
	filter := bson.M{"id": id}
	var resp DeviceSweeperData
	err := m.conn.Database("test").Collection(m.table).FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDeviceSweeperDataModel) FindOneByDeviceSnTimestamp(ctx context.Context, deviceSn string, timestamp int64) (*DeviceSweeperData, error) {
	filter := bson.M{"device_sn": deviceSn, "timestamp": timestamp}
	var resp DeviceSweeperData
	err := m.conn.Database("test").Collection(m.table).FindOne(ctx, filter).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultDeviceSweeperDataModel) FindByDeviceSnTimeRange(ctx context.Context, deviceSn string, startTimestamp int64, endTimestamp int64) ([]*DeviceSweeperData, error) {
	filter := bson.M{"device_sn": deviceSn, "timestamp": bson.M{"$gte": startTimestamp, "$lte": endTimestamp}}
	cursor, err := m.conn.Database("test").Collection(m.table).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []*DeviceSweeperData
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *defaultDeviceSweeperDataModel) Insert(ctx context.Context, data *DeviceSweeperData) (*mongo.InsertOneResult, error) {
	res, err := m.conn.Database("test").Collection(m.table).InsertOne(ctx, data)
	return res, err
}

func (m *defaultDeviceSweeperDataModel) Update(ctx context.Context, newData *DeviceSweeperData) (*mongo.UpdateResult, error) {
	filter := bson.M{"id": newData.Id}
	update := bson.M{"$set": newData}
	res, err := m.conn.Database("test").Collection(m.table).UpdateOne(ctx, filter, update)
	return res, err
}

func (m *defaultDeviceSweeperDataModel) Upsert(ctx context.Context, data []*DeviceSweeperData) (*BatchResult, error) {
	if len(data) == 0 {
		return &BatchResult{}, nil
	}

	result := &BatchResult{}

	for _, item := range data {
		filter := bson.M{"id": item.Id}
		update := bson.M{"$set": item}
		_, err := m.conn.Database("test").Collection(m.table).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
		if err != nil {
			//result.FailedBatch = append(result.FailedBatch, item.Id)
			result.Err = err
			return result, err
		}
		result.SuccessCount++
	}

	return result, nil
}

func (m *defaultDeviceSweeperDataModel) tableName() string {
	return m.table
}