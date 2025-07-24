<script lang="ts">
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import Search from '@lucide/svelte/icons/search';
	import { Modal, Toaster, createToaster } from '@skeletonlabs/skeleton-svelte';
	import { getKeys, addValue, updateValue, deleteKey } from '$lib/api';
	import type { KeyValues, KeyValuesResponse } from '$lib/types';
	import { formatJson, isJson } from '$lib/utils';

	const toaster = createToaster({
		placement: 'top-end'
	});

	let bucketName = '';
	let bucketKeyValues: KeyValues = [];
	let openState = false;
	let editState = false;
	let actionType = '';
	let selectedKey = '';
	let modalContent = '';
	let modalHeader = 'Key-Value Details';
	let search = '';

	$: filteredKeyValues = search
		? bucketKeyValues.filter((i) => i.key.toLowerCase().includes(search.toLowerCase()))
		: bucketKeyValues;

	$: bucketName = page.params.name;

  const reloadPage = () => {
    setTimeout(() => {
      window.location.reload(); // Reload to reflect changes
    }, 1000); // Ensure the UI reflects the changes. Time = 1000ms
  };

	const getValueForKey = (keyName: string) => {
		const item = bucketKeyValues.find((item) => item.key === keyName);
		return item ? item.value : '';
	};

	function modalViewHandler(keyName?: string, action?: string) {
		actionType = action || 'view';
		if (keyName) {
			modalHeader = `Details for Key: ${keyName}`;
			modalContent = getValueForKey(keyName);
		} else {
			modalHeader = 'Key-Value Details';
			modalContent = '';
		}
		openState = !openState;
	}

	const modalEditHandler = (keyName?: string, action?: string) => {
		actionType = action || 'edit';
		if (keyName) {
			selectedKey = keyName;
			modalHeader = keyName;
			modalContent = getValueForKey(keyName);

			modalContent = isJson(modalContent) ? formatJson(modalContent) : modalContent;
		} else {
			selectedKey = '';
			modalHeader = '';
			modalContent = '';
		}
		editState = !editState;
	};

	const updateCurrentList = () => {
		const index = bucketKeyValues.findIndex((item) => item.key === selectedKey);
		if (index !== -1) {
			bucketKeyValues[index].value = modalContent;
		} else {
			bucketKeyValues.push({ key: selectedKey, value: modalContent });
		}
		toaster.success({
			title: 'Edits saved successfully!',
			closable: true
		});
	};

	const handleNewEntry = async () => {
		if (modalHeader.trim() === '') {
			toaster.error({
				title: 'Key cannot be empty.',
				closable: true
			});
			return;
		} else {
      const index = bucketKeyValues.findIndex((item) => item.key === modalHeader);

      if (index !== -1) {
        toaster.error({
          title: `Key "${modalHeader}" already exists. Please choose a different key.`,
          closable: true
        });
        return;
      }

			await addValue(bucketName, modalHeader, modalContent)
				.then((res) => {
					if (res.status === 200) {
						bucketKeyValues.push({ key: modalHeader, value: modalContent });
						toaster.success({
							title: `Key "${modalHeader}" added successfully!`,
							closable: true
						});
					} else {
						toaster.error({
							title: `Failed to add key "${modalHeader}". Please try again.`,
							closable: true
						});
					}
				})
				.catch((error) => {
					console.error('Error adding new entry:', error);
					toaster.error({
						title: 'Failed to add new entry. Please try again.',
						closable: true
					});
				});
		}
	};

	const updateItem = async () => {
		const value = getValueForKey(selectedKey);

		if ((selectedKey !== modalHeader || modalContent !== value) && modalHeader.trim() !== '') {
			await updateValue(bucketName, selectedKey, modalContent, selectedKey)
				.then(updateCurrentList)
				.catch((error) => {
					console.error('Error updating item:', error);
					toaster.error({
						title: 'Failed to update item. Please try again.',
						closable: true
					});
				});
		} else {
			if (modalHeader.trim() === '') {
				toaster.error({
					title: 'Key cannot be empty.',
					closable: true
				});
			} else {
				toaster.info({
					title: 'No changes detected.',
					closable: true
				});
			}
		}
	};

	const saveEdits = async () => {
		if (actionType === 'add') {
			await handleNewEntry();
		} else if (actionType === 'edit') {
			await updateItem();
		}
		editState = false;
    
    reloadPage(); // Reload to reflect changes
	};

	const deleteItem = async (keyName: string) => {
		if (confirm(`Are you sure you want to delete the key "${keyName}"?`)) {
			await deleteKey(bucketName, keyName)
				.then((res) => {
					if (res.status === 200) {
						bucketKeyValues = bucketKeyValues.filter((item) => item.key !== keyName);
						toaster.success({
							title: `Key "${keyName}" deleted successfully!`,
							closable: true
						});
    
            reloadPage(); // Reload to reflect changes
					} else {
						toaster.error({
							title: `Failed to delete key "${keyName}". Please try again.`,
							closable: true
						});
					}
				})
				.catch((error) => {
					console.error('Error deleting key:', error);
					toaster.error({
						title: 'Failed to delete key. Please try again.',
						closable: true
					});
				});
		}
	};

	onMount(async () => {
		const res: KeyValuesResponse = await getKeys(bucketName);

		if (res.keyValues && res.keyValues.length > 0) {
			bucketKeyValues = res.keyValues;
		}
	});
