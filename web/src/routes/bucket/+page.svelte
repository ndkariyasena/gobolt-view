<script lang="ts">
	import { onMount } from 'svelte';
	import Search from '@lucide/svelte/icons/search';
	import ArrowRight from '@lucide/svelte/icons/arrow-right';
	import { getBuckets } from '$lib/api';
	import { goto } from '$app/navigation';

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
	<h1 class="h1">🪣 Buckets</h1>

	<button
		type="button"
		class="btn btn-m preset-filled"
		onclick={() => goto('/query-view')}
	>
		<span>Query View</span>
    <ArrowRight size={18} />
	</button>

	<div class="relative w-full flex-1">
		<input
			type="text"
			placeholder="Search buckets..."
			bind:value={search}
			class="form-input focus:outline-hidden w-full rounded-md p-2 pl-10 focus:border-hidden focus:ring-2 focus:ring-indigo-600"
		/>
		<span class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
			<Search size={18} />
		</span>
	</div>

	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-3">
		{#each filtered as bucket (bucket)}
			<div
				class="card card-body flex cursor-pointer items-center justify-between bg-slate-700 p-2 shadow-lg hover:bg-sky-500"
			>
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
