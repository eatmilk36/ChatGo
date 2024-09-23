package ChatroomMessageHistory

type ChatroomMessageHistoryRepository struct {
	Id        int    `json:"Id" gorm:"primaryKey"`
	UserId    int    `json:"UserId" gorm:"column:UserId"`
	GroupName string `json:"GroupName" gorm:"column:GroupName"`
	Message   string `json:"Message" gorm:"column:Message"`
	TimeStamp int64  `json:"TimeStamp" gorm:"column:TimeStamp"`
}

// TableName 指定資料表名稱為 'ChatroomMessageHistory'
func (ChatroomMessageHistoryRepository) TableName() string {
	return "ChatroomMessageHistory"
}
