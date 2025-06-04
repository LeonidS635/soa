package dto

import "time"

// Metric

const (
	MetricViews = iota
	MetricLikes
	MetricComments
)

// Statistics

type PostStats struct {
	Views    uint64
	Likes    uint64
	Comments uint64
}

// Dynamics

type PostDynamics []struct {
	Date        time.Time
	MetricValue uint64
}

// Top

type Top []struct {
	ID          int32 // Post or Author ID
	MetricValue uint64
}
