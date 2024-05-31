package bucketv1

import (
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
)

func ToMessage(b *model.Bucket) *Bucket {
	return &Bucket{
		Id:   b.ID.String(),
		Name: b.Name,
		RootDirectoryId: func() string {
			if b.RootDirectoryID == nil {
				return ""
			}
			return b.RootDirectoryID.String()
		}(),
	}
}

func ToMessages(buckets []*model.Bucket) []*Bucket {
	messages := make([]*Bucket, len(buckets))
	for i, bucket := range buckets {
		messages[i] = ToMessage(bucket)
	}
	return messages
}
