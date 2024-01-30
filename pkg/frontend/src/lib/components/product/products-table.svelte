<script lang="ts">
	import type { Product } from '$lib/api/types';
	import { Table } from '$lib/components/ui';
	import ProductActions from './product-actions.svelte';

	export let products: Product[] = [];
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			<Table.Head>Name</Table.Head>
			<Table.Head>Notizen</Table.Head>
			<Table.Head class="w-[1%] text-nowrap"></Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#if products.length === 0}
			<Table.Row>
				<Table.Cell colspan={4}>Keine Produkte vorhanden</Table.Cell>
			</Table.Row>
		{/if}
		{#each products as product}
			<Table.Row>
				<Table.Cell>{product.name}</Table.Cell>
				<Table.Cell>{@html product.notes}</Table.Cell>
				<Table.Cell>
					<ProductActions {product} />
				</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
