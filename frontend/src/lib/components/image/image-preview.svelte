<script lang="ts">
	import Image from 'virtual:icons/lucide/image';

	import { getFileUrl } from '$lib/utils';

	export let collection: string = '';
	export let record: string = '';
	export let filename: string = '';
	export let src = '';
	export let thumbnail: boolean = false;

	$: if (collection && record && filename) {
		src = getFileUrl(collection, record, filename);
	} else {
		src = '';
	}
</script>

<div class="relative aspect-video h-full w-full overflow-hidden rounded-md">
	<img
		src="{src}{thumbnail ? '?thumb=500x500f' : ''}"
		class="bg-secondary text-secondary h-full w-full object-contain"
		alt="Bild"
	/>
	{#if !src}
		<Image
			class="absolute left-1/2 top-1/2 h-12 w-12 -translate-x-1/2 -translate-y-1/2 transform opacity-70"
		/>
	{/if}
</div>
