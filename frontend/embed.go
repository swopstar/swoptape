// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package frontend

import "embed"

//go:generate npm run build
//go:embed all:dist
var Content embed.FS
