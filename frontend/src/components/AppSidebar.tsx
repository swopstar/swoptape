// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { Link, useNavigate } from "@tanstack/react-router";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
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
  useSidebar,
} from "@swopstar/react-ui";
import {
  ChevronsUpDown,
  DiscAlbum,
  Folder,
  House,
  LogOut,
  Monitor,
  Plus,
  Search,
  Settings,
  User,
  User2,
} from "lucide-react";
import { clearAuthTokens, getUsername } from "../auth";
import { useEndCurrentSession } from "../api/auth/auth";
import { useThemeMode } from "../theme-context";

const navItems = [
  { to: "/home", label: "Home", icon: House },
  { to: "/collections", label: "Collections", icon: Folder },
  { to: "/artists", label: "Artists", icon: User2 },
  { to: "/albums", label: "Albums", icon: DiscAlbum },
  { to: "/search", label: "Search", icon: Search },
] as const;

export function AppSidebar() {
  const navigate = useNavigate();
  const username = getUsername();
  const { isMobile } = useSidebar();
  const { mode, setMode } = useThemeMode();

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
        <div className="flex items-center gap-1 p-1">
          <img src="/favicon.svg" width={32} height={32} alt="" />
          <span
            style={{
              fontVariationSettings: "'wght' 800, 'YTLC' 540, 'wdth' 75",
              fontSize: "20px",
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
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <SidebarMenuButton>
                  <User size={16} />
                  {username || "Account"}
                  <ChevronsUpDown size={14} style={{ marginLeft: "auto" }} />
                </SidebarMenuButton>
              </DropdownMenuTrigger>
              <DropdownMenuContent
                side={isMobile ? "bottom" : "top"}
                align="start"
                style={{ minWidth: "220px" }}
              >
                <DropdownMenuGroup>
                  <DropdownMenuItem
                    onClick={() => navigate({ to: "/settings" })}
                  >
                    <User size={16} />
                    Account
                  </DropdownMenuItem>
                  <DropdownMenuItem
                    onClick={() => navigate({ to: "/settings" })}
                  >
                    <Settings size={16} />
                    Settings
                  </DropdownMenuItem>
                </DropdownMenuGroup>
                <DropdownMenuSeparator />
                <DropdownMenuSub>
                  <DropdownMenuSubTrigger>
                    <Monitor size={16} />
                    Theme
                  </DropdownMenuSubTrigger>
                  <DropdownMenuSubContent>
                    <DropdownMenuItem
                      onClick={() => setMode("light")}
                      data-active={mode === "light" || undefined}
                    >
                      Light
                    </DropdownMenuItem>
                    <DropdownMenuItem
                      onClick={() => setMode("dark")}
                      data-active={mode === "dark" || undefined}
                    >
                      Dark
                    </DropdownMenuItem>
                    <DropdownMenuItem
                      onClick={() => setMode("auto")}
                      data-active={mode === "auto" || undefined}
                    >
                      System
                    </DropdownMenuItem>
                  </DropdownMenuSubContent>
                </DropdownMenuSub>
                <DropdownMenuSeparator />
                <DropdownMenuItem
                  onClick={() => logout.mutate()}
                  disabled={logout.isPending}
                >
                  <LogOut size={16} />
                  Log out
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
}
