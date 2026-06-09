// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { Progress } from "@swopstar/react-ui";
import { ChevronRight, FileText, Github } from "lucide-react";
import {
  useGetInstanceStatus,
  useGetInstanceVersion,
} from "../api/system/system";

const MONTHS = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

function formatBytes(bytes: number): string {
  if (bytes >= 1024 ** 3) return `${(bytes / 1024 ** 3).toFixed(1)} GB`;
  if (bytes >= 1024 ** 2) return `${(bytes / 1024 ** 2).toFixed(1)} MB`;
  if (bytes >= 1024) return `${(bytes / 1024).toFixed(1)} KB`;
  return `${bytes} B`;
}

interface InfoRowProps {
  icon: React.ReactNode;
  title: string;
  subtitle: string;
  onClick?: () => void;
}

function InfoRow({ icon, title, subtitle, onClick }: InfoRowProps) {
  return (
    <button
      onClick={onClick}
      className="flex w-full items-center gap-4 rounded-lg bg-muted/50 px-4 py-3 text-left transition-colors hover:bg-muted"
    >
      <span className="shrink-0 text-muted-foreground">{icon}</span>
      <span className="flex-1">
        <span className="block text-sm font-semibold">{title}</span>
        <span className="block text-xs text-muted-foreground">{subtitle}</span>
      </span>
      <ChevronRight size={16} className="shrink-0 text-muted-foreground" />
    </button>
  );
}

interface StatusItemProps {
  label: string;
  used: number;
  total: number;
}

function StatusItem({ label, used, total }: StatusItemProps) {
  const pct = total > 0 ? Math.round((used / total) * 100) : 0;
  return (
    <div className="flex flex-col gap-2 rounded-lg bg-muted/50 px-4 py-3">
      <div className="flex items-center justify-between text-sm">
        <span className="font-semibold">{label}</span>
        <span className="text-xs text-muted-foreground">
          {formatBytes(used)} / {formatBytes(total)}
        </span>
      </div>
      <Progress value={pct} />
    </div>
  );
}

export function SystemInfoPage() {
  const { data: versionData } = useGetInstanceVersion();
  const { data: statusData } = useGetInstanceStatus();

  const build = versionData?.data?.version;
  const p = versionData?.data?.parsed;
  const label =
    p?.year != null && p?.month != null && p?.release != null
      ? `${MONTHS[(p.month ?? 1) - 1]} ${2000 + p.year} release ${p.release}`
      : (build ?? "—");
  const source = versionData?.data?.source;

  const mem = statusData?.data?.memory;
  const storage = statusData?.data?.storage;

  return (
    <div className="mx-auto flex w-full max-w-2xl flex-col gap-8 p-6">
      <div className="flex flex-col items-center gap-3 rounded-xl bg-muted/50 px-6 py-8">
        <img src="/favicon.svg" alt="swoptape" className="h-16 w-16" />
        <div className="text-center">
          <p className="text-xl font-bold">swoptape</p>
          <p className="text-sm text-muted-foreground">
            {label}
            {build && <span className="ml-1 font-mono text-xs">({build})</span>}
          </p>
        </div>
      </div>

      <div className="flex flex-col gap-2">
        <InfoRow
          icon={<FileText size={20} />}
          title="Software Licence"
          subtitle="GNU Affero General Public License 3.0 only, and third party dependencies"
          onClick={() =>
            window.open(
              "https://spdx.org/licenses/AGPL-3.0-only.html",
              "_blank",
              "noopener,noreferrer",
            )
          }
        />
        <InfoRow
          icon={<Github size={20} />}
          title="Source code"
          subtitle="View the source code"
          onClick={() =>
            source && window.open(source, "_blank", "noopener,noreferrer")
          }
        />
      </div>

      <div className="flex flex-col gap-2">
        <h2 className="px-1 text-lg font-bold">Status</h2>
        {mem && mem.usedBytes != null && mem.totalBytes != null && (
          <StatusItem
            label="Memory"
            used={mem.usedBytes}
            total={mem.totalBytes}
          />
        )}
        {storage && storage.usedBytes != null && storage.totalBytes != null && (
          <StatusItem
            label="Disk"
            used={storage.usedBytes}
            total={storage.totalBytes}
          />
        )}
      </div>
    </div>
  );
}
