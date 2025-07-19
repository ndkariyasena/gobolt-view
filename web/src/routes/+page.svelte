<script lang="ts">
	import { startDbConnection } from '$lib/api';
	import { saveConfig } from '$lib/storage';

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
			const { data, status } = await startDbConnection(payload);
      console.log({data, status})
			loading = false;

			if (data.message) {
				message = data.message;
			} else {
				error = data.error;
			}

      if (status == 200) {
        saveConfig({
          dbName: name,
          dbPort: parseInt(port, 10),
          dbFilePath: filePath
        });
        window.location.href = '/bucket'
      }
		} catch (_error) {
      error = 'Something went wrong. Please try again';
    }
	}
</script>

<div
	class="card preset-filled-surface-100-900 bg-card mx-auto mt-10 w-full max-w-2xl max-w-md space-y-6 rounded-xl border p-4 p-6 text-center shadow-lg"
>
  <h1 class="h1 text-center">🔗 Connect to Database</h1>

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
