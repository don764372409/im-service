package entity

type Message struct {
	No       string `gorm:"<-:create" gorm:"index:idx_1"  json:"no"` //消息唯一编号  4个字段的联合索引
	Type     string `gorm:"index:idx_1,unique" json:"type"`          //聊天类型 friend、group  4个字段的联合索引
	TargetId uint64 `gorm:"index:idx_1,unique" json:"targetId"`      //聊天目标 用户ID或群ID  4个字段的联合索引
	UserId   uint64 `gorm:"index:idx_1,unique" json:"userId"`        //当前聊天所有者  4个字段的联合索引
	From     string `json:"from"`                                    //消息发送者
	Data     string `json:"data"`                                    //消息内容 MessageData的json字符串
	Time     uint64 `json:"time"`                                    //消息发送时间
	Read     int    `json:"read"`                                    //是否已阅读 1-未读 2-已读
	Send     int    `json:"send"`                                    //发送状态 1-发送中 2-发送成功 -1-发送失败
	Ext1     string `json:"ext1"`
	Ext2     string `json:"ext2"`
	Ext3     string `json:"ext3"`
	Ext4     int    `json:"ext4"`
	Ext5     int    `json:"ext5"`
}

type MessageData struct {
	Type    int    `json:"type"`    //1-文本 2-图片 3-语音 4-视频 5-文件
	Content string `json:"content"` //消息内容
}
