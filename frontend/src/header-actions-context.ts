// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { createContext, useContext, useEffect, type ReactNode } from "react";

export interface HeaderActionsContextValue {
  actions: ReactNode;
  setActions: (actions: ReactNode) => void;
}

export const HeaderActionsContext =
  createContext<HeaderActionsContextValue | null>(null);

export function useHeaderActions(): HeaderActionsContextValue {
  const ctx = useContext(HeaderActionsContext);
  if (!ctx)
    throw new Error(
      "useHeaderActions must be used within HeaderActionsProvider",
    );
  return ctx;
}

export function useSetHeaderActions(actions: ReactNode): void {
  const { setActions } = useHeaderActions();
  useEffect(() => {
    setActions(actions);
    return () => setActions(null);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
}
