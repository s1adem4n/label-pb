<script lang="ts">
	import Plus from 'virtual:icons/lucide/plus';

	import { page } from '$app/stores';
	import type { Batch, Product } from '$lib/api/types';
	import BatchesTable from '$lib/components/batch/batches-table.svelte';
	import CreateBatchDialog from '$lib/components/batch/create-batch-dialog.svelte';
	import ProductAutocomplete from '$lib/components/product/product-autocomplete.svelte';
	import { Button } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';

	let product: Product | null = null;
	$: if (product) {
		productId = product.id;
	}
	let productId: string | null = null;
	$: if (productId) {
		$page.url.searchParams.set('productId', productId);
	}

	$: if (productId) {
		$pb
			.collection('batches')
			.getFullList({
				filter: `product = '${productId}'`
			})
			.then((res) => {
				batches = res;
			})
			.catch(() => {
				toast.error('Chargen konnten nicht geladen werden.');
			});
	}

	onMount(() => {
		productId = $page.url.searchParams.get('productId');

		$pb.collection('batches').subscribe('*', (e) => {
			if (e.record.product !== productId) return;
			if (e.action === 'create') {
				batches = [...batches, e.record];
			}
			if (e.action === 'update') {
				const index = batches.findIndex((i) => i.id === e.record.id);
				batches[index] = e.record;
			}
			if (e.action === 'delete') {
				batches = batches.filter((i) => i.id !== e.record.id);
			}
		});

		return () => {
			$pb.collection('batches').unsubscribe('*');
		};
	});

	let batches: Batch[] = [];
	let dialogOpen = false;
</script>

<CreateBatchDialog product={productId ?? ''} bind:open={dialogOpen} />
<div class="flex flex-wrap gap-2 p-2">
	<ProductAutocomplete bind:value={product} />
	<Button
		variant="outline"
		on:click={() => (dialogOpen = true)}
		disabled={!productId}
		class="gap-1"
	>
		<Plus class="h-4" />
		Neue Charge
	</Button>
</div>

<BatchesTable {batches} />
