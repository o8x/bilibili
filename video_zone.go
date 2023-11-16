package bilibili

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

type ZoneArchive struct {
	Aid       int    `json:"aid"`
	Videos    int    `json:"videos"`
	Tid       int    `json:"tid"`
	Tname     string `json:"tname"`
	Copyright int    `json:"copyright"`
	Pic       string `json:"pic"`
	Title     string `json:"title"`
	Pubdate   int    `json:"pubdate"`
	Ctime     int    `json:"ctime"`
	Desc      string `json:"desc"`
	State     int    `json:"state"`
	Duration  int    `json:"duration"`
	MissionId int    `json:"mission_id,omitempty"`
	Rights    struct {
		Bp            int `json:"bp"`
		Elec          int `json:"elec"`
		Download      int `json:"download"`
		Movie         int `json:"movie"`
		Pay           int `json:"pay"`
		Hd5           int `json:"hd5"`
		NoReprint     int `json:"no_reprint"`
		Autoplay      int `json:"autoplay"`
		UgcPay        int `json:"ugc_pay"`
		IsCooperation int `json:"is_cooperation"`
		UgcPayPreview int `json:"ugc_pay_preview"`
		NoBackground  int `json:"no_background"`
		ArcPay        int `json:"arc_pay"`
		PayFreeWatch  int `json:"pay_free_watch"`
	} `json:"rights"`
	Owner struct {
		Mid  int64  `json:"mid"`
		Name string `json:"name"`
		Face string `json:"face"`
	} `json:"owner"`
	Stat struct {
		Aid      int `json:"aid"`
		View     int `json:"view"`
		Danmaku  int `json:"danmaku"`
		Reply    int `json:"reply"`
		Favorite int `json:"favorite"`
		Coin     int `json:"coin"`
		Share    int `json:"share"`
		NowRank  int `json:"now_rank"`
		HisRank  int `json:"his_rank"`
		Like     int `json:"like"`
		Dislike  int `json:"dislike"`
		Vt       int `json:"vt"`
		Vv       int `json:"vv"`
	} `json:"stat"`
	Dynamic   string `json:"dynamic"`
	Cid       int    `json:"cid"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
	SeasonId    int         `json:"season_id,omitempty"`
	ShortLinkV2 string      `json:"short_link_v2"`
	FirstFrame  string      `json:"first_frame"`
	PubLocation string      `json:"pub_location"`
	Bvid        string      `json:"bvid"`
	SeasonType  int         `json:"season_type"`
	IsOgv       bool        `json:"is_ogv"`
	OgvInfo     interface{} `json:"ogv_info"`
	RcmdReason  string      `json:"rcmd_reason"`
	EnableVt    int         `json:"enable_vt"`
	AiRcmd      interface{} `json:"ai_rcmd"`
	UpFromV2    int         `json:"up_from_v2,omitempty"`
}

type ZoneNewVideos struct {
	Page struct {
		Num   int `json:"num"`
		Size  int `json:"size"`
		Count int `json:"count"`
	} `json:"page"`
	Archives []ZoneArchive `json:"archives"`
}

// GetZoneNewVideos 获取分区最新视频
func GetZoneNewVideos(tid, pageNum, pagesize int) (*ZoneNewVideos, error) {
	return std.GetZoneNewVideos(tid, pageNum, pagesize)
}

func (c *Client) GetZoneNewVideos(tid, pageNum, pagesize int) (*ZoneNewVideos, error) {
	resp, err := c.resty().R().
		SetQueryParams(map[string]string{
			"rid": strconv.Itoa(tid),
			"pn":  strconv.Itoa(pageNum),
			"ps":  strconv.Itoa(pagesize),
		}).
		Get("https://api.bilibili.com/x/web-interface/dynamic/region")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := getRespData(resp, "获取分区最新视频")
	if err != nil {
		return nil, err
	}

	var ret ZoneNewVideos
	err = json.Unmarshal(data, &ret)
	return &ret, errors.WithStack(err)
}

// GetZoneNewVideosWithTag 获取分区最新视频
func GetZoneNewVideosWithTag(tid int, tagID string, pageNum, pagesize int) (*ZoneNewVideos, error) {
	return std.GetZoneNewVideosWithTag(tid, tagID, pageNum, pagesize)
}

// 近期投稿  https://api.bilibili.com/x/web-interface/newlist_rank?main_ver=v3&search_type=video&view_type=hot_rank&copy_right=-1&new_web_tag=1&order=scores&cate_id=236&page=1&pagesize=30&time_from=20231105&time_to=20231112&keyword=%E5%B0%84%E5%87%BB
// 最新动态 https://api.bilibili.com/x/web-interface/dynamic/tag?ps=14&pn=1&rid=236&tag_id=6697
// 热门 https://api.bilibili.com/x/web-interface/ranking/tag?tag_id=6697&rid=236
func (c *Client) GetZoneNewVideosWithTag(tid int, tagID string, pageNum, pagesize int) (*ZoneNewVideos, error) {
	resp, err := c.resty().R().
		SetQueryParams(map[string]string{
			"tag_id": tagID,
			"rid":    strconv.Itoa(tid),
			"pn":     strconv.Itoa(pageNum),
			"ps":     strconv.Itoa(pagesize),
		}).
		Get("https://api.bilibili.com/x/web-interface/dynamic/tag")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := getRespData(resp, "获取分区标签最新视频")
	if err != nil {
		return nil, err
	}

	var ret ZoneNewVideos
	err = json.Unmarshal(data, &ret)
	return &ret, errors.WithStack(err)
}
