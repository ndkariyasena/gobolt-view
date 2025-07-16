const API = import.meta.env.VITE_API_BASE_URL;

export async function getBuckets() {
  const res = await fetch(`${API}/buckets`);
  return res.json();
}

export async function getKeys(bucket: string) {
  const res = await fetch(`${API}/bucket/${bucket}/keys`);
  return res.json();
}

export async function getValue(bucket: string, key: string) {
  const res = await fetch(`${API}/bucket/${bucket}/key/${key}`);
  return res.json();
}
