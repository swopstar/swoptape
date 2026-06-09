import { defineConfig } from "orval";

export default defineConfig({
  swoptape: {
    input: "../www/api/v0/openapi.yaml",
    output: {
      mode: "tags-split",
      target: "src/api",
      schemas: "src/api/schemas",
      client: "react-query",
      override: {
        mutator: {
          path: "src/api/client.ts",
          name: "apiClient",
        },
      },
    },
  },
});
