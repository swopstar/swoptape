// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { createContext, useContext } from "react";

export type ThemeMode = "light" | "dark" | "auto";

export interface ThemeModeContextValue {
  mode: ThemeMode;
  setMode: (mode: ThemeMode) => void;
}

export const ThemeModeContext = createContext<ThemeModeContextValue | null>(
  null,
);

export function useThemeMode(): ThemeModeContextValue {
  const ctx = useContext(ThemeModeContext);
  if (!ctx)
    throw new Error("useThemeMode must be used within ThemeModeProvider");
  return ctx;
}
