<script lang="ts">
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import Search from '@lucide/svelte/icons/search';
	import { Modal } from '@skeletonlabs/skeleton-svelte';
	import { getKeyValuePares } from '$lib/api';
	import { formatJson, isJson } from '$lib/utils';

	let bucketName = '';
	let key = '';
	let value = '';
	let bucketKeyValues: { key: string; value: string }[] = [];
	let openState = false;
	let modalContent = '';
	let modalHeader = 'Key-Value Details';
	let search = '';

	$: filteredKeyValues = search
		? bucketKeyValues.filter((i) => i.key.toLowerCase().includes(search.toLowerCase()))
		: bucketKeyValues;

	$: bucketName = page.params.name;

	function modalViewHandler(keyName?: string) {
		if (keyName) {
			modalHeader = `Details for Key: ${keyName}`;
			modalContent = bucketKeyValues.find((item) => item.key === keyName)?.value || '';
		} else {
			modalHeader = 'Key-Value Details';
			modalContent = '';
		}
		openState = !openState;
	}

	onMount(async () => {
		const res = await getKeyValuePares(bucketName);

		if (res.keyValues && res.keyValues.length > 0) {
			bucketKeyValues = res.keyValues;
		}
	});

	/* async function handleAdd() {
    const res = await fetch(`/api/bucket/${bucketName}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ key, value }),
    });
    if (res.ok) {
      items.push({ key, value });
      key = '';
      value = '';
    }
  }

  async function deleteItem(k: string) {
    const res = await fetch(`/api/bucket/${bucketName}/${k}`, {
      method: 'DELETE',
    });
    if (res.ok) {
      items = items.filter((i) => i.key !== k);
    }
  } */
</script>

<section class="container mx-auto max-w-7xl space-y-6 p-6">
	<h1 class="h1 text-center">📂 Bucket: <span class="text-primary">{bucketName}</span></h1>

	<div class="relative w-full flex-1">
		<input
			type="text"
			placeholder="Search keys..."
			bind:value={search}
			class="form-input focus:outline-hidden w-full rounded-md p-2 pl-10 focus:border-hidden focus:ring-2 focus:ring-indigo-600"
		/>
		<span class="pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-gray-400">
			<Search size={18} />
		</span>
	</div>

	<!-- Add New Entry Form -->
	<!-- <div class="card card-body space-y-4">
		<h2 class="h3">➕ Add Entry</h2>

		<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
			<input class="form-input" type="text" placeholder="Key" bind:value={key} />
			<input class="form-input" type="text" placeholder="Value" bind:value />
		</div>

		<div class="flex justify-end">
      <button class="btn btn-filled-primary" on:click={handleAdd}>Add Entry</button>
    </div>
	</div> -->

	<!-- Key-Value List -->
	<div class="card card-body space-y-4">
		<h2 class="h3">📋 Entries</h2>

		{#if bucketKeyValues.length === 0}
			<p class="text-muted">No items found in this bucket.</p>
		{/if}
	</div>
	<div class="overflow-x-auto">
		{#if bucketKeyValues.length > 0}
			<table
				class="relative min-w-full border-separate border-spacing-2 border border-gray-400 dark:border-gray-500"
			>
				<thead>
					<tr>
						<th class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-600">Key</th>
						<th class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-600">Value</th>
						<th class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-600">Action</th>
					</tr>
				</thead>
				<tbody>
					{#each filteredKeyValues as item}
						<tr>
							<td class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-700"
								>{item.key}</td
							>
							<td
								class="break-inside-avoid-column border border-gray-300 bg-slate-700 p-2 dark:border-gray-700"
							>
								{#if isJson(item.value)}
									<pre class="h-50 overflow-y-auto whitespace-pre-wrap">{formatJson(
											item.value
										)}</pre>
								{:else}
									<p
										class="max-w-250 max-h-80 overflow-y-auto whitespace-pre-wrap break-words text-justify"
									>
										{item.value}
									</p>
								{/if}
							</td>
							<td class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-700">
								<button
									type="button"
									class="btn btn-sm preset-filled"
									onclick={() => modalViewHandler(item.key)}>View</button
								>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		{/if}
	</div>

	<Modal
		open={openState}
		onOpenChange={(e) => (openState = e.open)}
		triggerBase="btn preset-tonal"
		contentBase="card bg-surface-100-900 p-5 space-y-4 shadow-xl w-dvh h-75 max-w-full max-h-75vh overflow-y-auto flex flex-col h-full"
		backdropClasses="backdrop-blur-sm"
	>
		{#snippet content()}
			<header class="flex justify-between">
				<h2 class="h2">{modalHeader}</h2>
			</header>
			<article class="flex-1 overflow-y-auto">
				{#if isJson(modalContent)}
					<pre class="h-75 overflow-y-auto whitespace-pre-wrap">{formatJson(modalContent)}</pre>
				{:else}
					<p
						class="max-h-80 overflow-y-auto whitespace-pre-wrap break-words text-justify opacity-60"
					>
						{modalContent}
					</p>
				{/if}
			</article>
			<footer class="flex shrink-0 justify-end gap-4">
				<button type="button" class="btn preset-filled" onclick={() => modalViewHandler(undefined)}>
					Close
				</button>
			</footer>
		{/snippet}
	</Modal>
</section>
