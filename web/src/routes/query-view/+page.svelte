<script lang="ts">
	import { onMount } from 'svelte';
	import { Toaster, createToaster } from '@skeletonlabs/skeleton-svelte';
	import Search from '@lucide/svelte/icons/search';
	import Delete from '@lucide/svelte/icons/delete';
	import { getBuckets, getValue } from '$lib/api';
	import { formatJson, isJson } from '$lib/utils';

	const toaster = createToaster({
    placement: 'top-end'
  });

	let buckets: string[] = $state([]);
	let bucketDetails: Record<string, number> = {};
	let queryMap: Map<number, Record<string, string>> = $state(new Map());

	const addQuery = (canDelete = true) => {
		const id = Date.now();
		const newQuery = {
			id: id.toString(),
			bucket: '',
			key: '',
			results: '',
			canDelete: canDelete ? 'true' : 'false'
		};
		queryMap.set(id, newQuery);
		queryMap = new Map(queryMap); // Trigger reactivity
	};

	const deleteQuery = (id: number) => {
		if (queryMap.has(id)) {
			queryMap.delete(id);
			queryMap = new Map(queryMap); // Trigger reactivity
		} else {
			console.warn(`Query with ID ${id} does not exist.`);
		}
	};

	const search = async (id: string) => {
		const query = queryMap.get(Number(id));
		if (!query || !query.bucket || !query.key) {
			toaster.error({
				title: 'Please select a bucket and enter a key name before searching.',
        closable: true,
			});

      return;
		}

		const results = await getValue(query.bucket, query.key);
		if (results?.value) {
			query.results = results.value;
		} else {
			query.results = 'No results found';
			console.warn(`No results found for ID ${id}`);
      toaster.error({
				title: `No results found for bucket: ${query.bucket} and key: ${query.key}`,
        closable: true,
			});
		}

		queryMap.set(Number(id), query);
		queryMap = new Map(queryMap); // Trigger reactivity
	};

	const selectOnChange = (event: Event, id: string) => {
		if (queryMap.has(Number(id))) {
			const query = queryMap.get(Number(id));
			if (query) {
				query.bucket = (event.target as HTMLSelectElement).value;
			}
		}
	};

	const selectOnChangeKey = (event: Event, id: string) => {
		console.log({
			key: (event.target as HTMLInputElement).value,
			id
		});
		if (queryMap.has(Number(id))) {
			const query = queryMap.get(Number(id));
			if (query) {
				query.key = (event.target as HTMLInputElement).value;
			}
		}
	};

	onMount(async () => {
		addQuery(false);
		const res = await getBuckets();
		bucketDetails = res.bucketDetails;
		buckets = Object.keys(bucketDetails);
	});
</script>

<Toaster {toaster}></Toaster>

<div class="flex flex-col p-6">
	<!-- Page Header -->
	<div class="mb-6 flex items-center justify-between border-b pb-3">
		<h1 class="h1">📋 Query View</h1>
		<button onclick={() => addQuery()} class="btn preset-tonal-primary">Add</button>
	</div>

	<!-- Query Columns -->
	<div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3">
		{#each Array.from(queryMap.values()) as query (query.id + query.results)}
			<div class="space-y-3 rounded-2xl p-4 shadow-sm bg-slate-700 shadow-lg shadow-cyan-500/50">
				<!-- Bucket Dropdown -->
				<div>
					<label for="buckets-list" class="mb-1 block font-medium">Buckets</label>
					<select
						id="buckets-list"
						class="w-full rounded border p-2"
						onchange={(e) => selectOnChange(e, query.id)}
						value={query.bucket}
					>
						{#each buckets as bucket}
							<option value={bucket}>{bucket}</option>
						{/each}
					</select>
				</div>

				<!-- Key Input -->
				<div>
					<label for="key-input" class="mb-1 block font-medium">Key Name</label>
					<input
						id="key-input"
						type="text"
						placeholder="Enter key"
						class="w-full rounded border p-2"
						bind:value={query.key}
						onchange={(e) => selectOnChangeKey(e, query.id)}
					/>
				</div>

				<!-- Search Button -->
				<button onclick={() => search(query.id)} type="button" class="btn preset-filled-secondary-500">
					<span>Search</span>
					<Search size={18} />
				</button>

				<!-- Results Area -->
				<div class="min-h-[100px] rounded border p-4">
					{#if isJson(query.results)}
						<pre class="h-75 overflow-y-auto whitespace-pre-wrap">{formatJson(query.results)}</pre>
					{:else}
						<p
							class="max-h-80 overflow-y-auto whitespace-pre-wrap break-words text-justify opacity-60"
						>
							{query.results ? query.results : 'No results found'}
						</p>
					{/if}
				</div>

				<!-- Delete view -->
				{#if query.canDelete === 'true'}
					<button
						type="button"
						class="btn btn-sm preset-filled"
						onclick={() => deleteQuery(Number(query.id))}
					>
						<span>Delete</span>
            <Delete size={18} />
					</button>
				{/if}
			</div>
		{/each}
	</div>
</div>
