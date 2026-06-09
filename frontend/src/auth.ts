// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

export function getAccessToken(): string | null {
  return localStorage.getItem("accessToken");
}

export function isAuthenticated(): boolean {
  return !!localStorage.getItem("accessToken");
}

export function getUsername(): string {
  return localStorage.getItem("username") ?? "";
}

export function setAuthTokens(
  accessToken: string,
  refreshToken: string,
  username?: string,
): void {
  localStorage.setItem("accessToken", accessToken);
  localStorage.setItem("refreshToken", refreshToken);
  if (username) localStorage.setItem("username", username);
}

export function setEntitlements(entitlements: string[]): void {
  localStorage.setItem("entitlements", JSON.stringify(entitlements));
}

export function getEntitlements(): string[] {
  try {
    const raw = localStorage.getItem("entitlements");
    if (!raw) return [];
    const parsed = JSON.parse(raw);
    return Array.isArray(parsed) ? parsed : [];
  } catch {
    return [];
  }
}

export function hasEntitlement(e: string): boolean {
  return getEntitlements().includes(e);
}

export function clearAuthTokens(): void {
  localStorage.removeItem("accessToken");
  localStorage.removeItem("refreshToken");
  localStorage.removeItem("username");
  localStorage.removeItem("entitlements");
}
