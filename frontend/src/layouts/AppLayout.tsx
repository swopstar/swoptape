// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { Outlet, useMatches, useNavigate } from "@tanstack/react-router";
import { Button, PageHeader, Scaffold } from "@swopstar/react-ui";
import { ArrowLeft } from "lucide-react";
import { AppSidebar } from "../components/AppSidebar";
import { HeaderActionsProvider } from "../header-actions";
import { useHeaderActions } from "../header-actions-context";

function AppLayoutInner() {
  const navigate = useNavigate();
  const matches = useMatches();
  const last = matches[matches.length - 1]?.staticData;
  const title = last?.title ?? "";
  const backTo = last?.backTo;
  const { actions } = useHeaderActions();

  return (
    <Scaffold
      sidebar={<AppSidebar />}
      header={
        <PageHeader
          title={title}
          prefixActions={
            backTo ? (
              <Button
                size="icon"
                variant="ghost"
                onClick={() => navigate({ to: backTo })}
              >
                <ArrowLeft size={18} />
              </Button>
            ) : undefined
          }
          actions={actions}
        />
      }
    >
      <Outlet />
    </Scaffold>
  );
}

export function AppLayout() {
  return (
    <HeaderActionsProvider>
      <AppLayoutInner />
    </HeaderActionsProvider>
  );
}
