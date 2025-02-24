package website

import "time"

type DataRegion struct {
	ID         uint   `gorm:"primarykey" json:"ID"`
	Pid        uint   `json:"pid" gorm:"index;comment:'父ID'"`
	Path       string `json:"path" gorm:"comment:'路径'"`
	Level      uint   `json:"level" gorm:"comment:'级别'"`
	Name       string `json:"name" gorm:"comment:'中文名称'"`
	NameEn     string `json:"name_en" gorm:"comment:'英文名称'"`
	NamePinyin string `json:"name_pinyin" gorm:"comment:'拼音名称'"`
	Code       string `json:"code" gorm:"comment:'代码'"`
}

func (DataRegion) TableName() string {
	return "data_region"
}

// CREATE TABLE `countries` (
// 	`id` mediumint unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`iso3` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`numeric_code` char(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`iso2` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`phonecode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`capital` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`currency` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`currency_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`currency_symbol` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`tld` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`native` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`region_id` mediumint unsigned DEFAULT NULL,
// 	`subregion` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`subregion_id` mediumint unsigned DEFAULT NULL,
// 	`nationality` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`timezones` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
// 	`translations` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
// 	`latitude` decimal(10,8) DEFAULT NULL,
// 	`longitude` decimal(11,8) DEFAULT NULL,
// 	`emoji` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`emojiU` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	`flag` tinyint(1) NOT NULL DEFAULT '1',
// 	`wikiDataId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Rapid API GeoDB Cities',
// 	PRIMARY KEY (`id`),
// 	KEY `country_continent` (`region_id`),
// 	KEY `country_subregion` (`subregion_id`),
// 	CONSTRAINT `country_continent_final` FOREIGN KEY (`region_id`) REFERENCES `regions` (`id`),
// 	CONSTRAINT `country_subregion_final` FOREIGN KEY (`subregion_id`) REFERENCES `subregions` (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

type CountryType struct {
	ID             uint      `gorm:"primarykey" json:"ID"`
	Name           string    `json:"name" gorm:"comment:'中文名称'"`
	ISO3           string    `json:"iso3" gorm:"comment:'iso3'"`
	NumericCode    string    `json:"numeric_code" gorm:"comment:'numeric_code'"`
	ISO2           string    `json:"iso2" gorm:"comment:'iso2'"`
	PhoneCode      string    `json:"phonecode" gorm:"comment:'phonecode'"`
	Capital        string    `json:"capital" gorm:"comment:'capital'"`
	Currency       string    `json:"currency" gorm:"comment:'currency'"`
	CurrencyName   string    `json:"currency_name" gorm:"comment:'currency_name'"`
	CurrencySymbol string    `json:"currency_symbol" gorm:"comment:'currency_symbol'"`
	TLD            string    `json:"tld" gorm:"comment:'tld'"`
	Native         string    `json:"native" gorm:"comment:'native'"`
	Region         string    `json:"region" gorm:"comment:'region'"`
	RegionID       uint      `json:"region_id" gorm:"comment:'region_id'"`
	SubRegion      string    `json:"subregion" gorm:"comment:'subregion'"`
	SubRegionID    uint      `json:"subregion_id" gorm:"comment:'subregion_id'"`
	Nationality    string    `json:"nationality" gorm:"comment:'nationality'"`
	TimeZones      string    `json:"timezones" gorm:"comment:'timezones'"`
	Translations   string    `json:"translations" gorm:"comment:'translations'"`
	Latitude       float64   `json:"latitude" gorm:"comment:'latitude'"`
	Longitude      float64   `json:"longitude" gorm:"comment:'longitude'"`
	Emoji          string    `json:"emoji" gorm:"comment:'emoji'"`
	EmojiU         string    `json:"emojiU" gorm:"comment:'emojiU'"`
	CreatedAt      time.Time `json:"created_at" gorm:"comment:'created_at'"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"comment:'updated_at'"`
	Flag           bool      `json:"flag" gorm:"comment:'flag'"`
	WikiDataID     string    `json:"wikiDataId" gorm:"comment:'wikiDataId'"`
}

func (CountryType) TableName() string {
	return "countries"
}

