import type { connectionParams } from "./api";

export function saveConfig(config: connectionParams) {
  sessionStorage.setItem("db_config", JSON.stringify(config));
}

export function getConfig(): connectionParams | null {
  const item = sessionStorage.getItem("db_config");
  return item ? JSON.parse(item) : null;
}
