export const isJson = (str: string) => {
	try {
		JSON.parse(str);
		return true;
	} catch {
		return false;
	}
};

export const formatJson = (str: string) => {
	try {
		return JSON.stringify(JSON.parse(str), null, 2);
	} catch {
		return str;
	}
};
