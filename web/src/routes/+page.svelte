<script lang="ts">
  let name = '';
  let city = '';
  let file: File | null = null;
  let filePath = '';
  let useFilePicker = true;

  let message = '';
  let error = '';
  let loading = false;

  async function submit() {
    loading = true;
    error = '';
    message = '';

    const payload = {
      name,
      city,
      path: useFilePicker && file ? await readFilePath(file) : filePath,
    };

    const res = await fetch('http://localhost:8080/config/db', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });

    const result = await res.json();
    loading = false;

    if (res.ok) {
      message = result.message;
    } else {
      error = result.error;
    }
  }

  async function readFilePath(file: File): Promise<string> {
    // Normally, browsers don't expose actual file paths due to security reasons.
    // You may need to read the file's content or upload it instead.
    // For now, we'll use file.name as a placeholder.
    return file.name;
  }
</script>

<div class="card w-full max-w-md preset-filled-surface-100-900 p-4 text-center max-w-2xl mx-auto mt-10 p-6 border rounded-xl bg-card shadow-lg space-y-6">
  <h2 class="text-2xl font-bold mb-4">🔗 Connect to BoltDB</h2>

  <!-- User Info -->
  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
    <div>
      <label for="" class="block text-sm font-medium text-muted-foreground mb-1">DB Name</label>
      <input id="db-name" type="text" bind:value={name} placeholder="Enter DB name"
        class="input" />
    </div>

    <div>
      <label for="db-port" class="block text-sm font-medium text-muted-foreground mb-1">DB Port</label>
      <input id="db-port" type="text" bind:value={city} placeholder="Enter DB port"
        class="input" />
    </div>
  </div>

  <!-- File Selection -->
  <div>
    <label for="db-path" class="block text-sm font-medium text-muted-foreground mb-2">BoltDB File Path</label>
    <input id="db-path" type="text" bind:value={filePath} placeholder="/path/to/database.db"
        class="input" />
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
    <p class="text-green-600 text-sm mt-2">{message}</p>
  {/if}
  {#if error}
    <p class="text-red-600 text-sm mt-2">{error}</p>
  {/if}
</div>
