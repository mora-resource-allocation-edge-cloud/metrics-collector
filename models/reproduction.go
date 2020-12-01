package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VideoReproduction struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	StartTime     primitive.DateTime `bson:"startTime" json:"startTime"`
	EndTime       primitive.DateTime `bson:"endTime" json:"endTime"`
	StartTimeUnix uint64             `bson:"startTimeUnix" json:"startTimeUnix"`
	EndTimeUnix   uint64             `bson:"endTimeUnix" json:"endTimeUnix"`
	Metrics       []Metric           `bson:"metrics" json:"metrics"`
}

type Metric struct {
	Video MetricContent `bson:"video" json:"video"`
	Audio MetricContent `bson:"audio" json:"audio"`
	Time  uint64        `bson:"time" json:"time"`
}

type MetricContent struct {
	Bitrate       float64    `bson:"bitrate" json:"bitrate"`
	BufferLevel   float64    `bson:"bufferLevel" json:"bufferLevel"`
	Download      LowAvgHigh `bson:"download" json:"download"`
	DroppedFrames int64      `bson:"droppedFrames" json:"droppedFrames"`
	Latency       LowAvgHigh `bson:"latency" json:"latency"`
	LiveLatency   float64    `bson:"liveLatency" json:"liveLatency"`
	MaxIndex      int64      `bson:"maxIndex" json:"maxIndex"`
	Ratio         LowAvgHigh `bson:"ratio" json:"ratio"`
	Time          uint64     `bson:"time" json:"time"`
	VideoTime     float64    `bson:"videoTime" json:"videoTime"`
}

type LowAvgHigh struct {
	Low  float64 `bson:"low" json:"low"`
	High float64 `bson:"high" json:"high"`
	Avg  float64 `bson:"avg" json:"avg"`
}
