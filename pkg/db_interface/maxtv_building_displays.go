package db_interface

type MaxtvBuildingDisplay struct {
	ID                 int    `gorm:"id" json:"id"`
	BuildingId         int    `json:"building_id"`
	Sysid              string `json:"sys_id"`
	Name               string `json:"name"`
	ThemeId            int    `json:"theme_id"`
	Type               string `json:"type"` //enum ('', 'touch', 'nontouch')
	CovidBlockActiveTo string `json:"covid_block_active_to"`
	DisplaySizeId      int    `json:"display_size_id" gorm:"column:display_size_id"`
	PlaceId            int    `json:"place_id"`

	//Orientation                    varchar(50)
	//DisplaySizeId                int
	//model_id                       int
	//in_elevator                    tinyint(1)
	//display_place                  varchar(255)
	//display_address                varchar(255)
	//display_address_style          text
	//video                          tinyint(1)
	//weather                        int
	//notices                        int
	//news                           int
	//sports                         int
	//gallery                        int
	//classified                     int
	//parcels                        int
	//transit                        int
	//survey                         int
	//fifa2014                       int
	//playlist2                      enum ('0', '1')
	//directory                      int
	//panam                          int
	//booking                        int
	//rio2016                        int
	//french                         int
	//events                         int
	//html_scroll                    int
	//resolution_x                   int
	//resolution_y                   int
	//app_rotation                   int
	//app_scale_x                    varchar(255)
	//app_scale_y                    varchar(255)
	//app_left                       varchar(255)
	//app_top                        varchar(255)
	//app_lock                       int
	//video_left                     int
	//video_top                      int
	//video_height                   int
	//video_width                    int
	//video_offset_top               int
	//video_offset_left              int
	//logo                           varchar(255)
	//logo_style                     text
	//logo_left                      varchar(255)
	//logo_top                       varchar(255)
	//logo_width                     varchar(255)
	//logo_height                    varchar(255)
	//theme_switcher                 text
	//playlist_survey_id             int
	//playlist_id                    int
	//playlist2_id                   int
	//playlist2_title                varchar(255)
	//playlist_classified_id         int
	//playlist_directory_id          int
	//playlist_amenities_id          int
	//playlist_path                  varchar(255)
	//domain_url                     varchar(255)
	//display_path                   varchar(255)
	//news_path                      varchar(255)
	//news_french_path               varchar(255)
	//sports_path                    varchar(255)
	//sports_french_path             varchar(255)
	//gallery_path                   varchar(255)
	//hourly_path                    varchar(255)
	//weather_path                   varchar(255)
	//transit_path                   varchar(255)
	//parcels_path                   varchar(255)
	//panam_path                     varchar(255)
	//fifa2014_path                  varchar(255)
	//blink_path                     varchar(255)
	//blinkmd5_path                  varchar(255)
	//booking_path                   varchar(255)
	//residents_path                 varchar(255)
	//survey_path                    varchar(255)
	//video_swf                      varchar(255)
	//createdon                      datetime
	//route_lock                     enum ('0', '1')
	//accessibility                  int
	//default_notice_replacement     varchar(255)
	//display_delay_classifieds      int
	//display_delay_news             int
	//active                         int
	//buildinglink_user              varchar(255)
	//buildinglink_password          varchar(255)
	//buildinglink_device            varchar(255)
	//buildinglink_unit_sufix        varchar(255)
	//internal_parcel                int
	//weather_farinhate              int
	//directory_address              varchar(255)
	//buttons_autonavigation         int
	//marge_parcels                  int
	//place                          varchar(255)
	//survey_builder_id              int
	//autorotation                   text
	//presentation                   int
	//presentation_displays          text
	//clickevent                     varchar(255)
	//background_feature_notices     varchar(255)
	//background_feature_news        varchar(255)
	//background_feature_sports      varchar(255)
	//background_feature_classified  varchar(255)
	//background_feature_weather     varchar(255)
	//background_feature_parcels     varchar(255)
	//like_feature_news              int
	//like_feature_sports            int
	//like_feature_notices           int
	//debug_mode                     int
	//debug_url                      varchar(255)
	//touch_rolling                  int
	//touch_with_arrows              int
	//debug_touch_layer              int
	//playlist_new_pmi               int
	//playlist_classified_new_pmi    int
	//playlist_id_new_pmi            int unsigned
	//playlist_classified_id_new_pmi int
	//directory_with_notices         int
	//global_style                   longtext
	//directory_alpha_sort           int
	//is_android                     int
	//allow_interactive_ad           int
	//layout_path                    varchar(255)
	//playlist2_new_pmi              int
	//playlist2_id_new_pmi           int
	//video_randomize                int
	//place_id                       int
	//covid_block_active_to          varchar(50)
	//time_download_config           int
	//time_download_themes           int
	//time_download_layout           int
	//time_download_backgrounds      int
	//time_download_ads              int
	//time_download_ads_json         int
	//time_download_booking          int
	//time_download_classified       int
	//time_download_directory        int
	//time_download_notices          int
	//time_download_weather          int
	//time_download_media            int
	//time_download_parcels          int
	//time_download_survey           int
	//time_upload_screenshots        int
	//time_upload_frameshots         int
	//time_upload_stat_video         int
	//time_upload_stat_health        int
	//time_upload_stat_clicks        int
	//time_upload_stat_survey        int
}
