<script lang="ts">
  import { AppBar } from '@skeletonlabs/skeleton-svelte';
  import ArrowLeft from '@lucide/svelte/icons/arrow-left';
  import Paperclip from '@lucide/svelte/icons/paperclip';
  import Calendar from '@lucide/svelte/icons/calendar';
  import CircleUser from '@lucide/svelte/icons/circle-user';

  let dbPath = '';
  let message = '';
  let error = '';

  async function connect() {
    const res = await fetch('http://localhost:8080/config/db', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ path: dbPath })
    });

    const result = await res.json();
    if (res.ok) {
      message = result.message;
      error = '';
    } else {
      message = '';
      error = result.error;
    }
  }
</script>

<AppBar>
  {#snippet lead()}
    <ArrowLeft size={24} />
  {/snippet}
  {#snippet trail()}
    <Paperclip size={20} />
    <Calendar size={20} />
    <CircleUser size={20} />
  {/snippet}
  {#snippet headline()}
    <h2 class="h2">Headline</h2>
  {/snippet}
</AppBar>

<section class="space-y-4">
  <h2 class="text-xl font-semibold">🔗 Connect to BoltDB File</h2>

  <input type="text" bind:value={dbPath} placeholder="/path/to/my.db"
    class="form-input w-full p-2 rounded border" />

  <button on:click={connect} class="btn btn-primary">
    Connect
  </button>

  {#if message}
    <p class="text-green-600">{message}</p>
  {/if}
  {#if error}
    <p class="text-red-600">{error}</p>
  {/if}
</section>

