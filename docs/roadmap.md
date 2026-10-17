# swoptape: Roadmap

## Done

- Go prototype (on `main`, kept as reference): config loading, database layer,
  users with granular permissions and TOTP, sessions

## Rust rewrite (MVP)

Developed on the `rewrite/rust` branch, which replaces `main` once the MVP is
reached.

### M0: Foundations

- Cargo workspace, CI, agent instructions and ADRs
- Build-time application version and source link, per the
  [versioning](https://github.com/swopstar/meta/blob/main/standards/versioning.md) and
  [licensing](https://github.com/swopstar/meta/blob/main/standards/licensing.md) standards

### M1: Library

- Scanning that preserves both folder structure (collections) and tags
- Scans run as a single background task with bounded concurrency, publishing
  progress for the web player to show. No general job system.
- Tracks whose files disappear are marked missing and kept for a grace period,
  so moved files keep their plays, stars and playlist entries
  ([ADR-T006](adr/T006-dual-mode-library.md))
- Library schema and migrations
- Artwork, with generated thumbnails in a cache directory under `data/`

### M2: Identity

- Users and sessions
- First-user setup through the API
- Permissions (`admin`, `upload`, `tag`) stored from the start; only `admin` is
  enforced until uploads and tag editing exist. Mapped to Subsonic roles where
  they have an equivalent.
- App passwords for Subsonic clients
- `swoptape` subcommands: `serve`, `migrate` and `user reset-password` for
  admin recovery

### M3: Subsonic API

- Browsing in both folder and tag modes
- Streaming and cover art
- Search, playlists, stars and scrobbles

### M4: Web player

- First-run setup screens
- Folder and tag browsing, and playback
- Empty states from the mockups, with "Upload tracks" hidden until uploads exist
- App password management
- **Settings → About this instance**: licence, version and source link, as the
  licensing standard requires before general use

## Open questions

- Defining "sufficient tag quality" for tag-based views, once a real library
  has been imported

## After the MVP

- Uploads: where uploaded files land, duplicate handling and size limits. This
  enables the `upload` permission.
- Tag editing (merging and splitting artists, corrections). This enables the
  `tag` permission.
- Second factor for web sign-in: TOTP, passkeys or both
- A general background job system, if more job types appear (transcode caching,
  artwork fetching)

## Not planned

- A separate `st` CLI binary. Admin tasks are `swoptape` subcommands instead.
- Soft deletes with a cleanup job. Missing tracks are handled by the scanner's
  grace period instead.

## Future

- iOS app (SwiftUI) with offline downloads
- Desktop app (Tauri) with offline downloads
- Syncing to MP3 players, then iPods, from the desktop app
- Native Android app
- Shared Rust library extracted for swopcart
