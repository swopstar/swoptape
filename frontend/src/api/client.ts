// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

const BASE_URL = "/api/v0";

export const apiClient = async <T>(
  url: string,
  options: RequestInit,
): Promise<T> => {
  const token = localStorage.getItem("accessToken");

  const response = await fetch(`${BASE_URL}${url}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options.headers,
    },
  });

  if (!response.ok) {
    throw await response.json();
  }

  const data = await response.json();
  return { data, status: response.status, headers: response.headers } as T;
};
