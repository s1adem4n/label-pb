<script lang="ts">
	import { Button, Input, Label, Dialog } from '$lib/components/ui';
	import { auth } from '$lib/stores/auth';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';
	import ImagePreview from './image-preview.svelte';

	let image: File | null = null;
	let name = '';

	let src = '';
	$: if (image) {
		src = URL.createObjectURL(image);
	}

	export let open = false;
	let loading = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Bild hochladen</Dialog.Title>
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
							image = e.currentTarget.files[0];
						}
					}}
				/>
			</div>
		</div>
		<Dialog.Footer>
			<Button
				type="submit"
				disabled={loading || !name || !image}
				on:click={() => {
					if (!image || !name) return;
					loading = true;
					const formData = new FormData();
					formData.append('name', name);
					formData.append('file', image);
					formData.append('owner', $auth.model.id);
					$pb
						.collection('images')
						.create(formData)
						.then(() => {
							loading = false;
							open = false;
						})
						.catch(() => {
							loading = false;
							toast.error('Bild konnte nicht hochgeladen werden.');
						});
				}}
			>
				Bild hochladen
			</Button>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
