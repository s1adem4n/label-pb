<script lang="ts">
	import { base } from '$app/paths';

	export let title: string;
	export let path: string;
	export let open = false;
</script>

<a
	class="flex h-10 w-full items-center justify-start gap-2 overflow-hidden p-0 px-2 hover:bg-transparent"
	href="{base}/{path}"
>
	<div class="h-6 w-6">
		<slot />
	</div>

	{#if open}
		<span class="text-sm">{title}</span>
	{/if}
</a>
