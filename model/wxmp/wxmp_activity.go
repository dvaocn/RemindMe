package wxmp

import (
    "RemindMe/model"
)

type Activity struct {
    models.Model
    SubId uint             `gorm:"-"` // 周期性活动子id
    Type  string           `json:"type" gorm:"column:type; type:varchar(16); comment:活动类型,db:代办,dk:打卡,sr:生日,sjb:时间表"`
    Title string           `json:"title" gorm:"column:title; type:varchar(64); comment:活动名称"`
    Time  models.LocalTime `json:"time" gorm:"column:time; type:datetime(3); comment:活动时间"`
    Lunar string           `json:"lunar" gorm:"column:lunar; type:varchar(32);comment:农历时间"`

    Periodic  int     `json:"periodic" gorm:"column:periodic; type:int(4); comment:周期间隔时间，0/1/7/30/365及355(农历每年)"`
    NWeek     int     `json:"nWeek" gorm:"column:nweek; type:int(4); comment:周期几，0-6"`
    Address   string  `json:"address" gorm:"column:address; type:varchar(128); comment:活动地址"`
    Latitude  float64 `json:"latitude" gorm:"column:latitude; comment:地址纬度"`
    Longitude float64 `json:"longitude" gorm:"column:longitude; comment:地址精度"`

    Privacy int    `json:"privacy" gorm:"column:privacy; type:int(4); comment:0:仅自己可见,1:可分享他人,2:完全公开"`
    Remark  string `json:"remark" gorm:"column:remark; type:varchar(512); comment:活动备注"`

    // 发布者信息
    Publisher User `gorm:"foreignKey:PublisherID"`
    // 发布者id
    PublisherID uint `gorm:"column:publisher_id; not null; index:idx_activity_user_id"`

    // 订阅者信息
    //Subscribers []User `gorm:"many2many:wxmp_activity_subscription"`

    // 订阅额外信息
    Subscriptions []ActivitySubscription `gorm:"foreignKey:ActivityID"`
}

func (Activity) TableName() string {
    return "wxmp_activity"
}
