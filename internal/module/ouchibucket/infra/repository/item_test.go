package repository

import (
	"testing"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestMakeItemRanegs(t *testing.T) {
	t.Parallel()
	type input struct {
		dir    *model.Directory
		workID uuid.UUID
	}
	tests := []struct {
		name  string
		input input
		want  []*db.Item
	}{
		{
			name: "Success",
			input: input{
				dir: &model.Directory{
					ID:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name: "dir01",
					Directorires: []*model.Directory{
						{
							ID:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
							Name: "dir02",
							Directorires: []*model.Directory{
								{
									ID:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
									Name: "dir03",
								},
							},
						},
						{
							ID:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
							Name: "dir04",
							Files: []*model.File{
								{
									ID:   uuid.MustParse("00000000-0000-0000-0000-000000000005"),
									Name: "file01.jpg",
								},
								{
									ID:   uuid.MustParse("00000000-0000-0000-0000-000000000006"),
									Name: "file02.png",
								},
							},
						},
					},
					Files: []*model.File{
						{
							ID:   uuid.MustParse("00000000-0000-0000-0000-000000000007"),
							Name: "file03.mp4",
						},
						{
							ID:   uuid.MustParse("00000000-0000-0000-0000-000000000008"),
							Name: "file04.mp3",
						},
					},
				},
				workID: uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			},
			want: []*db.Item{
				{
					ID:       "00000000-0000-0000-0000-000000000003",
					Name:     "dir03",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      3,
					Rgt:      4,
					IsFile:   false,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000002",
					Name:     "dir02",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      2,
					Rgt:      5,
					IsFile:   false,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000005",
					Name:     "file01.jpg",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      7,
					Rgt:      8,
					IsFile:   true,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000006",
					Name:     "file02.png",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      9,
					Rgt:      10,
					IsFile:   true,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000004",
					Name:     "dir04",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      6,
					Rgt:      11,
					IsFile:   false,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000007",
					Name:     "file03.mp4",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      12,
					Rgt:      13,
					IsFile:   true,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000008",
					Name:     "file04.mp3",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      14,
					Rgt:      15,
					IsFile:   true,
				},
				{
					ID:       "00000000-0000-0000-0000-000000000001",
					Name:     "dir01",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      1,
					Rgt:      16,
					IsFile:   false,
				},
			},
		},
		{
			name: "Success - Empty dir",
			input: input{
				dir: &model.Directory{
					ID:   uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name: "dir01",
				},
				workID: uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			},
			want: []*db.Item{
				{
					ID:       "00000000-0000-0000-0000-000000000001",
					Name:     "dir01",
					BucketID: "10000000-0000-0000-0000-000000000001",
					Lft:      1,
					Rgt:      2,
					IsFile:   false,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := []*db.Item{}
			count := 0
			makeItemRanges(1, tt.input.dir, tt.input.workID, &got, &count)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
		})
	}
}
