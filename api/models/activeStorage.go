package models

import (
	"time"
)

type activeStorageAttachments struct {
	name       string    `gorm:"not null;index:indexActiveStorageAttachmentsUniqueness,unique"`
	recordType string    `gorm:"not null;index:indexActiveStorageAttachmentsUniqueness,unique"`
	recordId   int       `gorm:"not null;index:indexActiveStorageAttachmentsUniqueness,unique"`
	blobId     int       `gorm:"not null;index:indexActiveStorageAttachmentsOnBlobId;index:indexActiveStorageAttachmentsUniqueness,unique;foreignKey:blobId"`
	createdAt  time.Time `gorm:"not null"`
}

type activeStorageBlobs struct {
	key         string `gorm:"not null;index:indexActiveStorageBlobsOnKey,unique"`
	filename    string `gorm:"not null"`
	contentType string
	metadata    string
	serviceName string `gorm:"not null"`
	byteSize    int    `gorm:"not null"`
	checksum    string
	createdAt   time.Time `gorm:"not null"`
}

type activeStorageVariantRecords struct {
	blobId          int    `gorm:"not null;index:indexActiveStorageVariantRecordsUniqueness,unique;foreignKey:blobId"`
	variationDigest string `gorm:"not null;index:indexActiveStorageVariantRecordsUniqueness,unique"`
}
