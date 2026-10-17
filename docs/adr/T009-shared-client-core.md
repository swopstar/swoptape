# ADR-T009: Native clients share a Rust core

**Status:** proposed
**Date:** 2026-10-17

## Context

After the MVP, the plan is an iOS app, a desktop app and later a native Android app. All of them need offline downloads and playback, and the desktop app may later write to MP3 players and iPods. Implementing the API client, download queue and offline cache once per platform would be wasteful and would drift.

The [licensing standard](https://github.com/swopstar/meta/blob/main/standards/licensing.md) gives desktop clients GPLv3 and mobile clients a custom copyleft licence that is still to be drafted.

## Decision

- A `swoptape-client` crate holds the API client, sync logic, download queue and the local offline database.
- The iOS app is written in SwiftUI and uses `swoptape-client` through UniFFI's Swift bindings.
- The Android app comes later, written natively in Kotlin through UniFFI's Kotlin bindings. It should feel native to Android rather than being a port of the iOS app.
- The desktop app uses Tauri. It reuses the web player's UI and calls `swoptape-client` from Tauri's Rust side.
- Device sync (MP3 players, iPods) lives in the desktop app's Rust side. iPod support comes last because of its proprietary database format.

## Alternatives considered

- **Separate native implementations per platform.** Native UI either way, but the sync and offline logic would be written and maintained once per platform.
- **A cross-platform UI toolkit for mobile (Flutter, React Native).** One UI codebase for iOS and Android, but neither app would feel fully native.
- **Electron for desktop.** The suite's earlier assumption, before ADR-0001. Tauri is chosen because the client core is already Rust, and Tauri uses far less memory.

## Consequences

- The client crate must build for iOS and Android targets. Its dependencies must stay compatible with both.
- Desktop and web share most UI, which limits how native the desktop app feels. This is accepted in exchange for less code.
- `swoptape-client` is linked into both a GPLv3 desktop app and custom-licensed mobile apps, so its own licence must allow both. MIT, as a shared library under the licensing standard, is the likely answer. This must be settled before client work starts.
