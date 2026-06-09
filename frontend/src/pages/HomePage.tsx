// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { useNavigate } from "@tanstack/react-router";
import { BodyText, BoldText, Button } from "@swopstar/react-ui";
import { BarChart3, Upload } from "lucide-react";

export function HomePage() {
  const navigate = useNavigate();

  return (
    <div
      style={{
        flex: 1,
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        flexDirection: "column",
        gap: "12px",
        padding: "24px",
        textAlign: "center",
      }}
    >
      <BoldText>The library is currently empty.</BoldText>
      <BodyText>
        Add tracks by adding folders or uploading tracks to the library.
      </BodyText>
      <div style={{ display: "flex", gap: "8px", marginTop: "4px" }}>
        <Button variant="primary" onClick={() => navigate({ to: "/settings" })}>
          <BarChart3 size={16} />
          Manage library
        </Button>
        <Button>
          <Upload size={16} />
          Upload tracks
        </Button>
      </div>
    </div>
  );
}
