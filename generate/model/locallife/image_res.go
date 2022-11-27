package locallife

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
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


CREATE TABLE `image_res` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '图片 id',
  `name` varchar(255) NOT NULL COMMENT '图片名称',
  `scene` varchar(255) NOT NULL DEFAULT '' COMMENT '使用场景',
  `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '图片地址',
  `is_delete` tinyint DEFAULT '0' COMMENT '是否删除(0:未删除, 1:已删除)',
  `status` tinyint DEFAULT '0' COMMENT '0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `key_name` (`name`),
  KEY `key_scene` (`scene`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='图片资源表'

JSON Sample
-------------------------------------
{    "id": 8,    "name": "tIIvETAIpKqesWBvEeMTOpVLt",    "scene": "JCWHqhvZytjSwCSwRFhjqHMFU",    "url": "aEWfsYqeSpbxHsWpjkwTvrBVD",    "is_delete": 58,    "status": 1,    "created_at": "2083-04-14T01:10:33.050461054+08:00",    "updated_at": "2072-08-17T07:55:49.233181871+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// ImageRes struct is a row record of the image_res table in the local_life database
type ImageRes struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 图片 id
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"` // 图片名称
	//[ 2] scene                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Scene string `gorm:"column:scene;type:varchar;size:255;" json:"scene"` // 使用场景
	//[ 3] url                                            varchar(1024)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	URL string `gorm:"column:url;type:varchar;size:1024;" json:"url"` // 图片地址
	//[ 4] is_delete                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsDelete sql.NullInt64 `gorm:"column:is_delete;type:tinyint;default:0;" json:"is_delete"` // 是否删除(0:未删除, 1:已删除)
	//[ 5] status                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Status sql.NullInt64 `gorm:"column:status;type:tinyint;default:0;" json:"status"` // 0未审核 1审核通过 2审核通过正在编辑 3审核不过 4自动审核过
	//[ 6] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	//[ 7] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间

}

var image_resTableInfo = &TableInfo{
	Name: "image_res",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `图片 id`,
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
			Comment:            `图片名称`,
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
			Name:               "scene",
			Comment:            `使用场景`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Scene",
			GoFieldType:        "string",
			JSONFieldName:      "scene",
			ProtobufFieldName:  "scene",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (i *ImageRes) TableName() string {
	return "image_res"
}

//// BeforeSave invoked before saving, return an error if field is not populated.
//func (i *ImageRes) BeforeSave() error {
//	return nil
//}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *ImageRes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *ImageRes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (i *ImageRes) TableInfo() *TableInfo {
	return image_resTableInfo
}
