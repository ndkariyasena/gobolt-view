import type {
	IsDbConnectedResponse,
	StartDbConnectionResponse,
	BucketDetails,
	KeyValuesResponse,
	ValueResponse,
	ConnectionParams
} from './types';

const API = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000';

export const isDbConnected = async (): Promise<IsDbConnectedResponse> => {
	return await fetch(`${API}/api/connection/is-active`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error checking DB connection:', err);
			return { isConnected: false };
		});
};

export const startDbConnection = async (
	payload: ConnectionParams
): Promise<StartDbConnectionResponse> => {
	return await fetch(`${API}/api/config/db`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify(payload)
	})
		.then(async (res) => {
			const data = await res.json();
			return {
				status: res.status,
				message: data.message || null,
				error: data.error || null
			} as StartDbConnectionResponse;
		})
		.catch((err) => {
			console.error('Error starting DB connection:', err);
			return {
				status: 400,
				message: null,
				error: 'Failed to start DB connection'
			} as StartDbConnectionResponse;
		});
};

export async function getBuckets(): Promise<BucketDetails> {
	return await fetch(`${API}/api/buckets`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching buckets:', err);
			return { status: 500, bucketDetails: {} };
		});
}

export async function getKeys(bucket: string): Promise<KeyValuesResponse> {
	return await fetch(`${API}/api/bucket/${bucket}/keys`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching keys:', err);
			return { status: 500, keyValues: [] };
		});
}

export async function getValue(bucket: string, key: string): Promise<ValueResponse> {
	return await fetch(`${API}/api/bucket/${bucket}/key/${key}`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching value:', err);
			return { status: 500, value: null };
		});
}
