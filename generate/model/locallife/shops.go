package locallife

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `shops` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商铺 id',
  `name` varchar(255) NOT NULL COMMENT '商铺名称',
  `type` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商铺类型 grid_option主键',
  `phone` varchar(255) NOT NULL DEFAULT '' COMMENT '电话号码',
  `address` varchar(1024) NOT NULL DEFAULT '' COMMENT '商铺地址',
  `business_hours` varchar(255) NOT NULL DEFAULT '' COMMENT '营业时间',
  `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址',
  `is_delete` tinyint DEFAULT '0' COMMENT '是否删除(0:未删除, 1:已删除)',
  `status` tinyint DEFAULT '0' COMMENT '0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `key_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商铺资源表'

JSON Sample
-------------------------------------
{    "id": 83,    "name": "BOXUpHjYSLeplpvLqCPuCYUjO",    "type": 45,    "phone": "NNQOXveZAKLNpAgexmROcebGY",    "address": "mPKimtkqTVcsMohXKjITOoeMn",    "business_hours": "FxNKKZMHemuGStJSCsPMCeAsY",    "url": "ZysYxvAveVtYbcoTBvSTagkQi",    "is_delete": 86,    "status": 89,    "created_at": "2152-12-10T09:51:19.599648914+08:00",    "updated_at": "2241-10-02T00:07:01.792362555+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned



*/

// Shops struct is a row record of the shops table in the local_life database
type Shops struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 商铺 id
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"` // 商铺名称
	//[ 2] type                                           ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	Type uint64 `gorm:"column:type;type:ubigint;default:0;" json:"type"` // 商铺类型 grid_option主键
	//[ 3] phone                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone string `gorm:"column:phone;type:varchar;size:255;" json:"phone"` // 电话号码
	//[ 4] address                                        varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	Address string `gorm:"column:address;type:varchar;size:1024;" json:"address"` // 商铺地址
	//[ 5] business_hours                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BusinessHours string `gorm:"column:business_hours;type:varchar;size:255;" json:"business_hours"` // 营业时间
	//[ 6] url                                            varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	URL string `gorm:"column:url;type:varchar;size:1024;" json:"url"` // 图片地址
	//[ 7] is_delete                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDelete sql.NullInt64 `gorm:"column:is_delete;type:tinyint;default:0;" json:"is_delete"` // 是否删除(0:未删除, 1:已删除)
	//[ 8] status                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status sql.NullInt64 `gorm:"column:status;type:tinyint;default:0;" json:"status"` // 0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过
	//[ 9] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	//[10] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间

}

var shopsTableInfo = &TableInfo{
	Name: "shops",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `商铺 id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            `商铺名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
			Comment:            `商铺类型 grid_option主键`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "uint64",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "phone",
			Comment:            `电话号码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Phone",
			GoFieldType:        "string",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "address",
			Comment:            `商铺地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "business_hours",
			Comment:            `营业时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "BusinessHours",
			GoFieldType:        "string",
			JSONFieldName:      "business_hours",
			ProtobufFieldName:  "business_hours",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "url",
			Comment:            `图片地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "URL",
			GoFieldType:        "string",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_delete",
			Comment:            `是否删除(0:未删除, 1:已删除)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsDelete",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "is_delete",
			ProtobufFieldName:  "is_delete",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "status",
			Comment:            `0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "created_at",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "updated_at",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Shops) TableName() string {
	return "shops"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Shops) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Shops) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Shops) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Shops) TableInfo() *TableInfo {
	return shopsTableInfo
}
