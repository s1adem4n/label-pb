<script lang="ts">
	import type { Image, Product } from '$lib/api/types';
	import { Dialog } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';
	import ImagePreview from '../image/image-preview.svelte';

	export let open = false;
	export let product: Product;
	let image: Image;

	$: if (product.image && open) {
		$pb
			.collection('images')
			.getOne(product.image)
			.then((img) => {
				image = img;
			})
			.catch((err) => {
				toast.error('Bild konnte nicht geladen werden.');
			});
	}
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Produktdetails</Dialog.Title>
		</Dialog.Header>
		<div class="flex flex-col gap-4">
			<div class="flex w-full flex-col gap-1">
				<span class="font-bold">Name</span>
				<span class="text-sm">{product.name}</span>
			</div>
			<div class="flex w-full flex-col gap-1">
				<span class="font-bold">Notizen</span>
				{#if product.notes}
					<span class="text-sm">{@html product.notes}</span>
				{:else}
					<span class="text-muted-foreground text-sm italic">Keine Notizen</span>
				{/if}
			</div>
			<div class="flex w-full flex-col gap-1">
				<span class="font-bold">Bild</span>
				{#if image}
					<span class="text-sm">{image.name}</span>
					<ImagePreview collection={image.collectionId} record={image.id} filename={image.file} />
				{:else}
					<span class="text-muted-foreground text-sm italic">Kein Bild</span>
				{/if}
			</div>
		</div></Dialog.Content
	>
</Dialog.Root>
