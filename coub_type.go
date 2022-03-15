package main

import "time"

type CoubInfo struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
	Coubs      []Coub `json:"coubs"`
}

type Coub struct {
	Flag                   bool          `json:"flag"`
	Abuses                 interface{}   `json:"abuses"`
	RecoubsByUsersChannels []interface{} `json:"recoubs_by_users_channels"`
	Favourite              bool          `json:"favourite"`
	PromotedID             interface{}   `json:"promoted_id"`
	Recoub                 bool          `json:"recoub"`
	Like                   bool          `json:"like"`
	Dislike                bool          `json:"dislike"`
	Reaction               interface{}   `json:"reaction"`
	InMyBest2015           bool          `json:"in_my_best2015"`
	ID                     int           `json:"id"`
	Type                   string        `json:"type"`
	Permalink              string        `json:"permalink"`
	Title                  string        `json:"title"`
	VisibilityType         string        `json:"visibility_type"`
	OriginalVisibilityType string        `json:"original_visibility_type"`
	ChannelID              int           `json:"channel_id"`
	CreatedAt              time.Time     `json:"created_at"`
	UpdatedAt              time.Time     `json:"updated_at"`
	IsDone                 bool          `json:"is_done"`
	ViewsCount             int           `json:"views_count"`
	Cotd                   interface{}   `json:"cotd"`
	CotdAt                 interface{}   `json:"cotd_at"`
	VisibleOnExploreRoot   bool          `json:"visible_on_explore_root"`
	VisibleOnExplore       bool          `json:"visible_on_explore"`
	Featured               bool          `json:"featured"`
	Published              bool          `json:"published"`
	PublishedAt            time.Time     `json:"published_at"`
	Reversed               bool          `json:"reversed"`
	FromEditorV2           bool          `json:"from_editor_v2"`
	IsEditable             bool          `json:"is_editable"`
	OriginalSound          bool          `json:"original_sound"`
	HasSound               bool          `json:"has_sound"`
	RecoubTo               interface{}   `json:"recoub_to"`
	FileVersions           struct {
		HTML5 struct {
			Video struct {
				Higher struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"higher"`
				High struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"high"`
				Med struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"med"`
			} `json:"video"`
			Audio struct {
				High struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"high"`
				Med struct {
					URL  string `json:"url"`
					Size int    `json:"size"`
				} `json:"med"`
				SampleDuration float64 `json:"sample_duration"`
			} `json:"audio"`
		} `json:"html5"`
		Mobile struct {
			Video string   `json:"video"`
			Audio []string `json:"audio"`
		} `json:"mobile"`
		Share struct {
			Default string `json:"default"`
		} `json:"share"`
	} `json:"file_versions"`
	AudioVersions struct {
	} `json:"audio_versions"`
	ImageVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
	} `json:"image_versions"`
	FirstFrameVersions struct {
		Template string   `json:"template"`
		Versions []string `json:"versions"`
	} `json:"first_frame_versions"`
	Dimensions struct {
		Big []int `json:"big"`
		Med []int `json:"med"`
	} `json:"dimensions"`
	SiteWH               []int       `json:"site_w_h"`
	PageWH               []int       `json:"page_w_h"`
	SiteWHSmall          []int       `json:"site_w_h_small"`
	Size                 []int       `json:"size"`
	AgeRestricted        bool        `json:"age_restricted"`
	AgeRestrictedByAdmin bool        `json:"age_restricted_by_admin"`
	NotSafeForWork       bool        `json:"not_safe_for_work"`
	AllowReuse           bool        `json:"allow_reuse"`
	DontCrop             bool        `json:"dont_crop"`
	Banned               bool        `json:"banned"`
	GlobalSafe           bool        `json:"global_safe"`
	AudioFileURL         interface{} `json:"audio_file_url"`
	ExternalDownload     interface{} `json:"external_download"`
	Application          interface{} `json:"application"`
	Channel              struct {
		ID                     int           `json:"id"`
		Permalink              string        `json:"permalink"`
		Title                  string        `json:"title"`
		Description            string        `json:"description"`
		IFollowHim             bool          `json:"i_follow_him"`
		FollowsByUsersChannels []interface{} `json:"follows_by_users_channels"`
		FollowersCount         int           `json:"followers_count"`
		FollowingCount         int           `json:"following_count"`
		AvatarVersions         struct {
			Template string   `json:"template"`
			Versions []string `json:"versions"`
		} `json:"avatar_versions"`
	} `json:"channel"`
	File            interface{} `json:"file"`
	Picture         string      `json:"picture"`
	TimelinePicture string      `json:"timeline_picture"`
	SmallPicture    string      `json:"small_picture"`
	SharingPicture  interface{} `json:"sharing_picture"`
	PercentDone     int         `json:"percent_done"`
	Tags            []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Value string `json:"value"`
	} `json:"tags"`
	Categories []struct {
		ID                 int    `json:"id"`
		Title              string `json:"title"`
		Permalink          string `json:"permalink"`
		SubscriptionsCount int    `json:"subscriptions_count"`
		BigImageURL        string `json:"big_image_url"`
		SmallImageURL      string `json:"small_image_url"`
		MedImageURL        string `json:"med_image_url"`
		Visible            bool   `json:"visible"`
	} `json:"categories"`
	Communities []struct {
		ID                            int    `json:"id"`
		Title                         string `json:"title"`
		Permalink                     string `json:"permalink"`
		SubscriptionsCount            int    `json:"subscriptions_count"`
		BigImageURL                   string `json:"big_image_url"`
		SmallImageURL                 string `json:"small_image_url"`
		MedImageURL                   string `json:"med_image_url"`
		ISubscribed                   bool   `json:"i_subscribed"`
		CommunityNotificationsEnabled bool   `json:"community_notifications_enabled"`
		Description                   struct {
			ID              int           `json:"id"`
			Description     string        `json:"description"`
			Rules           []interface{} `json:"rules"`
			DescriptionHTML string        `json:"description_html"`
			RulesHTML       []interface{} `json:"rules_html"`
		} `json:"description"`
	} `json:"communities"`
	RecoubsCount  int `json:"recoubs_count"`
	RemixesCount  int `json:"remixes_count"`
	LikesCount    int `json:"likes_count"`
	DislikesCount int `json:"dislikes_count"`
	//RawVideoID           int  `json:"raw_video_id"`
	UploadedByIosApp     bool `json:"uploaded_by_ios_app"`
	UploadedByAndroidApp bool `json:"uploaded_by_android_app"`
	MediaBlocks          struct {
		//UploadedRawVideos []interface{} `json:"uploaded_raw_videos"`
		ExternalRawVideos []struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			URL         string `json:"url"`
			Image       string `json:"image"`
			ImageRetina string `json:"image_retina"`
			Meta        struct {
				Service  string `json:"service"`
				Duration string `json:"duration"`
			} `json:"meta"`
			Duration float64 `json:"duration"`
			//RawVideoID int     `json:"raw_video_id"`
			HasEmbed bool `json:"has_embed"`
		} `json:"external_raw_videos"`
		RemixedFromCoubs []interface{} `json:"remixed_from_coubs"`
		ExternalVideo    struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			URL         string `json:"url"`
			Image       string `json:"image"`
			ImageRetina string `json:"image_retina"`
			Meta        struct {
				Service  string `json:"service"`
				Duration string `json:"duration"`
			} `json:"meta"`
			Duration float64 `json:"duration"`
			//RawVideoID int     `json:"raw_video_id"`
			HasEmbed bool `json:"has_embed"`
		} `json:"external_video"`
	} `json:"media_blocks"`
	RawVideoThumbnailURL string      `json:"raw_video_thumbnail_url"`
	RawVideoTitle        string      `json:"raw_video_title"`
	VideoBlockBanned     bool        `json:"video_block_banned"`
	Duration             float64     `json:"duration"`
	PromoWinner          bool        `json:"promo_winner"`
	PromoWinnerRecoubers interface{} `json:"promo_winner_recoubers"`
	EditorialInfo        struct {
	} `json:"editorial_info"`
	PromoHint              interface{}   `json:"promo_hint"`
	BeelineBest2014        interface{}   `json:"beeline_best_2014"`
	FromWebEditor          bool          `json:"from_web_editor"`
	NormalizeSound         bool          `json:"normalize_sound"`
	NormalizeChangeAllowed bool          `json:"normalize_change_allowed"`
	Best2015Addable        bool          `json:"best2015_addable"`
	AhmadPromo             interface{}   `json:"ahmad_promo"`
	PromoData              interface{}   `json:"promo_data"`
	AudioCopyrightClaim    interface{}   `json:"audio_copyright_claim"`
	AdsDisabled            bool          `json:"ads_disabled"`
	IsSafeForAds           bool          `json:"is_safe_for_ads"`
	MegafonContents        []interface{} `json:"megafon_contents"`
	PositionOnPage         int           `json:"position_on_page"`
}

type ExternalDownload struct {
	Type        string `json:"type"`
	ServiceName string `json:"service_name"`
	URL         string `json:"url"`
	HasEmbed    bool   `json:"has_embed"`
}
