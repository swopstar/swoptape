# SPDX-License-Identifier: AGPL-3.0-only
# SPDX-FileCopyrightText: 2026 Rareș Nistor

version_year    := `date +%y`
version_month   := `date +%m`
version_release := "0"
version_patch   := "0"
version_pre     := "dev"
version_branch  := `git rev-parse --abbrev-ref HEAD`
version_commit  := `git rev-parse --short HEAD`
source_url      := ""
