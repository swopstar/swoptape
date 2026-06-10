#!/usr/bin/env node
// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

import { readFileSync, writeFileSync, existsSync, readdirSync } from "fs";
import { join, dirname } from "path";
import { fileURLToPath } from "url";

const root = join(dirname(fileURLToPath(import.meta.url)), "..");
const GO_CSV = join(root, "tmp/notices/go.csv");
const GO_TEXTS = join(root, "tmp/notices/go-texts");
const GO_VERSIONS = join(root, "tmp/notices/go-versions.txt");
const NPM_JSON = join(root, "tmp/notices/npm.json");
const AGPL_TEXT = join(root, "LICENSES/AGPL-3.0-only.txt");
const OUTPUT = join(root, "internal/notices/data.json");

function findLicenseInDir(dir) {
  if (!existsSync(dir)) return null;
  const entries = readdirSync(dir);
  const file = entries.find((e) =>
    /^(un)?licen[sc]e|^copying|^notice/i.test(e),
  );
  return file ? join(dir, file) : null;
}

function findLicenseFile(baseDir, modulePath) {
  const parts = modulePath.split("/").filter(Boolean);
  for (let depth = parts.length; depth > 0; depth--) {
    const found = findLicenseInDir(join(baseDir, ...parts.slice(0, depth)));
    if (found) return found;
  }
  return null;
}

function readText(path) {
  if (!path || !existsSync(path)) return "";
  return readFileSync(path, "utf8").trim();
}

// --- Go versions ---
const goVersions = new Map();
const goVersionsRaw = readFileSync(GO_VERSIONS, "utf8").trim();
for (const line of goVersionsRaw.split("\n")) {
  const parts = line.split(" ");
  if (parts.length >= 2) goVersions.set(parts[0], parts[1]);
}

// --- Go ---
const goEntries = [];
const goCsv = readFileSync(GO_CSV, "utf8").trim();
for (const line of goCsv.split("\n")) {
  const [module, , spdx] = line.split(",");
  const version = goVersions.get(module) ?? "";
  const licenseFile = findLicenseFile(GO_TEXTS, module);
  goEntries.push({ module, version, spdx, text: readText(licenseFile) });
}

// --- npm ---
const npmEntries = [];
const npmRaw = JSON.parse(readFileSync(NPM_JSON, "utf8"));
const seen = new Set();
for (const [key, info] of Object.entries(npmRaw)) {
  const at = key.lastIndexOf("@");
  const name = key.slice(0, at);
  const version = key.slice(at + 1);
  if (seen.has(`${name}@${version}`)) continue;
  seen.add(`${name}@${version}`);
  const licenseFile =
    info.licenseFile ?? (info.path ? findLicenseInDir(info.path) : null);
  npmEntries.push({
    module: name,
    version,
    spdx: info.licenses ?? "",
    text: readText(licenseFile),
  });
}
npmEntries.sort((a, b) => a.module.localeCompare(b.module));

const app = { spdx: "AGPL-3.0-only", text: readText(AGPL_TEXT) };
const output = { app, go: goEntries, npm: npmEntries };
writeFileSync(OUTPUT, JSON.stringify(output, null, 2) + "\n");
console.log(
  `wrote app + ${goEntries.length} Go + ${npmEntries.length} npm entries → ${OUTPUT}`,
);
