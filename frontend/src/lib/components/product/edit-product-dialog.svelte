<script lang="ts">
	import type { Image, Product } from '$lib/api/types';
	import { Button, Input, Label, Dialog } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';
	import ImageAutocomplete from '../image/image-autocomplete.svelte';
	import ImagePreview from '../image/image-preview.svelte';

	export let product: Product;
	let image: Image | null = null;

	$: if (product.image) {
		$pb
			.collection('images')
			.getOne(product.image)
			.then((img) => {
				image = img;
			});
	}

	export let open = false;
	let loading = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Produkt bearbeiten</Dialog.Title>
		</Dialog.Header>
		<div class="flex flex-col gap-4">
			<ImagePreview
				collection={image?.collectionId}
				record={image?.id}
				filename={image?.file}
				thumbnail
			/>
			<div class="flex w-full flex-col gap-1">
				<Label for="name">Name</Label>
				<Input placeholder="Name" id="name" bind:value={product.name} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label>Bild</Label>
				<ImageAutocomplete bind:value={image} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label for="notes">Notizen</Label>
				<Input placeholder="Notizen" id="notes" bind:value={product.notes} />
			</div>
		</div>
		<Dialog.Footer>
			<Button
				type="submit"
				disabled={loading || !product.name}
				on:click={() => {
					loading = true;
					$pb
						.collection('products')
						.update(product.id, {
							...product,
							image: image?.id ?? null
						})
						.then(() => {
							loading = false;
							open = false;
						})
						.catch(() => {
							loading = false;
							toast.error('Produkt konnte nicht bearbeitet werden.');
						});
				}}>Produkt bearbeiten</Button
			>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
