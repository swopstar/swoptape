# ADR-T006: Dual-mode library model

**Status:** proposed
**Date:** 2026-10-17

## Context

Most music servers force a choice. Either they flatten the library into an artist and album hierarchy built from tags and lose how the files are organised, or they expose raw folders with no tag context. swoptape's goal is to serve both ways of working with a library.

The [swoptape overview](https://github.com/swopstar/meta/blob/main/projects/swoptape/overview.md) in swopstar/meta describes swoptape as file-first, with tag-based views shown only when tag quality is good enough to trust. This ADR keeps folders as the primary structure and makes tag views first-class wherever the tags support them.

The Subsonic API already has both modes. `getIndexes` and `getMusicDirectory` browse folders, and `getArtists`, `getArtist` and `getAlbum` browse by ID3 tags. Most servers implement folder mode poorly.

## Decision

- A **library** is a configured root directory.
- **Folders** (shown to users as collections) form a tree (`parent_id`) under each library. They are first-class records, not derived from track paths at query time.
- Every **track** belongs to exactly one folder and keeps its path relative to the library root.
- **Artists** and **albums** are entities built from tags, linked to tracks many-to-many (multiple artists, compilations).
- Tag-derived entities can be corrected without editing files (for example, merging or splitting artists). These corrections survive rescans.
- When a track's file disappears, the scanner marks the track **missing** rather than deleting it. If a file with the same content fingerprint appears elsewhere, the track is moved to the new path and keeps its ID. Tracks still missing after a grace period are deleted.
- Both Subsonic modes are backed by these tables directly. Folder IDs and tag IDs are distinct and stable across rescans.
- How tag quality gates tag views (per library, per folder, or per track) is decided separately, once the scanner can measure it.

## Alternatives considered

- **Derive folders from track paths at query time.** Less to store, but makes folder browsing slow and gives folders no stable identity.
- **Tags only, with folders as a filter.** The model most servers use, and the one swoptape exists to avoid.
- **Delete tracks as soon as their files disappear.** Simplest, but moving or renaming a folder would lose plays, stars and playlist entries.
- **Soft-delete every table and run a periodic cleanup job.** The Go design's approach. It applies to every table, when the real need is specific to tracks that go missing during scans.

## Consequences

- Scanning must preserve folder structure and tags, and reconcile both incrementally.
- The web player and the native clients can offer a folder view and a tag view of the same tracks, and move between them (for example, from an album to the folder it lives in).
- Stable IDs across rescans need a deliberate identity strategy: a path plus a content fingerprint, rather than IDs that change on every scan.
- Missing tracks are hidden from browsing, search and the Subsonic API, but their plays, stars and playlist entries are kept until the grace period ends. The grace period is configurable, so a library on a drive that is temporarily unmounted isn't wiped by the next scan. A scan that finds a library's root directory missing or unreadable stops with an error instead of marking its tracks missing.
- Scanning runs as a single background task with bounded concurrency and reports its progress. It does not need a general job system.
