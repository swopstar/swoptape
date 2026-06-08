// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import "context"

func (h *Handlers) Search(ctx context.Context, request SearchRequestObject) (SearchResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) SearchArtists(ctx context.Context, request SearchArtistsRequestObject) (SearchArtistsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) SearchAlbums(ctx context.Context, request SearchAlbumsRequestObject) (SearchAlbumsResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) SearchTracks(ctx context.Context, request SearchTracksRequestObject) (SearchTracksResponseObject, error) {
	return nil, errNotImplemented
}

func (h *Handlers) SearchAutocomplete(ctx context.Context, request SearchAutocompleteRequestObject) (SearchAutocompleteResponseObject, error) {
	return nil, errNotImplemented
}