</script>

<Toaster {toaster}></Toaster>

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
	<div class="card card-body space-y-4">
		<h2 class="h3">➕ Add Entry</h2>

		<div class="flex justify-start">
			<button
				class="btn preset-filled-success-500"
				onclick={() => modalEditHandler(undefined, 'add')}>Add Entry</button
			>
		</div>
	</div>

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
								<p
									class="max-w-200 max-h-80 overflow-y-auto whitespace-pre-wrap break-words text-justify"
								>
									{item.value}
								</p>
							</td>
							<td class="border border-gray-300 bg-slate-700 p-2 dark:border-gray-700">
								<div class="mb-5 flex flex-row items-center justify-center gap-2">
									<button
										type="button"
										class="btn btn-sm preset-filled mr-2"
										onclick={() => modalViewHandler(item.key, 'view')}
									>
										View
									</button>
									<button
										type="button"
										class="btn btn-sm preset-filled-warning-500"
										onclick={() => modalEditHandler(item.key, 'edit')}
									>
										Edit
									</button>
								</div>
								<div class="mb-4 flex flex-row items-center justify-center gap-2">
									<button
										type="button"
										class="btn btn-sm preset-filled-error-500"
										onclick={() => deleteItem(item.key)}
									>
										Delete
									</button>
								</div>
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
				<button
					type="button"
					class="btn preset-filled"
					onclick={() => modalViewHandler(undefined, '')}
				>
					Close
				</button>
			</footer>
		{/snippet}
	</Modal>

	<Modal
		open={editState}
		onOpenChange={(e) => (editState = e.open)}
		triggerBase="btn preset-tonal"
		contentBase="card bg-surface-100-900 p-5 space-y-4 shadow-xl w-dvh h-75 max-w-full max-h-75vh overflow-y-auto flex flex-col h-full"
		backdropClasses="backdrop-blur-sm"
	>
		{#snippet content()}
			<header class="flex justify-between">
				<h2 class="h2">Edit {selectedKey}</h2>
			</header>
			<article class="flex-1 overflow-y-auto">
				<input
					type="text"
					bind:value={modalHeader}
					class="form-input mb-5 w-full bg-slate-700"
					placeholder="Edit key here..."
					tabindex="0"
				/>
				<textarea
					rows="20"
					bind:value={modalContent}
					class="form-textarea mt-2 w-full bg-slate-700"
					placeholder="Edit value here..."
				></textarea>
			</article>
			<footer class="flex shrink-0 justify-end gap-4">
				<button type="button" class="btn preset-filled-tertiary-500" onclick={saveEdits}>
					Save
				</button>
				<button
					type="button"
					class="btn preset-filled"
					onclick={() => modalEditHandler(undefined, '')}
				>
					Close
				</button>
			</footer>
		{/snippet}
	</Modal>
</section>
