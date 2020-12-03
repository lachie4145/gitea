// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"testing"

	"code.gitea.io/gitea/models"
	"github.com/stretchr/testify/assert"
)

func TestCombineLabelComments(t *testing.T) {
	var kases = []struct {
		beforeCombined []*models.Comment
		afterCombined  []*models.Comment
	}{
		{
			beforeCombined: []*models.Comment{
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "1",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
			afterCombined: []*models.Comment{
				{
					Type:        models.CommentTypeLabel,
					PosterID:    1,
					Content:     "1",
					CreatedUnix: 0,
					AddedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					RemovedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
		},
		{
			beforeCombined: []*models.Comment{
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "1",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 70,
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
			afterCombined: []*models.Comment{
				{
					Type:        models.CommentTypeLabel,
					PosterID:    1,
					Content:     "1",
					CreatedUnix: 0,
					AddedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
				{
					Type:        models.CommentTypeLabel,
					PosterID:    1,
					Content:     "",
					CreatedUnix: 70,
					RemovedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
		},
		{
			beforeCombined: []*models.Comment{
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "1",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:     models.CommentTypeLabel,
					PosterID: 2,
					Content:  "",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
			afterCombined: []*models.Comment{
				{
					Type:        models.CommentTypeLabel,
					PosterID:    1,
					Content:     "1",
					CreatedUnix: 0,
					AddedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
				{
					Type:        models.CommentTypeLabel,
					PosterID:    2,
					Content:     "",
					CreatedUnix: 0,
					RemovedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
				{
					Type:        models.CommentTypeComment,
					PosterID:    1,
					Content:     "test",
					CreatedUnix: 0,
				},
			},
		},
		{
			beforeCombined: []*models.Comment{
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "1",
					Label: &models.Label{
						Name: "kind/bug",
					},
					CreatedUnix: 0,
				},
				{
					Type:     models.CommentTypeLabel,
					PosterID: 1,
					Content:  "1",
					Label: &models.Label{
						Name: "kind/backport",
					},
					CreatedUnix: 10,
				},
			},
			afterCombined: []*models.Comment{
				{
					Type:        models.CommentTypeLabel,
					PosterID:    1,
					Content:     "1",
					CreatedUnix: 10,
					AddedLabels: []*models.Label{
						{
							Name: "kind/bug",
						},
						{
							Name: "kind/backport",
						},
					},
					Label: &models.Label{
						Name: "kind/bug",
					},
				},
			},
		},
	}

	for _, kase := range kases {
		var issue = models.Issue{
			Comments: kase.beforeCombined,
		}
		combineLabelComments(&issue)
		assert.EqualValues(t, kase.afterCombined, issue.Comments)
	}
}