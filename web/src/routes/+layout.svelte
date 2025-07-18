<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';
	import { isDbConnected } from '$lib/api';

	onMount(async () => {
		if (window.location.pathname != '/') {
			const res = await isDbConnected();
			if (!res.isConnected) {
				window.location.href = '/';
			}
		}
	});

	let { children } = $props();
</script>

{@render children()}
