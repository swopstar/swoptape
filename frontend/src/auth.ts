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

export function clearAuthTokens(): void {
  localStorage.removeItem("accessToken");
  localStorage.removeItem("refreshToken");
  localStorage.removeItem("username");
}
