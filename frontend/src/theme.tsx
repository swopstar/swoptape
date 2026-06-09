// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { useState } from "react";
import { ThemeProvider } from "@swopstar/react-ui";
import { ThemeModeContext, type ThemeMode } from "./theme-context";

const STORAGE_KEY = "swoptape:theme";

function readStoredMode(): ThemeMode {
  try {
    const v = localStorage.getItem(STORAGE_KEY);
    if (v === "light" || v === "dark" || v === "auto") return v;
  } catch {
    // ignore
  }
  return "auto";
}

export function ThemeModeProvider({ children }: { children: React.ReactNode }) {
  const [mode, setModeState] = useState<ThemeMode>(readStoredMode);

  const setMode = (next: ThemeMode) => {
    try {
      localStorage.setItem(STORAGE_KEY, next);
    } catch {
      // ignore
    }
    setModeState(next);
  };

  return (
    <ThemeModeContext.Provider value={{ mode, setMode }}>
      <ThemeProvider seedColor="#863bff" mode={mode}>
        {children}
      </ThemeProvider>
    </ThemeModeContext.Provider>
  );
}
