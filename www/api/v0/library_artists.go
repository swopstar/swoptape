// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "context"

func (h *Handlers) ListArtists(ctx context.Context, request ListArtistsRequestObject) (ListArtistsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) CreateArtist(ctx context.Context, request CreateArtistRequestObject) (CreateArtistResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) GetArtist(ctx context.Context, request GetArtistRequestObject) (GetArtistResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) UpdateArtist(ctx context.Context, request UpdateArtistRequestObject) (UpdateArtistResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) ListArtistAlbums(ctx context.Context, request ListArtistAlbumsRequestObject) (ListArtistAlbumsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) MergeArtist(ctx context.Context, request MergeArtistRequestObject) (MergeArtistResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) SplitArtist(ctx context.Context, request SplitArtistRequestObject) (SplitArtistResponseObject, error) {
	return nil, errNotImplemented
}
