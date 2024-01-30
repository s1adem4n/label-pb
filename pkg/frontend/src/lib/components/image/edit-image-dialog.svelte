<script lang="ts">
	import type { Image } from '$lib/api/types';
	import { Button, Input, Label, Dialog } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { getFileUrl } from '$lib/utils';
	import ImagePreview from './image-preview.svelte';

	export let image: Image;

	let name = image.name;
	let file: File | null = null;

	let src = getFileUrl(image.collectionId, image.id, image.file);
	$: if (file) {
		src = URL.createObjectURL(file);
	}

	export let open = false;
	let loading = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Bild bearbeiten</Dialog.Title>
		</Dialog.Header>
		<div class="flex flex-col gap-4">
			<ImagePreview {src} />
			<div class="flex w-full flex-col gap-1">
				<Label for="name">Name</Label>
				<Input placeholder="Name" id="name" bind:value={name} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label for="image">Bild</Label>
				<Input
					type="file"
					id="image"
					on:input={(e) => {
						if (e.currentTarget.files && e.currentTarget.files.length > 0) {
							file = e.currentTarget.files[0];
						}
					}}
				/>
			</div>
		</div>
		<Dialog.Footer>
			<Button
				type="submit"
				disabled={loading || !name}
				on:click={() => {
					if (!name) return;
					loading = true;

					const formData = new FormData();
					formData.append('name', name);
					if (file) {
						formData.append('file', file);
					}
					formData.append('owner', image.owner);

					$pb
						.collection('images')
						.update(image.id, formData)
						.then(() => {
							loading = false;
							open = false;
						});
				}}
			>
				Bild hochladen
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
