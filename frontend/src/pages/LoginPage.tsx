// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { useState } from "react";
import { useNavigate } from "@tanstack/react-router";
import {
  Alert,
  AlertDescription,
  Button,
  Card,
  CardContent,
  CardHeader,
  Input,
  Label,
  Spinner,
} from "@swopstar/react-ui";
import { LogIn } from "lucide-react";
import { useCreateSession } from "../api/auth/auth";
import type { CreateSession200 } from "../api/schemas/createSession200";
import { setAuthTokens } from "../auth";

const ALL_ENTITLEMENTS = ["admin", "tag", "upload", "internal"];

export function LoginPage() {
  const navigate = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [pin, setPin] = useState("");
  const [pinRequired, setPinRequired] = useState(false);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  const mutation = useCreateSession({
    mutation: {
      onSuccess: (data) => {
        const session = data as unknown as CreateSession200;
        if (session.accessToken && session.refreshToken) {
          setAuthTokens(
            session.accessToken,
            session.refreshToken,
            session.user?.username,
          );
          navigate({ to: "/home" });
        }
      },
      onError: (error) => {
        const body = (error as { response?: { data?: unknown } })?.response
          ?.data as { error?: string; parameters?: Record<string, string> };

        if (
          body?.error === "incorrect_credentials" &&
          body?.parameters?.pin === "required"
        ) {
          setPinRequired(true);
          setErrorMsg("Please enter your two-factor PIN.");
          return;
        }

        setErrorMsg("Incorrect username or password.");
      },
    },
  });

  return (
    <div
      style={{
        minHeight: "100svh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        padding: "16px",
      }}
    >
      <Card style={{ width: "100%", maxWidth: "400px" }}>
        <CardHeader>
          <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
            <img src="/favicon.svg" width={36} height={34} alt="" />
            <span
              style={{
                fontSize: "22px",
                fontVariationSettings: "'wght' 800, 'YTLC' 540",
              }}
            >
              swoptape
            </span>
          </div>
        </CardHeader>
        <CardContent>
          <form
            onSubmit={(e) => {
              e.preventDefault();
              setErrorMsg(null);
              mutation.mutate({
                data: {
                  username,
                  password,
                  entitlements: ALL_ENTITLEMENTS,
                  ...(pinRequired && pin ? { pin } : {}),
                },
              });
            }}
            style={{ display: "flex", flexDirection: "column", gap: "16px" }}
          >
            {errorMsg && (
              <Alert variant="destructive">
                <AlertDescription>{errorMsg}</AlertDescription>
              </Alert>
            )}
            <div
              style={{ display: "flex", flexDirection: "column", gap: "6px" }}
            >
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                autoComplete="username"
                disabled={pinRequired}
                required
              />
            </div>
            <div
              style={{ display: "flex", flexDirection: "column", gap: "6px" }}
            >
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                autoComplete="current-password"
                disabled={pinRequired}
                required
              />
            </div>
            {pinRequired && (
              <div
                style={{ display: "flex", flexDirection: "column", gap: "6px" }}
              >
                <Label htmlFor="pin">Two-factor PIN</Label>
                <Input
                  id="pin"
                  type="text"
                  inputMode="numeric"
                  pattern="[0-9]*"
                  maxLength={6}
                  value={pin}
                  onChange={(e) => setPin(e.target.value)}
                  autoComplete="one-time-code"
                  autoFocus
                  required
                />
              </div>
            )}
            <div style={{ display: "flex", justifyContent: "flex-end" }}>
              <Button
                type="submit"
                variant="primary"
                disabled={mutation.isPending}
              >
                {mutation.isPending ? <Spinner /> : <LogIn size={16} />}
                Sign in
              </Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
