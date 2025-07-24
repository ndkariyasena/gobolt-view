import type {
	IsDbConnectedResponse,
	StartDbConnectionResponse,
	BucketDetails,
	KeyValuesResponse,
	ValueResponse,
	ConnectionParams,
	FetchResponse
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

export const getBuckets = async (): Promise<BucketDetails> => {
	return await fetch(`${API}/api/buckets`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching buckets:', err);
			return { status: 500, bucketDetails: {} };
		});
};

export const getKeys = async (bucket: string): Promise<KeyValuesResponse> => {
	return await fetch(`${API}/api/bucket/${bucket}/keys`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching keys:', err);
			return { status: 500, keyValues: [] };
		});
};

export const getValue = async (bucket: string, key: string): Promise<ValueResponse> => {
	return await fetch(`${API}/api/bucket/${bucket}/key/${key}`)
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error fetching value:', err);
			return { status: 500, value: null };
		});
};

export const updateValue = async (
	bucket: string,
	key: string,
	value: string,
	originalKey: string
): Promise<FetchResponse> => {
	return await fetch(`${API}/api/bucket/${bucket}/key/${key}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ value, originalKey })
	})
		.then((res) => res.json())
		.catch((err) => {
			console.error('Error updating value:', err);
			return { status: 500, value: null };
		});
};

export const addValue = async (
	bucket: string,
	key: string,
	value: string
): Promise<FetchResponse> => {
	return await fetch(`${API}/api/bucket/${bucket}/key/${key}`, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ value })
	})
		.then(async (res) => {
			const data = await res.json();
			return { status: res.status, message: data.message || 'Value added successfully' };
		})
		.catch((err) => {
			console.error('Error updating value:', err);
			return { status: 500, error: err.message || 'Internal server error' };
		});
};

export const deleteKey = async (bucket: string, key: string): Promise<FetchResponse> => {
	return await fetch(`${API}/api/bucket/${bucket}/key/${key}`, {
		method: 'DELETE'
	})
		.then(async (res) => {
			const data = await res.json();
			if (res.ok) {
				return { status: res.status, message: data.message || 'Key deleted successfully' };
			} else {
				return { status: res.status, error: data.message || 'Failed to delete key' };
			}
		})
		.catch((err) => {
			console.error('Error deleting key:', err);
			return { status: 500, error: 'Internal server error' };
		});
};
