package models

type FeatureFlag struct {
	ID           uint   `gorm:"primaryKey"`
	Namespace    string `gorm:"index;not null"`
	Key          string `gorm:"index;not null"`
	Value        string `gorm:"type:text"`
	Type         string `gorm:"type:varchar(50);not null"`
	ABTestValue  string `gorm:"type:text"`
	ABTestType   string `gorm:"type:varchar(50)"`
	TargetGroup  string `gorm:"type:varchar(100)"`
	TargetGroupB string `gorm:"type:varchar(100)"`
}
