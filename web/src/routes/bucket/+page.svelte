<script lang="ts">
  import { page } from '$app/stores';
  import { getKeys } from '$lib/api';
  import { onMount } from 'svelte';

  let keys: string[] = [];
  let bucket = '';

  $: bucket = $page.params.name;

  onMount(async () => {
    const res = await getKeys(bucket);
    keys = res.keys;
  });
</script>

<h2 class="text-xl font-semibold mb-2">🔑 Keys in <code>{bucket}</code></h2>

<ul class="space-y-1">
  {#each keys as key}
    <li>
      <a href={`/bucket/${bucket}/key/${key}`} class="text-green-600 hover:underline">
        {key}
      </a>
    </li>
  {/each}
</ul>
