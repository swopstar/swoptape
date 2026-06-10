// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import (
	"context"
	"encoding/json"

	"github.com/swopstar/swoptape/internal/notices"
)

type noticesData struct {
	App noticeEntry   `json:"app"`
	Go  []noticeEntry `json:"go"`
	Npm []noticeEntry `json:"npm"`
}

type noticeEntry struct {
	Module  string `json:"module"`
	Version string `json:"version"`
	Spdx    string `json:"spdx"`
	Text    string `json:"text"`
}

func (h *Handlers) GetInstanceNotices(_ context.Context, _ GetInstanceNoticesRequestObject) (GetInstanceNoticesResponseObject, error) {
	var data noticesData
	if err := json.Unmarshal(notices.Data, &data); err != nil {
		return nil, err
	}

	toEntries := func(src []noticeEntry) *[]NoticeEntry {
		out := make([]NoticeEntry, len(src))
		for i, e := range src {
			e := e
			out[i] = NoticeEntry{
				Module:  &e.Module,
				Version: &e.Version,
				Spdx:    &e.Spdx,
				Text:    &e.Text,
			}
		}
		return &out
	}

	return GetInstanceNotices200JSONResponse{
		App: &NoticeEntry{Spdx: &data.App.Spdx, Text: &data.App.Text},
		Go:  toEntries(data.Go),
		Npm: toEntries(data.Npm),
	}, nil
}
