// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

//! HTTP server: mounts `/api`, `/rest` (Subsonic) and the embedded web player.

use axum::Router;
use axum::http::StatusCode;
use axum::routing::get;

pub fn router() -> Router {
    Router::new().route("/healthz", get(health))
}

async fn health() -> StatusCode {
    StatusCode::NO_CONTENT
}
