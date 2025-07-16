<script lang="ts">
  import { page } from '$app/stores';
  import { getValue } from '$lib/api';
  import { onMount } from 'svelte';

  let value = '';
  let error = '';
  let bucket = '', key = '';

  $: bucket = $page.params.name;
  $: key = $page.params.key;

  onMount(async () => {
    const res = await getValue(bucket, key);
    value = res?.value ?? '';
    error = res?.error ?? '';
  });
</script>

<h3 class="text-lg font-medium">🧾 Value for <code>{key}</code></h3>

{#if error}
  <div class="text-red-500 mt-4">{error}</div>
{:else}
  <pre class="bg-gray-100 p-4 mt-2 rounded overflow-x-auto">{value}</pre>
{/if}
