<script lang="ts">
	import Upload from 'virtual:icons/lucide/upload';

	import type { Image } from '$lib/api/types';
	import Gallery from '$lib/components/image/gallery.svelte';
	import { pb } from '$lib/stores/pb';
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui';
	import CreateImageDialog from '$lib/components/image/create-image-dialog.svelte';

	let images: Image[] = [];
	let dialogOpen = false;

	const fetchImages = () => {
		$pb
			.collection('images')
			.getFullList()
			.then((res) => {
				images = res;
			});
	};

	onMount(() => {
		fetchImages();

		$pb.collection('images').subscribe('*', (e) => {
			if (e.action === 'create') {
				images = [...images, e.record];
			}
			if (e.action === 'update') {
				const index = images.findIndex((i) => i.id === e.record.id);
				images[index] = e.record;
			}
			if (e.action === 'delete') {
				images = images.filter((i) => i.id !== e.record.id);
			}
		});

		return () => {
			$pb.collection('images').unsubscribe('*');
		};
	});
</script>

<CreateImageDialog bind:open={dialogOpen} />
<div class="bg-background sticky top-0 z-10 flex flex-wrap gap-2 p-2">
	<Button variant="outline" class="gap-1" on:click={() => (dialogOpen = true)}>
		<Upload class="h-4" />
		Bild hochladen
	</Button>
</div>
<div class="flex w-full px-2 pb-2">
	{#if images.length === 0}
		<span class="text-muted-foreground italic"> Keine Bilder vorhanden. </span>
	{:else}
		<Gallery {images} />
	{/if}
</div>
