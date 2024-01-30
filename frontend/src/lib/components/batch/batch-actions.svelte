<script lang="ts">
	import MoreHorizontal from 'virtual:icons/lucide/more-horizontal';
	import Trash from 'virtual:icons/lucide/trash';
	import Pen from 'virtual:icons/lucide/pen';

	import { pb } from '$lib/stores/pb';
	import { Button, DropdownMenu } from '$lib/components/ui';
	import type { Batch } from '$lib/api/types';
	import EditBatchDialog from './edit-batch-dialog.svelte';

	let open = false;
	let dialogOpen = false;
	export let batch: Batch;
</script>

<EditBatchDialog {batch} bind:open={dialogOpen} />
<DropdownMenu.Root bind:open>
	<DropdownMenu.Trigger asChild let:builder>
		<Button builders={[builder]} variant="ghost" size="sm" aria-label="Aktionen">
			<MoreHorizontal />
		</Button>
	</DropdownMenu.Trigger>
	<DropdownMenu.Content>
		<DropdownMenu.Label>Aktionen</DropdownMenu.Label>
		<DropdownMenu.Separator />
		<DropdownMenu.Item on:click={() => (dialogOpen = true)}>
			<Pen class="mr-2 h-4 w-4" />
			Bearbeiten
		</DropdownMenu.Item>
		<DropdownMenu.Item
			on:click={() => {
				$pb.collection('batches').delete(batch.id);
			}}
			class="text-red-500"
		>
			<Trash class="mr-2 h-4 w-4" />
			Löschen
		</DropdownMenu.Item>
	</DropdownMenu.Content>
</DropdownMenu.Root>