// CREATE TABLE `states` (
// 	`id` mediumint unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`country_id` mediumint unsigned NOT NULL,
// 	`country_code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`fips_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`iso2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`latitude` decimal(10,8) DEFAULT NULL,
// 	`longitude` decimal(11,8) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	`flag` tinyint(1) NOT NULL DEFAULT '1',
// 	`wikiDataId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Rapid API GeoDB Cities',
// 	PRIMARY KEY (`id`),
// 	KEY `country_region` (`country_id`),
// 	CONSTRAINT `country_region_final` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=5303 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

type StateType struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	Name        string    `json:"name" gorm:"comment:'中文名称'"`
	CountryID   uint      `json:"country_id" gorm:"comment:'country_id'"`
	CountryCode string    `json:"country_code" gorm:"comment:'country_code'"`
	FipsCode    string    `json:"fips_code" gorm:"comment:'fips_code'"`
	ISO2        string    `json:"iso2" gorm:"comment:'iso2'"`
	Type        string    `json:"type" gorm:"comment:'type'"`
	Latitude    float64   `json:"latitude" gorm:"comment:'latitude'"`
	Longitude   float64   `json:"longitude" gorm:"comment:'longitude'"`
	CreatedAt   time.Time `json:"created_at" gorm:"comment:'created_at'"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"comment:'updated_at'"`
	Flag        bool      `json:"flag" gorm:"comment:'flag'"`
	WikiDataID  string    `json:"wikiDataId" gorm:"comment:'wikiDataId'"`
}

func (StateType) TableName() string {
	return "states"
}

// CREATE TABLE `cities` (
// 	`id` mediumint unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`state_id` mediumint unsigned NOT NULL,
// 	`state_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`country_id` mediumint unsigned NOT NULL,
// 	`country_code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`latitude` decimal(10,8) NOT NULL,
// 	`longitude` decimal(11,8) NOT NULL,
// 	`created_at` timestamp NOT NULL DEFAULT '2014-01-01 12:01:01',
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	`flag` tinyint(1) NOT NULL DEFAULT '1',
// 	`wikiDataId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Rapid API GeoDB Cities',
// 	PRIMARY KEY (`id`),
// 	KEY `cities_test_ibfk_1` (`state_id`),
// 	KEY `cities_test_ibfk_2` (`country_id`),
// 	CONSTRAINT `cities_ibfk_1` FOREIGN KEY (`state_id`) REFERENCES `states` (`id`),
// 	CONSTRAINT `cities_ibfk_2` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=154082 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

type CityType struct {
	ID          uint      `gorm:"primarykey" json:"ID"`
	Name        string    `json:"name" gorm:"comment:'中文名称'"`
	StateID     uint      `json:"state_id" gorm:"comment:'state_id'"`
	StateCode   string    `json:"state_code" gorm:"comment:'state_code'"`
	CountryID   uint      `json:"country_id" gorm:"comment:'country_id'"`
	CountryCode string    `json:"country_code" gorm:"comment:'country_code'"`
	Latitude    float64   `json:"latitude" gorm:"comment:'latitude'"`
	Longitude   float64   `json:"longitude" gorm:"comment:'longitude'"`
	CreatedAt   time.Time `json:"created_at" gorm:"comment:'created_at'"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"comment:'updated_at'"`
	Flag        bool      `json:"flag" gorm:"comment:'flag'"`
	WikiDataID  string    `json:"wikiDataId" gorm:"comment:'wikiDataId'"`
}

func (CityType) TableName() string {
	return "cities"
}

// CREATE TABLE `regions` (
// 	`id` mediumint unsigned NOT NULL AUTO_INCREMENT,
// 	`name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
// 	`translations` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	`flag` tinyint(1) NOT NULL DEFAULT '1',
// 	`wikiDataId` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Rapid API GeoDB Cities',
// 	PRIMARY KEY (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

type RegionType struct {
	ID           uint      `gorm:"primarykey" json:"ID"`
	Name         string    `json:"name" gorm:"comment:'name'"`
	Translations string    `json:"translations" gorm:"comment:'translations'"`
	CreatedAt    time.Time `json:"created_at" gorm:"comment:'created_at'"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"comment:'updated_at'"`
	Flag         bool      `json:"flag" gorm:"comment:'flag'"`
	WikiDataID   string    `json:"wikiDataId" gorm:"comment:'wikiDataId'"`
}

func (RegionType) TableName() string {
	return "regions"
}
