<script lang="ts">
	import { startDbConnection } from '$lib/api';

	let name = '';
	let port = '';
	let filePath = '';

	let message = '';
	let error = '';
	let loading = false;

	async function submit() {
		loading = true;
		error = '';
		message = '';

		const payload = {
			dbName: name,
			dbPort: parseInt(port, 10),
			dbFilePath: filePath
		};

		try {
			const result = await startDbConnection(payload);
			console.log({ result });
			loading = false;

			if (result.message) {
				message = result.message;
			} else {
				error = result.error;
			}
		} catch (_error) {
      error = 'Something went wrong. Please try again';
    }
	}
</script>

<div
	class="card preset-filled-surface-100-900 bg-card mx-auto mt-10 w-full max-w-2xl max-w-md space-y-6 rounded-xl border p-4 p-6 text-center shadow-lg"
>
	<h2 class="mb-4 text-2xl font-bold">🔗 Connect to BoltDB</h2>

	<!-- User Info -->
	<div class="grid grid-cols-1 gap-4 md:grid-cols-2">
		<div>
			<label for="" class="text-muted-foreground mb-1 block text-sm font-medium">DB Name</label>
			<input id="db-name" type="text" bind:value={name} placeholder="Enter DB name" class="input" />
		</div>

		<div>
			<label for="db-port" class="text-muted-foreground mb-1 block text-sm font-medium"
				>DB Port</label
			>
			<input id="db-port" type="text" bind:value={port} placeholder="Enter DB port" class="input" />
		</div>
	</div>

	<!-- File Selection -->
	<div>
		<label for="db-path" class="text-muted-foreground mb-2 block text-sm font-medium"
			>BoltDB File Path</label
		>
		<input
			id="db-path"
			type="text"
			bind:value={filePath}
			placeholder="/path/to/database.db"
			class="input"
		/>
	</div>

	<!-- Submit Button -->
	<div class="flex justify-end">
		<!-- <button type="button" class="btn preset-tonal-primary">Button</button> -->
		<button on:click={submit} class="btn preset-tonal-primary" disabled={loading}>
			{#if loading}
				<span class="animate-pulse">Connecting…</span>
			{:else}
				Connect
			{/if}
		</button>
	</div>

	<!-- Feedback -->
	{#if message}
		<p class="mt-2 text-sm text-green-600">{message}</p>
	{/if}
	{#if error}
		<p class="mt-2 text-sm text-red-600">{error}</p>
	{/if}
</div>
