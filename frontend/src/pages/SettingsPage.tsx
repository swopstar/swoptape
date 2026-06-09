// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { useState, type ReactNode } from "react";
import { useNavigate } from "@tanstack/react-router";
import {
  Badge,
  Button,
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  Input,
  Label,
} from "@swopstar/react-ui";
import {
  CheckCircle,
  ChevronRight,
  CircleGauge,
  Clock,
  KeyRound,
  Library,
  MonitorCog,
  Music,
  ShieldCheck,
  Smartphone,
  Timer,
  Users,
} from "lucide-react";
import { hasEntitlement } from "../auth";
import { useGetInstanceVersion } from "../api/system/system";

interface SettingsRowProps {
  icon: ReactNode;
  title: string;
  subtitle?: string;
  onClick: () => void;
}

function SettingsRow({ icon, title, subtitle, onClick }: SettingsRowProps) {
  return (
    <button
      onClick={onClick}
      className="flex w-full items-center gap-4 rounded-lg bg-muted/50 px-4 py-3 text-left transition-colors hover:bg-muted"
    >
      <span className="shrink-0 text-muted-foreground">{icon}</span>
      <span className="flex-1">
        <span className="block text-sm font-semibold">{title}</span>
        {subtitle && (
          <span className="block text-xs text-muted-foreground">
            {subtitle}
          </span>
        )}
      </span>
      <ChevronRight size={16} className="shrink-0 text-muted-foreground" />
    </button>
  );
}

function Section({ label, children }: { label?: string; children: ReactNode }) {
  return (
    <div className="flex flex-col gap-2">
      {label && <h2 className="px-1 text-lg font-bold">{label}</h2>}
      {children}
    </div>
  );
}

export function SettingsPage() {
  const navigate = useNavigate();
  const isAdmin = hasEntitlement("admin");

  const [changePasswordOpen, setChangePasswordOpen] = useState(false);
  const [mfaOpen, setMfaOpen] = useState(false);

  const { data: versionData } = useGetInstanceVersion();
  const versionSubtitle = versionData?.data?.version
    ? `swoptape ${versionData.data.version}`
    : undefined;

  return (
    <div className="mx-auto flex w-full max-w-2xl flex-col gap-8 p-6">
      <Section>
        <SettingsRow
          icon={<KeyRound size={20} />}
          title="Change password"
          onClick={() => setChangePasswordOpen(true)}
        />
        <SettingsRow
          icon={<ShieldCheck size={20} />}
          title="Set up multi-factor authentication"
          onClick={() => setMfaOpen(true)}
        />
        <SettingsRow
          icon={<Smartphone size={20} />}
          title="Active sessions"
          onClick={() => navigate({ to: "/settings/sessions" })}
        />
        <SettingsRow
          icon={<CircleGauge size={20} />}
          title="Application passwords"
          subtitle="Log into less secure applications"
          onClick={() => navigate({ to: "/settings/tokens" })}
        />
      </Section>

      {isAdmin && (
        <Section label="Administration">
          <SettingsRow
            icon={<Users size={20} />}
            title="Users"
            onClick={() => navigate({ to: "/settings/users" })}
          />
          <SettingsRow
            icon={<Library size={20} />}
            title="Libraries"
            onClick={() => navigate({ to: "/settings/libraries" })}
          />
          <SettingsRow
            icon={<Music size={20} />}
            title="Playback and transcoding"
            onClick={() => navigate({ to: "/settings/playback" })}
          />
        </Section>
      )}

      <Section label="System">
        <SettingsRow
          icon={<MonitorCog size={20} />}
          title="System information"
          subtitle={versionSubtitle}
          onClick={() => navigate({ to: "/settings/system" })}
        />
        {isAdmin && (
          <>
            <SettingsRow
              icon={<CheckCircle size={20} />}
              title="Background jobs"
              onClick={() => navigate({ to: "/settings/jobs" })}
            />
            <SettingsRow
              icon={<Timer size={20} />}
              title="Scheduled tasks"
              onClick={() => navigate({ to: "/settings/tasks" })}
            />
          </>
        )}
      </Section>

      <Dialog open={changePasswordOpen} onOpenChange={setChangePasswordOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Change password</DialogTitle>
          </DialogHeader>
          <div className="flex flex-col gap-4 pt-2">
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="current-password">Current password</Label>
              <Input
                id="current-password"
                type="password"
                autoComplete="current-password"
              />
            </div>
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="new-password">New password</Label>
              <Input
                id="new-password"
                type="password"
                autoComplete="new-password"
              />
            </div>
            <div className="flex flex-col gap-1.5">
              <Label htmlFor="confirm-password">Confirm new password</Label>
              <Input
                id="confirm-password"
                type="password"
                autoComplete="new-password"
              />
            </div>
            <div className="flex justify-end gap-2 pt-2">
              <Button
                variant="outline"
                onClick={() => setChangePasswordOpen(false)}
              >
                Cancel
              </Button>
              <Button variant="primary">Save</Button>
            </div>
          </div>
        </DialogContent>
      </Dialog>

      <Dialog open={mfaOpen} onOpenChange={setMfaOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Set up multi-factor authentication</DialogTitle>
          </DialogHeader>
          <div className="flex flex-col items-center gap-4 pt-2">
            <Badge variant="secondary" className="flex items-center gap-1.5">
              <Clock size={12} />
              Coming soon
            </Badge>
            <p className="text-sm text-muted-foreground text-center">
              Multi-factor authentication setup will be available in a future
              update.
            </p>
            <Button variant="outline" onClick={() => setMfaOpen(false)}>
              Close
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  );
}
