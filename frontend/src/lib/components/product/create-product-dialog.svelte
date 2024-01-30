<script lang="ts">
	import { Button, Input, Label, Dialog } from '$lib/components/ui';
	import type { Image } from '$lib/api/types';
	import { auth } from '$lib/stores/auth';
	import { pb } from '$lib/stores/pb';
	import ImageAutocomplete from '../image/image-autocomplete.svelte';
	import ImagePreview from '../image/image-preview.svelte';

	let product = {
		name: '',
		notes: '',
		owner: $auth.model.id,
		image: ''
	};

	let image: Image | null = null;
	$: if (image) {
		product.image = image.id;
	} else {
		product.image = '';
	}

	export let open = false;
	let loading = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Produkt erstellen</Dialog.Title>
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
						.create(product)
						.then(() => {
							loading = false;
							open = false;
						});
				}}>Produkt erstellen</Button
			>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
