// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import type { NoticeEntry } from "../api/schemas/noticeEntry";
import { useGetInstanceNotices } from "../api/system/system";

function Badge({
  text,
  variant = "muted",
}: {
  text: string;
  variant?: "muted" | "primary";
}) {
  const cls =
    variant === "primary"
      ? "rounded bg-primary/10 px-2 py-0.5 font-mono text-xs text-primary"
      : "rounded bg-muted px-2 py-0.5 font-mono text-xs text-muted-foreground";
  return <span className={cls}>{text}</span>;
}

function Entry({ entry }: { entry: NoticeEntry }) {
  const hasContent = entry.text && entry.text.trim().length > 0;
  const name = entry.module ?? "";
  const badges = (
    <span className="flex shrink-0 items-center gap-1.5">
      {entry.version && <Badge text={entry.version} variant="muted" />}
      {entry.spdx && <Badge text={entry.spdx} variant="primary" />}
    </span>
  );

  if (!hasContent) {
    return (
      <div className="flex items-center justify-between gap-3 rounded-lg border border-border bg-muted/30 px-4 py-3">
        <span className="flex-1 truncate font-mono text-sm">{name}</span>
        {badges}
      </div>
    );
  }

  return (
    <details className="group rounded-lg border border-border bg-muted/30">
      <summary className="flex cursor-pointer list-none items-center justify-between gap-3 px-4 py-3 hover:bg-muted/50">
        <span className="flex-1 truncate font-mono text-sm">{name}</span>
        <span className="flex items-center gap-1.5">
          {badges}
          <svg
            className="h-4 w-4 shrink-0 text-muted-foreground transition-transform group-open:rotate-180"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </span>
      </summary>
      <pre className="overflow-x-auto whitespace-pre-wrap break-words border-t border-border px-4 py-3 font-mono text-xs text-muted-foreground">
        {entry.text}
      </pre>
    </details>
  );
}

function AppLicenseSection({ entry }: { entry: NoticeEntry }) {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="text-2xl font-bold">swoptape</h1>
      <details className="group rounded-lg border border-border bg-muted/30">
        <summary className="flex cursor-pointer list-none items-center justify-between gap-3 px-4 py-3 hover:bg-muted/50">
          <span className="flex-1 text-sm font-medium">
            GNU Affero General Public License, version 3.0 only
          </span>
          <span className="flex items-center gap-1.5">
            {entry.spdx && <Badge text={entry.spdx} variant="primary" />}
            <svg
              className="h-4 w-4 shrink-0 text-muted-foreground transition-transform group-open:rotate-180"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M19 9l-7 7-7-7"
              />
            </svg>
          </span>
        </summary>
        <pre className="overflow-x-auto whitespace-pre-wrap break-words border-t border-border px-4 py-3 font-mono text-xs text-muted-foreground">
          {entry.text}
        </pre>
      </details>
    </div>
  );
}

function DepsGroup({
  title,
  entries,
}: {
  title: string;
  entries: NoticeEntry[];
}) {
  if (!entries.length) return null;
  return (
    <div className="flex flex-col gap-2">
      <h2 className="px-1 text-lg font-semibold">{title}</h2>
      {entries.map((e, i) => (
        <Entry key={i} entry={e} />
      ))}
    </div>
  );
}

export function LicensesPage() {
  const { data } = useGetInstanceNotices();
  const appEntry = data?.data?.app;
  const goEntries = data?.data?.go ?? [];
  const npmEntries = data?.data?.npm ?? [];

  return (
    <div className="mx-auto flex w-full max-w-2xl flex-col gap-10 p-6">
      {appEntry && <AppLicenseSection entry={appEntry} />}

      {(goEntries.length > 0 || npmEntries.length > 0) && (
        <div className="flex flex-col gap-6">
          <h1 className="text-2xl font-bold">Third-party dependencies</h1>
          <DepsGroup title="Go modules" entries={goEntries} />
          <DepsGroup title="Node modules" entries={npmEntries} />
        </div>
      )}
    </div>
  );
}
