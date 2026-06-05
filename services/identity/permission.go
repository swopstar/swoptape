// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package identity

type Permission string

var (
	PermAdmin  Permission = "admin"
	PermTag    Permission = "tag"
	PermUpload Permission = "upload"

	Permissions = map[Permission]bool{
		PermAdmin:  true,
		PermTag:    true,
		PermUpload: true,
	}
)
