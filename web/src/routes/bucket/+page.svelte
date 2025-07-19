<script lang="ts">
	import { onMount } from 'svelte';
	import { getBuckets } from '$lib/api';

	let buckets: string[] = [];
	let bucketDetails: Record<string, number> = {};

	onMount(async () => {
		const res = await getBuckets();
		bucketDetails = res.bucketDetails;
    buckets = Object.keys(bucketDetails);
	});

	let search = '';

	$: filtered = search
		? buckets.filter((b) => b.toLowerCase().includes(search.toLowerCase()))
		: buckets;
</script>

<section class="container mx-auto space-y-6 p-6">
	<h1 class="h1">🔖 Buckets</h1>

	<div class="flex flex-col items-center gap-4 md:flex-row">
		<input
			type="text"
			placeholder="Search buckets..."
			bind:value={search}
			class="form-input flex-1 focus:outline-hidden focus:border-indigo-600 focus:outline-hidden rounded-md p-2"
		/>
	</div>

	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
		{#each filtered as bucket (bucket)}
			<div class="card card-body flex cursor-pointer items-center justify-between hover:bg-sky-500 p-2 shadow-lg bg-gray-900">
				<span class="font-semibold">{bucket}</span>
				<span class="font-italic">{bucketDetails[bucket]}</span>
				<a class="btn btn-sm preset-filled" href={`/bucket/${bucket}`}> View &rarr; </a>
			</div>
		{/each}
		{#if filtered.length === 0}
			<p class="text-muted-foreground col-span-full mt-4 text-center">No buckets found.</p>
		{/if}
	</div>
</section>

