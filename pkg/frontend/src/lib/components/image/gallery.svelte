<script lang="ts">
	import type { Image } from '$lib/api/types';
	import ImageActions from './image-actions.svelte';
	import ImagePreview from './image-preview.svelte';

	export let images: Image[] = [];
</script>

<div class="grid w-full grid-cols-1 gap-2 sm:grid-cols-2 lg:grid-cols-3">
	{#each images as image}
		<div class="border-border flex flex-col gap-2 rounded-md border p-2 shadow-sm">
			<ImagePreview
				collection={image.collectionId}
				record={image.id}
				filename={image.file}
				thumbnail
			/>
			<div class="flex items-center">
				<span class="mr-auto">{image.name}</span>
				<ImageActions {image} />
			</div>
		</div>
	{/each}
</div>
