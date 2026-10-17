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
- Library schema and migrations
- Artwork

### M2: Identity

- Users and sessions
- App passwords for Subsonic clients

### M3: Subsonic API

- Browsing in both folder and tag modes
- Streaming and cover art
- Search, playlists, stars and scrobbles

### M4: Web player

- Folder and tag browsing, and playback
- App password management
- **Settings → About this instance**: licence, version and source link, as the
  licensing standard requires before general use

## To review

Carried over from the Go plan and not yet scheduled:

- Granular permissions (`admin`, `upload`, `tag`)
- TOTP
- Uploads and the `storage` service
- `st` command-line tool (migrations, inspection, admin tasks)
- Background job system (previously planned via gokit)
- Defining "sufficient tag quality" for tag-based views, once a real library
  has been imported

## Future

- iOS app (SwiftUI) with offline downloads
- Desktop app (Tauri) with offline downloads
- Syncing to MP3 players, then iPods, from the desktop app
- Native Android app
- Shared Rust library extracted for swopcart
