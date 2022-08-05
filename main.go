package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/follower", requestForFollower)
	http.HandleFunc("/roominfo", getRoomInfo)
	http.HandleFunc("/config", getconfig)
	http.ListenAndServe(":8081", nil)
}

func requestForFollower(w http.ResponseWriter, req *http.Request) {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/relation/stat?vmid=475210&jsonp=jsonp", nil)
	if err != nil {
		// handle err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	var marshalled Response
	err = json.NewDecoder(resp.Body).Decode(&marshalled)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%d", marshalled.Data.Follower)
}

func getconfig(w http.ResponseWriter, req *http.Request) {
	fd, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	cnf, err := io.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", string(cnf))
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Mid       int `json:"mid"`
		Following int `json:"following"`
		Whisper   int `json:"whisper"`
		Black     int `json:"black"`
		Follower  int `json:"follower"`
	} `json:"data"`
}

func getRoomInfo(w http.ResponseWriter, req *http.Request) {
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/space/acc/info?mid=475210&jsonp=jsonp", nil)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	req.Header.Set("Authority", "api.bilibili.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	defer resp.Body.Close()
	resptext, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var marshalled UserInfoResponse
	err = json.Unmarshal(resptext, &marshalled)
	if err != nil {
		log.Fatal(err)
	}
	marshalledJson, err := json.Marshal(&marshalled.Data.LiveRoom)
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Fprintf(w, "%s", marshalledJson)
}

type UserInfoResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Mid       int    `json:"mid"`
		Name      string `json:"name"`
		Sex       string `json:"sex"`
		Face      string `json:"face"`
		FaceNft   int    `json:"face_nft"`
		Sign      string `json:"sign"`
		Rank      int    `json:"rank"`
		Level     int    `json:"level"`
		Jointime  int    `json:"jointime"`
		Moral     int    `json:"moral"`
		Silence   int    `json:"silence"`
		Coins     int    `json:"coins"`
		FansBadge bool   `json:"fans_badge"`
		FansMedal struct {
			Show  bool `json:"show"`
			Wear  bool `json:"wear"`
			Medal struct {
				UID              int    `json:"uid"`
				TargetID         int    `json:"target_id"`
				MedalID          int    `json:"medal_id"`
				Level            int    `json:"level"`
				MedalName        string `json:"medal_name"`
				MedalColor       int    `json:"medal_color"`
				Intimacy         int    `json:"intimacy"`
				NextIntimacy     int    `json:"next_intimacy"`
				DayLimit         int    `json:"day_limit"`
				MedalColorStart  int    `json:"medal_color_start"`
				MedalColorEnd    int    `json:"medal_color_end"`
				MedalColorBorder int    `json:"medal_color_border"`
				IsLighted        int    `json:"is_lighted"`
				GuardLevel       int    `json:"guard_level"`
				LightStatus      int    `json:"light_status"`
				WearingStatus    int    `json:"wearing_status"`
				Score            int    `json:"score"`
			} `json:"medal"`
		} `json:"fans_medal"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		Vip struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path        string `json:"path"`
				Text        string `json:"text"`
				LabelTheme  string `json:"label_theme"`
				TextColor   string `json:"text_color"`
				BgStyle     int    `json:"bg_style"`
				BgColor     string `json:"bg_color"`
				BorderColor string `json:"border_color"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
		} `json:"vip"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`
			Name       string `json:"name"`
			Image      string `json:"image"`
			ImageSmall string `json:"image_small"`
			Level      string `json:"level"`
			Condition  string `json:"condition"`
		} `json:"nameplate"`
		UserHonourInfo struct {
			Mid    int           `json:"mid"`
			Colour interface{}   `json:"colour"`
			Tags   []interface{} `json:"tags"`
		} `json:"user_honour_info"`
		IsFollowed bool   `json:"is_followed"`
		TopPhoto   string `json:"top_photo"`
		Theme      struct {
		} `json:"theme"`
		SysNotice struct {
		} `json:"sys_notice"`
		LiveRoom struct {
			RoomStatus    int    `json:"roomStatus"`
			LiveStatus    int    `json:"liveStatus"`
			URL           string `json:"url"`
			Title         string `json:"title"`
			Cover         string `json:"cover"`
			Roomid        int    `json:"roomid"`
			RoundStatus   int    `json:"roundStatus"`
			BroadcastType int    `json:"broadcast_type"`
			WatchedShow   struct {
				Switch       bool   `json:"switch"`
				Num          int    `json:"num"`
				TextSmall    string `json:"text_small"`
				TextLarge    string `json:"text_large"`
				Icon         string `json:"icon"`
				IconLocation string `json:"icon_location"`
				IconWeb      string `json:"icon_web"`
			} `json:"watched_show"`
		} `json:"live_room"`
		Birthday string `json:"birthday"`
		School   struct {
			Name string `json:"name"`
		} `json:"school"`
		Profession struct {
			Name       string `json:"name"`
			Department string `json:"department"`
			Title      string `json:"title"`
			IsShow     int    `json:"is_show"`
		} `json:"profession"`
		Tags   interface{} `json:"tags"`
		Series struct {
			UserUpgradeStatus int  `json:"user_upgrade_status"`
			ShowUpgradeWindow bool `json:"show_upgrade_window"`
		} `json:"series"`
		IsSeniorMember int `json:"is_senior_member"`
	} `json:"data"`
}
