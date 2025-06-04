package impl

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickHouseStatisticsStorageImpl struct {
	db driver.Conn
}

func NewClickHouseStatisticsStorageImpl(conn driver.Conn) (*ClickHouseStatisticsStorageImpl, error) {
	return &ClickHouseStatisticsStorageImpl{conn}, nil
}

func (c *ClickHouseStatisticsStorageImpl) Close() {
	c.db.Close()
}
