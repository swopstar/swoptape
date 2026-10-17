// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

use tokio::net::TcpListener;

const ADDRESS: &str = "0.0.0.0:8080";

#[tokio::main]
async fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind(ADDRESS).await?;
    axum::serve(listener, swoptape_server::router()).await
}
