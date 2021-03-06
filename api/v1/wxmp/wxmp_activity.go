package wxmp

import (
    "RemindMe/global"
    models "RemindMe/model"
    "RemindMe/model/common/response"
    "RemindMe/model/wxmp"
    wxmpReq "RemindMe/model/wxmp/request"
    wxmpRes "RemindMe/model/wxmp/response"
    "RemindMe/utils"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "strconv"
    "time"
)

type ActivityApi struct {
}

func (api *ActivityApi) CreateActivity(c *gin.Context) {
    var req wxmpReq.ActivityCreateRequest
    _ = c.ShouldBindJSON(&req)

    id := utils.GetWxmpUserID(c)
    if id == 0 {
        response.FailWithMessage("创建活动失败", c)
        return
    }

    t, err := time.ParseInLocation("2006-01-02 15:04", req.Time, time.Local)
    if err != nil {
        global.Log.Error("解析时间失败", zap.Any("err", err))
        response.FailWithMessage("创建活动失败", c)
        return
    }
    var ac = wxmp.Activity{
        Type:      req.Type,
        Title:     req.Title,
        Time:      models.LocalTime{Time: t},
        Lunar:     req.Lunar,
        Periodic:  req.Periodic,
        NWeek:     req.NWeek,
        Address:   req.Location.Address,
        Latitude:  req.Location.Latitude,
        Longitude: req.Location.Longitude,
        Privacy:   req.Privacy,
        Remark:    req.Remark,
    }

    if err := activityService.CreateActivity(id, &ac); err != nil {
        response.FailWithMessage("创建活动失败", c)
        return
    }

    response.OkWithData(nil, c)
}

func (api *ActivityApi) ActivityList(c *gin.Context) {
    user := utils.GetWxmpUserInfo(c)
    if user.ID == 0 {
        response.FailWithMessage("获取用户id失败", c)
        return
    }

    // 通过用户id获取所有相关的活动
    acType, cursor := c.Query("type"), c.Query("cursor")
    activities, nextCursor, err := activityService.QueryActivities(user.ID, acType, cursor)
    if err != nil {
        response.FailWithMessage("获取活动列表失败", c)
        return
    }

    var list = make([]wxmpRes.ActivityResponse, 0)
    for _, ac := range activities {
        res := wxmpRes.ActivityResponse{
            Id:          ac.ID,
            SubId:       ac.SubId,
            Type:        ac.Type,
            Title:       ac.Title,
            TimeText:    ac.Time.Format("2006-01-02 15:04") + " " + getWeekdayString(ac.NWeek),
            DateTime:    models.LocalTime{Time: ac.Time.Time},
            Lunar:       ac.Lunar,
            Periodic:    ac.Periodic,
            NWeek:       ac.NWeek,
            ObviousDate: getObviousDate(ac.Time.Time),
            ObviousTime: getObviousTime(ac.Time.Time),
            Location: wxmpRes.ActivityAddr{
                Address:   ac.Address,
                Latitude:  ac.Latitude,
                Longitude: ac.Longitude,
            },
            Publisher: wxmpRes.ActivityUser{
                Id:     ac.Publisher.ID,
                Name:   ac.Publisher.Nickname,
                Avatar: ac.Publisher.Avatar,
                Phone:  ac.Publisher.Phone,
            },
            MySubIndex: -1,
        }

        // 当前用户是否为发布者
        if res.Publisher.Id == user.ID {
            res.IsPublisher = true
        }
        // 当前用户是否订阅此活动
        for idx, item := range ac.Subscriptions {
            res.Subscribers = append(res.Subscribers, wxmpRes.ActivityUser{
                Id:     item.ID,
                Name:   item.Subscriber.Nickname,
                Avatar: item.Subscriber.Avatar,
                Phone:  item.Subscriber.Phone,
            })
            if item.ID == user.ID {
                res.IsSubscribed = true
                res.MySubIndex = idx
            }
        }
        list = append(list, res)
    }
    response.OkWithData(&wxmpRes.ActivitiesResponse{List: list, Cursor: nextCursor}, c)
}

func (api *ActivityApi) UpdateActivity(c *gin.Context) {
    var req wxmpReq.ActivityUpdateRequest
    _ = c.ShouldBindJSON(&req)

    t, err := time.ParseInLocation(global.SecLocalTimeFormat, req.DateTime, time.Local)
    if err != nil {
        global.Log.Error("解析时间失败", zap.Any("err", err))
        response.FailWithMessage("更新活动失败", c)
        return
    }
    var ac = wxmp.Activity{
        Type:      req.Type,
        Title:     req.Title,
        Time:      models.LocalTime{Time: t},
        Lunar:     req.Lunar,
        Periodic:  req.Periodic,
        NWeek:     req.NWeek,
        Address:   req.Location.Address,
        Latitude:  req.Location.Latitude,
        Longitude: req.Location.Longitude,
        Privacy:   req.Privacy,
        Remark:    req.Remark,
    }
    if err = activityService.UpdateActivity(req.Id, &ac); err != nil {
        response.FailWithMessage("更新失败", c)
        return
    }
    response.OkWithMessage("更新成功", c)
}

