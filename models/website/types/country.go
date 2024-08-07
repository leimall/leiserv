package website

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
