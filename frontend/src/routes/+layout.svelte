<script lang="ts">
	import '../app.css';
	import { auth } from '$lib/stores/auth';
	import Sidebar from '$lib/components/sidebar/sidebar.svelte';
	import { afterNavigate, goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { page } from '$app/stores';
	import { Toaster } from 'svelte-sonner';
	import { settings } from '$lib/stores/settings';

	let currentTitle: string = '';
	let sidebarOpen = window.innerWidth > 640 ? true : false;

	let toasterTheme: 'dark' | 'light' = 'light';
	$: toasterTheme = $settings.dark ? 'dark' : 'light';

	$: if (
		!$auth.isValid &&
		$page.url.pathname !== `${base}/auth/login` &&
		$page.url.pathname !== `${base}/auth/register`
	) {
		goto(`${base}/auth/login`);
	}

	afterNavigate(() => {
		setTimeout(() => {
			document.getElementById('scroller')?.scrollTo(0, 0);
		}, 100);
	});
</script>

<Toaster theme={toasterTheme} />

{#if $auth.isValid}
	<div class="flex w-full overflow-hidden">
		<Sidebar bind:currentTitle bind:open={sidebarOpen} />
		<div class="relative h-[100dvh] w-full overflow-auto" id="scroller">
			<slot />
		</div>
	</div>
{:else}
	<div
		class="relative mx-auto flex h-[100dvh] w-full max-w-md flex-col items-center justify-center gap-2 p-4"
	>
		<slot />
	</div>
{/if}
