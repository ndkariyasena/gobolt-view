export interface IsDbConnectedResponse {
  isConnected: boolean;
}

export interface StartDbConnectionResponse {
  status: number;
  message: string | null;
  error: string | null;
}

export interface BucketDetails {
  status: number;
  bucketDetails: Record<string, number>;
}

export type KeyValues = Array<{ key: string; value: string }>;

export interface KeyValuesResponse {
  status: number;
  keyValues: KeyValues;
}

export interface ValueResponse {
  status: number;
  value: string | null;
  error: string | null;
}

export interface ConnectionParams {
  dbFilePath: string;
}

export interface FetchError {
  status: number;
  message: string;
}