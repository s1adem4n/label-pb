<script lang="ts">
	import MoreHorizontal from 'virtual:icons/lucide/more-horizontal';
	import Trash from 'virtual:icons/lucide/trash';
	import Pen from 'virtual:icons/lucide/pen';
	import Eye from 'virtual:icons/lucide/eye';
	import Download from 'virtual:icons/lucide/download';

	import { pb } from '$lib/stores/pb';
	import { Button, DropdownMenu } from '$lib/components/ui';
	import type { Image } from '$lib/api/types';
	import EditImageDialog from './edit-image-dialog.svelte';
	import ViewImageDialog from './view-image-dialog.svelte';
	import { downloadUrl, getFileUrl } from '$lib/utils';
	import { toast } from 'svelte-sonner';

	let open = false;
	let editDialogOpen = false;
	let viewDialogOpen = false;
	export let image: Image;
</script>

<ViewImageDialog {image} bind:open={viewDialogOpen} />
<EditImageDialog {image} bind:open={editDialogOpen} />
<DropdownMenu.Root bind:open>
	<DropdownMenu.Trigger asChild let:builder>
		<Button builders={[builder]} variant="outline" size="sm" aria-label="Aktionen">
			<MoreHorizontal />
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content>
		<DropdownMenu.Label>Aktionen</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Item on:click={() => (editDialogOpen = true)}>
			<Pen class="mr-2 h-4 w-4" />
			Bearbeiten
		</DropdownMenu.Item>
		<DropdownMenu.Item on:click={() => (viewDialogOpen = true)}>
			<Eye class="mr-2 h-4 w-4" />
			Anzeigen
		</DropdownMenu.Item>
		<DropdownMenu.Item
			on:click={() => {
				downloadUrl(getFileUrl(image.collectionId, image.id, image.file), image.file);
			}}
		>
			<Download class="mr-2 h-4 w-4" />
			Herunterladen
		</DropdownMenu.Item>
		<DropdownMenu.Item
			on:click={() => {
				$pb
					.collection('images')
					.delete(image.id)
					.catch(() => {
						toast.error('Bild konnte nicht gelöscht werden.');
					});
			}}
			class="text-red-500"
		>
			<Trash class="mr-2 h-4 w-4" />
			Löschen
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
