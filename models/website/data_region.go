package website

type DataRegion struct {
	Id          int    `json:"id" gorm:"column:id;comment:主键"`
	Pid         int    `json:"pid" gorm:"column:pid;comment:父ID"`
	Path        string `json:"path" gorm:"column:path;comment:路径"`
	Level       int    `json:"level" gorm:"column:level;comment:层级"`
	Name        string `json:"name" gorm:"column:name;comment:中文名称"`
	Name_en     string `json:"name_en" gorm:"column:name_en;comment:英文名称"`
	Name_pinyin string `json:"name_pinyin" gorm:"column:name_pinyin;comment:中文拼音"`
	Code        string `json:"code" gorm:"column:code;comment:代码"`
}

func (DataRegion) TableName() string {
	return "data_region"
}
