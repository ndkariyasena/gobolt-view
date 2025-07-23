import type { ConnectionParams } from "./types";

export function saveConfig(config: ConnectionParams) {
  sessionStorage.setItem("db_config", JSON.stringify(config));
}

export function getConfig(): ConnectionParams | null {
  const item = sessionStorage.getItem("db_config");
  return item ? JSON.parse(item) : null;
}
