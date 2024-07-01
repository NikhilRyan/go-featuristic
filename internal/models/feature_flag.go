package models

type FeatureFlag struct {
	ID        uint   `gorm:"primaryKey"`
	Namespace string `gorm:"index;not null"`
	Key       string `gorm:"index;not null"`
	Value     string `gorm:"type:text"`
	Type      string `gorm:"type:varchar(50);not null"`
}
