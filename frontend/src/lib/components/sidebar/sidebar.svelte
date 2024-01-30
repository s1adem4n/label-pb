<script lang="ts">
	import Box from 'virtual:icons/lucide/box';
	import LogOut from 'virtual:icons/lucide/log-out';
	import Home from 'virtual:icons/lucide/home';
	import Image from 'virtual:icons/lucide/image';
	import Archive from 'virtual:icons/lucide/archive';
	import X from 'virtual:icons/lucide/x';
	import Menu from 'virtual:icons/lucide/menu';
	import Settings from 'virtual:icons/lucide/settings';

	import { Button } from '$lib/components/ui';
	import { auth } from '$lib/stores/auth';

	import SidebarItem from './sidebar-item.svelte';
	import DarkModeSwitch from '../dark-mode-switch.svelte';
	import { page } from '$app/stores';

	const titles: Record<string, string> = {
		'': 'Startseite',
		products: 'Produkte',
		batches: 'Chargen',
		images: 'Bilder',
		settings: 'Einstellungen'
	};

	const titlesLength = Object.keys(titles).length;

	export let currentTitle: string = '';
	let activeIndex = 0;

	$: activeIndex = Object.keys(titles).indexOf($page.url.pathname.split('/')[2]);
	$: currentTitle = titles[$page.url.pathname.split('/')[2]] ?? '';
	export let open = false;
</script>

<div
	class="border-border flex h-[100dvh] w-full flex-col gap-2 overflow-hidden border-r p-2 transition-[max-width] {open
		? 'max-w-[14rem]'
		: 'max-w-[3.5rem]'}"
>
	<div class="flex items-center justify-between">
		{#if open}
			<span class="whitespace-nowrap pl-2">
				Hallo, {$auth.model.username}!
			</span>
		{/if}
		<Button size="icon" variant="ghost" class="ml-auto" on:click={() => (open = !open)}>
			{#if open}
				<X class="h-6 w-6" />
			{:else}
				<Menu class="h-6 w-6" />
			{/if}
		</Button>
	</div>
	<div class="relative flex flex-col gap-2">
		<div
			class="bg-accent absolute inset-x-0 -z-10 h-10 rounded-md transition-[top]"
			style="top: calc({activeIndex} * 2.5rem + {activeIndex} * 0.5rem)"
		/>
		<SidebarItem path="" title={titles['']} bind:open>
			<Home class="h-6 w-6" />
		</SidebarItem>
		<SidebarItem path="products" title={titles['products']} bind:open>
			<Box class="h-6 w-6" />
		</SidebarItem>
		<SidebarItem path="batches" title={titles['batches']} bind:open>
			<Archive class="h-6 w-6" />
		</SidebarItem>
		<SidebarItem path="images" title={titles['images']} bind:open>
			<Image class="h-6 w-6" />
		</SidebarItem>
		<SidebarItem path="settings" title={titles['settings']} bind:open>
			<Settings class="h-6 w-6" />
		</SidebarItem>
	</div>

	<div class="mt-auto flex flex-wrap items-center justify-between gap-2">
		<DarkModeSwitch />
		<Button size="icon" variant="ghost" on:click={() => auth.logout()}>
			<LogOut class="h-6 w-6" />
			<span class="sr-only">Ausloggen</span>
		</Button>
	</div>
</div>