func (api *ActivityApi) ActivityDetail(c *gin.Context) {
    activityId, err := strconv.Atoi(c.Query("id"))
    subId, err := strconv.Atoi(c.Query("subId"))
    if err != nil {
        response.FailWithMessage("活动id不存在", c)
        return
    }
    user := utils.GetWxmpUserInfo(c)
    if user.ID == 0 {
        response.FailWithMessage("获取用户id失败", c)
        return
    }

    // 查询活动信息
    ac, err := activityService.ActivityDetail(uint(activityId), uint(subId))
    if err != nil {
        response.FailWithMessage("活动不存在", c)
        return
    }

    res := wxmpRes.ActivityResponse{
        Id:          ac.ID,
        SubId:       ac.SubId,
        Type:        ac.Type,
        Title:       ac.Title,
        TimeText:    ac.Time.Format("2006-01-02 15:04") + " " + getWeekdayString(ac.NWeek),
        DateTime:    models.LocalTime{Time: ac.Time.Time},
        Lunar:       ac.Lunar,
        Periodic:    ac.Periodic,
        NWeek:       ac.NWeek,
        Week:        getWeekdayString(ac.NWeek),
        ObviousDate: getObviousDate(ac.Time.Time),
        ObviousTime: getObviousTime(ac.Time.Time),
        Location: wxmpRes.ActivityAddr{
            Address:   ac.Address,
            Latitude:  ac.Latitude,
            Longitude: ac.Longitude,
        },
        Remark:       ac.Remark,
        Privacy:      ac.Privacy,
        IsPublisher:  false,
        IsSubscribed: false,
        MySubIndex:   -1,
        Publisher: wxmpRes.ActivityUser{
            Id:     ac.Publisher.ID,
            Name:   ac.Publisher.Nickname,
            Avatar: ac.Publisher.Avatar,
            Phone:  ac.Publisher.Phone,
        },
        Subscribers: nil,
    }
    // 当前用户是否为发布者
    if res.Publisher.Id == user.ID {
        res.IsPublisher = true
    }
    var subscriber []wxmpRes.ActivityUser
    for idx, item := range ac.Subscriptions {
        if item.Status == 0 {
            continue
        }
        subscriber = append(subscriber, wxmpRes.ActivityUser{
            Id:          item.Subscriber.ID,
            Name:        item.Subscriber.Nickname,
            Avatar:      item.Subscriber.Avatar,
            Phone:       item.Subscriber.Phone,
            IsTplRemind: item.IsTplRemind,
            IsSmsRemind: item.IsSmsRemind,
            RemindAt:    item.RemindAt,
        })
        // 订阅者id与用户id相同，表示当前用户已订阅活动
        if item.Subscriber.ID == user.ID {
            res.IsSubscribed = true
            res.MySubIndex = idx
        }
    }
    res.Subscribers = subscriber

    response.OkWithData(&res, c)
}

func (api *ActivityApi) DeleteActivity(c *gin.Context) {
    var req wxmpReq.ActivityDeleteRequest
    _ = c.ShouldBindJSON(&req)
    id := utils.GetWxmpUserID(c)

    if err := activityService.DeleteActivity(id, req.Id); err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    response.Ok(c)
}

func (api *ActivityApi) SubscribeActivity(c *gin.Context) {
    id := utils.GetWxmpUserID(c)

    var req wxmpReq.ActivitySubscribeRequest
    _ = c.ShouldBindJSON(&req)

    if err := activityService.SubscribeActivity(id, &req); err != nil {
        response.FailWithMessage("订阅成功", c)
        return
    }
    response.OkWithMessage("订阅成功", c)
}

func (api *ActivityApi) UnsubscribeActivity(c *gin.Context) {
    var (
        id     int
        err    error
        userId = utils.GetWxmpUserID(c)
    )

    activityId := c.Query("id")
    if id, err = strconv.Atoi(activityId); err != nil {
        response.FailWithMessage("取消订阅活动失败", c)
        return
    }

    if err = activityService.UnsubscribeActivity(userId, uint(id)); err != nil {
        response.FailWithMessage("取消订阅失败", c)
        return
    }
    response.OkWithMessage("取消订阅成功", c)
}

func getObviousDate(tim time.Time) (d string) {
    days := diffDays(time.Now(), tim)
    half := getObviousHalfDate(tim)
    //nweek := tim.Weekday()
    //nextSunday := time.Now()
    //lastMonday := time.Now()
    switch days {
    case -2:
        d = "前天" + half
    case -1:
        d = "昨天" + half
    case 0:
        d = "今天" + half
    case 1:
        d = "明天" + half
    case 2:
        d = "后天" + half
    default:
        d = strconv.Itoa(int(tim.Month())) + "月" + strconv.Itoa(tim.Day()) + "日"
    }
    return d
}

func getObviousTime(tim time.Time) (t string) {
    return tim.Format("15:04")
}

func getObviousHalfDate(t time.Time) string {
    var h, m = t.Hour(), t.Minute()
    if h >= 0 && h < 6 {
        return "凌晨"
    } else if h >= 6 && h < 8 {
        return "早上"
    } else if h >= 8 && h < 11 || (h == 11 && m < 30) {
        return "上午"
    } else if h >= 13 && h < 18 || (h == 12 && m > 30) {
        return "下午"
    } else if h >= 18 && h < 24 {
        return "晚上"
    } else {
        return "中午"
    }
}

func diffDays(t1, t2 time.Time) (days int) {
    var (
        df = t2.Sub(t1)
        du = time.Hour * 24
    )

    // >1天时间间隔
    if df > du || -df > du {
        tmp := int(df / du)
        days += tmp
        t1 = t1.Add(time.Duration(tmp) * du)
    }

    // 不到一天的时间内，如果跨天，days+1
    if t1.Format("20060102") != t2.Format("20060102") {
        if df > 0 {
            days += 1
        } else {
            days -= 1
        }
    }

    return days
}

func getWeekdayString(n int) string {
    var w = map[int]string{
        0: "星期日",
        1: "星期一",
        2: "星期二",
        3: "星期三",
        4: "星期四",
        5: "星期五",
        6: "星期六",
    }
    return w[n]
}
