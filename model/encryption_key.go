package model

type EncryptionKey struct {
	BaseModel  BaseModel `gorm:"embedded"`
	ContentID  string    `gorm:"uniqueIndex:idx_content_package_quality;size:12;not null"`
	PackageID  string    `gorm:"uniqueIndex:idx_content_package_quality;size:16;not null"`
	Quality    string    `gorm:"uniqueIndex:idx_content_package_quality;not null"`
	ProviderID string    `gorm:"size:8"`
	DrmScheme  string
	KeyID      string `gorm:"size:32"`
	KeyIV      string `gorm:"size:32"`
	Key        string `gorm:"size:64"`
}
