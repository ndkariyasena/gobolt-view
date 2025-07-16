<script lang="ts">
  import { onMount } from 'svelte';
  import { getBuckets } from '$lib/api';

  let buckets: string[] = [];

  onMount(async () => {
    const res = await getBuckets();
    buckets = res.buckets;
  });
</script>

<h1 class="text-2xl font-bold mb-4">📁 Buckets</h1>

<ul class="space-y-2">
  {#each buckets as bucket}
    <li>
      <a href={`/bucket/${bucket}`} class="text-blue-600 hover:underline">
        {bucket}
      </a>
    </li>
  {/each}
</ul>
