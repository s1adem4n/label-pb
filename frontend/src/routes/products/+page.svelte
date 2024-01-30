<script lang="ts">
	import Plus from 'virtual:icons/lucide/plus';

	import type { Product } from '$lib/api/types';
	import CreateProductDialog from '$lib/components/product/create-product-dialog.svelte';
	import ProductsTable from '$lib/components/product/products-table.svelte';
	import { Button } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { onMount } from 'svelte';

	let products: Product[] = [];
	let dialogOpen = false;

	const fetchProducts = () => {
		$pb
			.collection('products')
			.getFullList()
			.then((res) => {
				products = res;
			});
	};

	onMount(() => {
		fetchProducts();

		$pb.collection('products').subscribe('*', (e) => {
			console.log(e);
			if (e.action === 'create') {
				products = [...products, e.record];
			}
			if (e.action === 'update') {
				const index = products.findIndex((i) => i.id === e.record.id);
				products[index] = e.record;
			}
			if (e.action === 'delete') {
				products = products.filter((i) => i.id !== e.record.id);
			}
		});

		return () => {
			$pb.collection('products').unsubscribe('*');
		};
	});
</script>

<CreateProductDialog
	bind:open={dialogOpen}
	on:update={() => {
		fetchProducts();
	}}
/>
<div class="gap-2 p-2">
	<Button variant="outline" class="gap-1" on:click={() => (dialogOpen = true)}>
		<Plus class="h-4" />
		Neues Produkt
	</Button>
</div>
<ProductsTable bind:products />
