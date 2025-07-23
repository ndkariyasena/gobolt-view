<script lang="ts">
	import '../app.css';
  import { AppBar } from '@skeletonlabs/skeleton-svelte';
  import ArrowLeft from '@lucide/svelte/icons/arrow-left';
  import ArrowRight from '@lucide/svelte/icons/arrow-right';
  import Home from '@lucide/svelte/icons/home';
	import { onMount } from 'svelte';
	import { isDbConnected } from '$lib/api';
	import type { IsDbConnectedResponse } from '$lib/types';

	onMount(async () => {
		if (window.location.pathname != '/') {
			const res: IsDbConnectedResponse = await isDbConnected();
			if (!res.isConnected) {
        console.warn('Database not connected, redirecting to home');
				window.location.href = '/';
			}
		}
	});

	let { children } = $props();
</script>

<AppBar>
  {#snippet lead()}
    <button
      type="button"
      class="hover:bg-slate-700 rounded-full p-1"
      onclick={() => history.length > 1 ? history.back() : window.location.href = '/'}
      aria-label="Go Back"
    >
      <ArrowLeft size={24} />
    </button>
    <button
      type="button"
      class="hover:bg-slate-700 rounded-full p-1"
      onclick={() => history.length > 1 ? history.forward() : window.location.href = '/'}
      aria-label="Go Forward"
    >
      <ArrowRight size={24} />
    </button>
  {/snippet}
  {#snippet trail()}
    <a
      href="/bucket"
      class="hover:bg-slate-700 rounded-full p-1"
      aria-label="Buckets"
    >
      <Home size={30} />
    </a>
  {/snippet}
  <span>Go Bolt View</span>
</AppBar>

{@render children()}
