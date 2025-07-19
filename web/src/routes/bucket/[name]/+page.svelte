<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { Modal } from '@skeletonlabs/skeleton-svelte';
	import { getKeyValuePares } from '$lib/api';

	let bucketName = '';
	let key = '';
	let value = '';
	let items: { key: string; value: string }[] = [];
	let openState = false;

	$: bucketName = $page.params.name;

	function modalClose() {
		openState = false;
	}

  function modalViewHandler() {
    openState = !openState;
  }

	onMount(async () => {
		/* const dbPath = localStorage.getItem('gobolt-db-path');
    if (!dbPath) return goto('/'); */

		const res = await getKeyValuePares(bucketName);
		console.log({ res });
		if (res.keyValues && res.keyValues.length > 0) {
			items = res.keyValues;
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
	function isJson(str: string) {
		try {
			JSON.parse(str);
			return true;
		} catch {
			return false;
		}
	}

	function formatJson(str: string) {
		try {
			return JSON.stringify(JSON.parse(str), null, 2);
		} catch {
			return str;
		}
	}
</script>

<section class="container mx-auto max-w-7xl space-y-6 p-6">
	<h1 class="h1 text-center">📂 Bucket: <span class="text-primary">{bucketName}</span></h1>

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

		{#if items.length === 0}
			<p class="text-muted">No items found in this bucket.</p>
		{/if}
	</div>
	<div class="overflow-x-auto">
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
				{#each items as item}
					<tr>
						<td class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-700">{item.key}</td>
						<td
							class="break-inside-avoid-column border border-gray-300 bg-slate-700 p-2 dark:border-gray-700"
						>
							{#if isJson(item.value)}
								<pre class="h-50 overflow-y-auto whitespace-pre-wrap">{formatJson(item.value)}</pre>
							{:else}
								{item.value}
							{/if}
						</td>
						<td class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-700">
							<button type="button" class="btn btn-sm preset-filled" onclick={modalViewHandler}>View</button>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<Modal
		open={openState}
		onOpenChange={(e) => (openState = e.open)}
		triggerBase="btn preset-tonal"
		contentBase="card bg-surface-100-900 p-4 space-y-4 shadow-xl max-w-screen-sm"
		backdropClasses="backdrop-blur-sm"
	>
		{#snippet content()}
			<header class="flex justify-between">
				<h2 class="h2">Modal Example</h2>
			</header>
			<article>
				<p class="opacity-60">
					Lorem ipsum dolor sit amet consectetur adipisicing elit. Nam, ab adipisci. Libero cumque
					sunt quis error veritatis amet, expedita voluptatem. Quos repudiandae consequuntur
					voluptatem et dicta quas, reprehenderit velit excepturi?
				</p>
			</article>
			<footer class="flex justify-end gap-4">
				<button type="button" class="btn preset-filled" onclick={modalViewHandler}>Close</button>
			</footer>
		{/snippet}
	</Modal>
</section>
