package model

type Answer struct {
	Id         int    `gorm:"id" json:"id"`
	Content    string `gorm:"content" json:"content"`         // 内容
	Username   string `gorm:"username" json:"username"`       // 用户
	Uid        int    `gorm:"uid" json:"uid"`                 // 用户id
	GoodCount  int    `gorm:"good_count" json:"good_count"`   // 点赞数
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	UpdateTime string `gorm:"update_time" json:"update_time"` // 更新时间
}

type Comment struct {
	Id         int    `gorm:"id" json:"id"`
	Content    string `gorm:"content" json:"content"`         // 内容
	Username   string `gorm:"username" json:"username"`       // 用户
	Uid        int    `gorm:"uid" json:"uid"`                 // 用户id
	GoodCount  int    `gorm:"good_count" json:"good_count"`   // 点赞数
	CreateTime string `gorm:"create_time" json:"create_time"` // 创建时间
	UpdateTime string `gorm:"update_time" json:"update_time"` // 更新时间
}

type Topic struct {
	Id          int    `gorm:"id" json:"id"`
	Title       string `gorm:"title" json:"title"`               // 标题
	Content     string `gorm:"content" json:"content"`           // 内容
	Username    string `gorm:"username" json:"username"`         // 用户
	Uid         int    `gorm:"uid" json:"uid"`                   // 用户id
	AnswerCount int    `gorm:"answer_count" json:"answer_count"` // 回答数量
	GoodCount   int    `gorm:"good_count" json:"good_count"`     // 点赞数
	Priority    int    `gorm:"priority" json:"priority"`         // 排序优先级
	CategoryId  int    `gorm:"category_id" json:"category_id"`   // 分类id
	CreateTime  string `gorm:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  string `gorm:"update_time" json:"update_time"`   // 更新时间
}
