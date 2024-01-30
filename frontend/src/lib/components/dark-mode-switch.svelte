<script lang="ts">
	import { settings } from '$lib/stores/settings';
	import Sun from 'virtual:icons/lucide/sun';
	import Moon from 'virtual:icons/lucide/moon';

	import { Button } from '$lib/components/ui';
</script>

<Button
	size="icon"
	variant="ghost"
	on:click={() => {
		$settings.dark = !$settings.dark;
	}}
>
	{#if $settings.dark}
		<Sun class="h-6 w-6" />
	{:else}
		<Moon class="h-6 w-6" />
	{/if}
</Button>
