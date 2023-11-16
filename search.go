package bilibili

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Spi struct {
	B3 string `json:"b_3"`
	B4 string `json:"b_4"`
}

// GetFingerSpi 获取SPI指纹
func GetFingerSpi() (*Spi, error) {
	return std.GetFingerSpi()
}

func (c *Client) GetFingerSpi() (*Spi, error) {
	resp, err := c.resty().R().Get("https://api.bilibili.com/x/frontend/finger/spi")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := getRespData(resp, "获取SPI指纹")
	if err != nil {
		return nil, err
	}

	var ret Spi
	err = json.Unmarshal(data, &ret)
	return &ret, errors.WithStack(err)
}

type SearchArchive struct {
	Type             string        `json:"type"`
	Id               int           `json:"id"`
	Author           string        `json:"author"`
	Mid              int64         `json:"mid"`
	Typeid           string        `json:"typeid"`
	Typename         string        `json:"typename"`
	Arcurl           string        `json:"arcurl"`
	Aid              int           `json:"aid"`
	Bvid             string        `json:"bvid"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	Arcrank          string        `json:"arcrank"`
	Pic              string        `json:"pic"`
	Play             int           `json:"play"`
	VideoReview      int           `json:"video_review"`
	Favorites        int           `json:"favorites"`
	Tag              string        `json:"tag"`
	Review           int           `json:"review"`
	Pubdate          int           `json:"pubdate"`
	Senddate         int           `json:"senddate"`
	Duration         string        `json:"duration"`
	Badgepay         bool          `json:"badgepay"`
	HitColumns       []string      `json:"hit_columns"`
	ViewType         string        `json:"view_type"`
	IsPay            int           `json:"is_pay"`
	IsUnionVideo     int           `json:"is_union_video"`
	RecTags          interface{}   `json:"rec_tags"`
	NewRecTags       []interface{} `json:"new_rec_tags"`
	RankScore        int           `json:"rank_score"`
	Like             int           `json:"like"`
	Upic             string        `json:"upic"`
	Corner           interface{}   `json:"corner"`
	Cover            string        `json:"cover"`
	Desc             string        `json:"desc"`
	Url              string        `json:"url"`
	RecReason        string        `json:"rec_reason"`
	Danmaku          int           `json:"danmaku"`
	BizData          interface{}   `json:"biz_data"`
	IsChargeVideo    int           `json:"is_charge_video"`
	Vt               int           `json:"vt"`
	EnableVt         int           `json:"enable_vt"`
	VtDisplay        string        `json:"vt_display"`
	Subtitle         string        `json:"subtitle"`
	EpisodeCountText string        `json:"episode_count_text"`
	ReleaseStatus    int           `json:"release_status"`
	IsIntervene      int           `json:"is_intervene"`
}

type SearchResult struct {
	Seid           string `json:"seid"`
	Page           int    `json:"page"`
	Pagesize       int    `json:"pagesize"`
	NumResults     int    `json:"numResults"`
	NumPages       int    `json:"numPages"`
	SuggestKeyword string `json:"suggest_keyword"`
	RqtType        string `json:"rqt_type"`
	CostTime       struct {
		Total               string `json:"total"`
		ParamsCheck         string `json:"params_check"`
		IsRiskQuery         string `json:"is_risk_query"`
		IllegalHandler      string `json:"illegal_handler"`
		MainHandler         string `json:"main_handler"`
		AsRequestFormat     string `json:"as_request_format"`
		AsRequest           string `json:"as_request"`
		DeserializeResponse string `json:"deserialize_response"`
		AsResponseFormat    string `json:"as_response_format"`
	} `json:"cost_time"`
	Pageinfo struct {
		Video struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"video"`
		Bangumi struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"bangumi"`
		Special struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"special"`
		Topic struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"topic"`
		Upuser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"upuser"`
		Tv struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"tv"`
		Movie struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"movie"`
		MediaBangumi struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"media_bangumi"`
		MediaFt struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"media_ft"`
		RelatedSearch struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"related_search"`
		User struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"user"`
		Activity struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"activity"`
		OperationCard struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"operation_card"`
		Pgc struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"pgc"`
		Live struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live"`
		LiveAll struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_all"`
		BiliUser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"bili_user"`
		LiveRoom struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_room"`
		Article struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"article"`
		LiveUser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_user"`
		LiveMaster struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_master"`
	} `json:"pageinfo"`
	TopTlist struct {
		Video         int `json:"video"`
		Bangumi       int `json:"bangumi"`
		Special       int `json:"special"`
		Topic         int `json:"topic"`
		Upuser        int `json:"upuser"`
		Tv            int `json:"tv"`
		Movie         int `json:"movie"`
		Card          int `json:"card"`
		MediaBangumi  int `json:"media_bangumi"`
		MediaFt       int `json:"media_ft"`
		Pgc           int `json:"pgc"`
		Live          int `json:"live"`
		User          int `json:"user"`
		Activity      int `json:"activity"`
		OperationCard int `json:"operation_card"`
		BiliUser      int `json:"bili_user"`
		LiveRoom      int `json:"live_room"`
		Article       int `json:"article"`
		LiveUser      int `json:"live_user"`
		LiveMaster    int `json:"live_master"`
	} `json:"top_tlist"`
	ShowColumn       int      `json:"show_column"`
	ShowModuleList   []string `json:"show_module_list"`
	AppDisplayOption struct {
		IsSearchPageGrayed int `json:"is_search_page_grayed"`
	} `json:"app_display_option"`
	InBlackKey int `json:"in_black_key"`
	InWhiteKey int `json:"in_white_key"`
	Result     []struct {
		ResultType string          `json:"result_type"`
		Data       []SearchArchive `json:"data"`
	} `json:"result"`
	IsSearchPageGrayed int `json:"is_search_page_grayed"`
}

