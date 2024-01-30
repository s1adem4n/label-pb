<script lang="ts">
	import { Dialog } from '$lib/components/ui';
	import type { Image } from '$lib/api/types';
	import ImagePreview from './image-preview.svelte';

	export let open = false;
	export let image: Image;
</script>

<Dialog.Root bind:open>
	<Dialog.Content fullWidth>
		<Dialog.Header>
			<Dialog.Title>{image.name}</Dialog.Title>
		</Dialog.Header>
		<ImagePreview collection={image.collectionId} record={image.id} filename={image.file} />
	</Dialog.Content>
</Dialog.Root>
