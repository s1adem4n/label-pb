<script lang="ts">
	import MoreHorizontal from 'virtual:icons/lucide/more-horizontal';
	import Trash from 'virtual:icons/lucide/trash';
	import Pen from 'virtual:icons/lucide/pen';
	import Barcode from 'virtual:icons/lucide/barcode';
	import Archive from 'virtual:icons/lucide/archive';
	import Info from 'virtual:icons/lucide/info';

	import { pb } from '$lib/stores/pb';
	import { Button, DropdownMenu } from '$lib/components/ui';
	import type { Product } from '$lib/api/types';
	import EditProductDialog from './edit-product-dialog.svelte';
	import ViewProductDialog from './view-product-dialog.svelte';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { toast } from 'svelte-sonner';

	let open = false;
	let editDialogOpen = false;
	let detailsDialogOpen = false;
	export let product: Product;
</script>

<ViewProductDialog {product} bind:open={detailsDialogOpen} />
<EditProductDialog {product} bind:open={editDialogOpen} />
<DropdownMenu.Root bind:open>
	<DropdownMenu.Trigger asChild let:builder>
		<Button builders={[builder]} variant="ghost" size="sm" aria-label="Aktionen">
			<MoreHorizontal />
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content>
		<DropdownMenu.Label>Aktionen</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Item on:click={() => (detailsDialogOpen = true)}>
			<Info class="mr-2 h-4 w-4" />
			Details
		</DropdownMenu.Item>
		<DropdownMenu.Item on:click={() => (editDialogOpen = true)}>
			<Pen class="mr-2 h-4 w-4" />
			Bearbeiten
		</DropdownMenu.Item>
		<DropdownMenu.Item
			on:click={() => {
				goto(`${base}/batches?productId=${product.id}`);
			}}
		>
			<Archive class="mr-2 h-4 w-4" />
			Chargen anzeigen
		</DropdownMenu.Item>
		<DropdownMenu.Item on:click={() => {}}>
			<Barcode class="mr-2 h-4 w-4" />
			Barcode herunterladen
		</DropdownMenu.Item>
		<DropdownMenu.Item
			on:click={() => {
				$pb
					.collection('products')
					.delete(product.id)
					.catch(() => {
						toast.error('Produkt konnte nicht gelöscht werden.');
					});
			}}
			class="text-red-500"
		>
			<Trash class="mr-2 h-4 w-4" />
			Löschen
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
