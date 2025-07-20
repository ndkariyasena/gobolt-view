const API = import.meta.env.VITE_API_BASE_URL;

export interface connectionParams {
	dbName: string;
	dbPort: number;
	dbFilePath: string;
}

export const isDbConnected = async () => {
  const res = await fetch(`${API}/api/connection/is-active`);
	return res.json();
}

export const startDbConnection = async (payload: connectionParams) => {
	const res = await fetch(`${API}/api/config/db`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload)
	});
	return {data: await res.json(), status: res.status};
};

export async function getBuckets() {
	const res = await fetch(`${API}/api/buckets`);
	return res.json();
}

export async function getKeys(bucket: string) {
	const res = await fetch(`${API}/api/bucket/${bucket}/keys`);
	return res.json();
}

export async function getValue(bucket: string, key: string) {
	const res = await fetch(`${API}/api/bucket/${bucket}/key/${key}`);
	return res.json();
}

export async function getKeyValuePares(bucket: string) {
	const res = await fetch(`${API}/api/bucket/${bucket}/key-value`);
	return res.json();
}
