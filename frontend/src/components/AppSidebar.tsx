// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { Link, useNavigate } from "@tanstack/react-router";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarSeparator,
} from "@swopstar/react-ui";
import {
  ChevronsUpDown,
  Disc3,
  Download,
  FolderOpen,
  House,
  Plus,
  Search,
  Smartphone,
  User,
} from "lucide-react";
import { clearAuthTokens, getUsername } from "../auth";
import { useEndCurrentSession } from "../api/auth/auth";

const navItems = [
  { to: "/home", label: "Home", icon: House },
  { to: "/collections", label: "Collections", icon: FolderOpen },
  { to: "/artists", label: "Artists", icon: User },
  { to: "/albums", label: "Albums", icon: Disc3 },
  { to: "/search", label: "Search", icon: Search },
] as const;

export function AppSidebar() {
  const navigate = useNavigate();
  const username = getUsername();

  const logout = useEndCurrentSession({
    mutation: {
      onSettled: () => {
        clearAuthTokens();
        navigate({ to: "/login" });
      },
    },
  });

  return (
    <Sidebar>
      <SidebarHeader>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "8px",
            padding: "4px 0",
          }}
        >
          <img src="/favicon.svg" width={28} height={26} alt="" />
          <span
            style={{
              fontVariationSettings: "'wght' 800, 'YTLC' 540",
              fontSize: "16px",
            }}
          >
            swoptape
          </span>
        </div>
      </SidebarHeader>

      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu>
              {navItems.map(({ to, label, icon: Icon }) => (
                <SidebarMenuItem key={to}>
                  <SidebarMenuButton asChild>
                    <Link
                      to={to}
                      activeOptions={{ exact: true }}
                      activeProps={{ "data-active": "true" }}
                    >
                      <Icon size={16} />
                      {label}
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupLabel>Playlists</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              <SidebarMenuItem>
                <SidebarMenuButton>
                  <Plus size={16} />
                  New playlist
                </SidebarMenuButton>
              </SidebarMenuItem>
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupLabel>Recents</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu />
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>

      <SidebarSeparator />

      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton>
              <Download size={16} />
              swoptape app
            </SidebarMenuButton>
          </SidebarMenuItem>
          <SidebarMenuItem>
            <SidebarMenuButton>
              <Smartphone size={16} />
              Listen on the go
            </SidebarMenuButton>
          </SidebarMenuItem>
          <SidebarMenuItem>
            <SidebarMenuButton
              onClick={() => logout.mutate()}
              disabled={logout.isPending}
            >
              <User size={16} />
              {username || "Account"}
              <ChevronsUpDown size={14} style={{ marginLeft: "auto" }} />
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
}
