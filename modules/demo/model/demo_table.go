package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameDemoTable = "demo_table"

// DemoTable mapped from table <demo_table>
type DemoTable struct {
	*cool.Model
	// Name string `gorm:"column:name;not null;comment:名称" json:"name"`
	Columnstring  string  `gorm:"column:columnstring;not null;comment:列string" json:"columnstring"`
	Columnint     int     `gorm:"column:columnint;not null;comment:列int" json:"columnint"`
	ColumnBool    bool    `gorm:"column:columnbool;not null;comment:列bool" json:"columnbool"`
	ColumnFloat32 float32 `gorm:"column:columnfloat32;not null;comment:列float32" json:"columnfloat32"`
	ColumnFloat64 float64 `gorm:"column:columnfloat64;not null;comment:列float" json:"columnfloat64"`
}

// TableName DemoTable's table name
func (*DemoTable) TableName() string {
	return TableNameDemoTable
}

// GroupName DemoTable's table group
func (*DemoTable) GroupName() string {
	return "default"
}

// NewDemoTable create a new DemoTable
func NewDemoTable() *DemoTable {
	return &DemoTable{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&DemoTable{})
}
