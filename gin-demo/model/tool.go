package model

type Tool struct {
    ID string `json:"tool_id" xorm:"'id' not null pk char(32) comment('uuid')"`// ID
    CategoryId string `json:"category_id" xorm:"'category_id' not null char(32)"`
    Type uint8 `json:"type" xorm:"'type' not null tinyint(3)"`
    Title string `json:"title" xorm:"'title' not null varchar(8)"`
    Icon string `json:"icon" xorm:"'icon' not null varchar(128)"`
    Uri string `json:"uri" xorm:"'uri' not null default '' varchar(128)"`
    Action string `json:"action" xorm:"'action' default '' varchar(8)"`
    ParamStr string `json:"param_str" xorm:"'param' default '' varchar(255)"`
    Param interface{} `json:"param"`
    Status uint8 `json:"status" xorm:"'status' not null tinyint(1)"`
    Sort uint32 `json:"sort" xorm:"'sort' not null int(10)"`
    CreatedAt uint32 `json:"created_at" xorm:"'created_at' not null int(10)"` // 创建时间
    UpdatedAt uint32 `json:"updated_at" xorm:"'updated_at' not null int(10)"` // 最后更新时间
}

// DatabaseAlias 数据库别名
func (*Tool) DatabaseAlias() string {
    return "lx_toolkit"
}

// TableName 表名
func (*Tool) TableName() string {
    return "lx_tool"
}