// SearchByKeyword 搜索
func SearchByKeyword(keyword string, pageNum int) (*SearchResult, error) {
	return std.SearchByKeyword(keyword, pageNum)
}

func (c *Client) SearchByKeyword(keyword string, pageNum int) (*SearchResult, error) {
	spi, err := c.GetFingerSpi()
	if err != nil {
		return nil, err
	}

	resp, err := c.resty().R().
		SetQueryParams(map[string]string{
			"page":    strconv.Itoa(pageNum),
			"keyword": keyword,
		}).
		SetCookies([]*http.Cookie{
			{Name: "SESSDATA", Value: ""},
			{Name: "buvid3", Value: spi.B3},
			{Name: "bili_jct", Value: ""},
			{Name: "DedeUserID", Value: ""},
			{Name: "ac_time_value", Value: ""},
			{Name: "Domain", Value: ".bilibili.com"},
		}).
		Get("https://api.bilibili.com/x/web-interface/search/all/v2")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := getRespData(resp, "获取搜索结果")
	if err != nil {
		return nil, err
	}

	var ret *SearchResult
	err = json.Unmarshal(data, &ret)
	return ret, errors.WithStack(err)
}

type Game5012 struct {
	G5012 []struct {
		Id         int    `json:"id"`
		ContractId string `json:"contract_id"`
		ResId      int    `json:"res_id"`
		AsgId      int    `json:"asg_id"`
		PosNum     int    `json:"pos_num"`
		Name       string `json:"name"`
		Pic        string `json:"pic"`
		Litpic     string `json:"litpic"`
		Url        string `json:"url"`
		Style      int    `json:"style"`
		Archive    struct {
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
				Mid  int    `json:"mid"`
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
			SeasonId    int    `json:"season_id,omitempty"`
			ShortLinkV2 string `json:"short_link_v2"`
			FirstFrame  string `json:"first_frame"`
			PubLocation string `json:"pub_location"`
			Bvid        string `json:"bvid"`
			EnableVt    int    `json:"enable_vt"`
			MissionId   int    `json:"mission_id,omitempty"`
			UpFromV2    int    `json:"up_from_v2,omitempty"`
		} `json:"archive"`
		Agency       string      `json:"agency"`
		Label        string      `json:"label"`
		Intro        string      `json:"intro"`
		CreativeType int         `json:"creative_type"`
		RequestId    string      `json:"request_id"`
		SrcId        int         `json:"src_id"`
		Area         int         `json:"area"`
		IsAdLoc      bool        `json:"is_ad_loc"`
		AdCb         string      `json:"ad_cb"`
		Title        string      `json:"title"`
		ServerType   int         `json:"server_type"`
		CmMark       int         `json:"cm_mark"`
		Stime        int         `json:"stime"`
		Mid          string      `json:"mid"`
		ActivityType int         `json:"activity_type"`
		Epid         int         `json:"epid"`
		Season       interface{} `json:"season"`
		Room         interface{} `json:"room"`
		SubTitle     string      `json:"sub_title"`
		AdDesc       string      `json:"ad_desc"`
		AdverName    string      `json:"adver_name"`
		NullFrame    bool        `json:"null_frame"`
		PicMainColor string      `json:"pic_main_color"`
		CardType     int         `json:"card_type"`
		BusinessMark interface{} `json:"business_mark"`
		Inline       struct {
			InlineUseSame       int    `json:"inline_use_same"`
			InlineType          int    `json:"inline_type"`
			InlineUrl           string `json:"inline_url"`
			InlineBarrageSwitch int    `json:"inline_barrage_switch"`
		} `json:"inline"`
		Operater string `json:"operater"`
	} `json:"5012"`
}

func GetGame5012() (*Game5012, error) {
	return std.GetGame5012()
}

// https://api.bilibili.com/x/web-show/res/locs?pf=0&ids=4991,5012,5028,255,257,259
// 5012 前方高能
func (c *Client) GetGame5012() (*Game5012, error) {
	resp, err := c.resty().R().
		SetQueryParams(map[string]string{
			"ids": "5012",
			"pf":  "0",
		}).
		Get("https://api.bilibili.com/x/web-show/res/locs")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	data, err := getRespData(resp, "获取前方高能最新视频")
	if err != nil {
		return nil, err
	}

	var ret Game5012
	err = json.Unmarshal(data, &ret)
	return &ret, errors.WithStack(err)
}
