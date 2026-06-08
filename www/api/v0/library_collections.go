// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "context"

func (h *Handlers) ListCollections(ctx context.Context, request ListCollectionsRequestObject) (ListCollectionsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) GetCollection(ctx context.Context, request GetCollectionRequestObject) (GetCollectionResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) ListSubcollections(ctx context.Context, request ListSubcollectionsRequestObject) (ListSubcollectionsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) ListCollectionTracks(ctx context.Context, request ListCollectionTracksRequestObject) (ListCollectionTracksResponseObject, error) {
	return nil, errNotImplemented
}
