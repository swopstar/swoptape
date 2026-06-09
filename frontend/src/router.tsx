// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import {
  createRootRoute,
  createRoute,
  createRouter,
  Outlet,
  redirect,
} from "@tanstack/react-router";
import { isAuthenticated } from "./auth";
import { AppLayout } from "./layouts/AppLayout";
import { LoginPage } from "./pages/LoginPage";
import { HomePage } from "./pages/HomePage";
import { CollectionsPage } from "./pages/CollectionsPage";
import { ArtistsPage } from "./pages/ArtistsPage";
import { AlbumsPage } from "./pages/AlbumsPage";
import { SearchPage } from "./pages/SearchPage";
import { SettingsPage } from "./pages/SettingsPage";

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
  interface StaticDataRouteOption {
    title?: string;
  }
}

const rootRoute = createRootRoute({ component: Outlet });

const loginRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/login",
  beforeLoad: () => {
    if (isAuthenticated()) throw redirect({ to: "/home" });
  },
  component: LoginPage,
});

const authLayoutRoute = createRoute({
  getParentRoute: () => rootRoute,
  id: "_auth",
  beforeLoad: ({ location }) => {
    if (!isAuthenticated()) {
      throw redirect({ to: "/login", search: { redirect: location.href } });
    }
  },
  component: AppLayout,
});

const indexRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/",
  beforeLoad: () => {
    throw redirect({ to: "/home" });
  },
});

const homeRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/home",
  staticData: { title: "Home" },
  component: HomePage,
});

const collectionsRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/collections",
  staticData: { title: "Collections" },
  component: CollectionsPage,
});

const artistsRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/artists",
  staticData: { title: "Artists" },
  component: ArtistsPage,
});

const albumsRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/albums",
  staticData: { title: "Albums" },
  component: AlbumsPage,
});

const searchRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/search",
  staticData: { title: "Search" },
  component: SearchPage,
});

const settingsRoute = createRoute({
  getParentRoute: () => authLayoutRoute,
  path: "/settings",
  staticData: { title: "Settings" },
  component: SettingsPage,
});

const routeTree = rootRoute.addChildren([
  loginRoute,
  authLayoutRoute.addChildren([
    indexRoute,
    homeRoute,
    collectionsRoute,
    artistsRoute,
    albumsRoute,
    searchRoute,
    settingsRoute,
  ]),
]);

export const router = createRouter({ routeTree });
