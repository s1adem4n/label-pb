<script lang="ts">
	import type { Batch } from '$lib/api/types';
	import { Table } from '$lib/components/ui';
	import { DateFormatter } from '@internationalized/date';
	import BatchActions from './batch-actions.svelte';

	export let batches: Batch[] = [];

	const df = new DateFormatter('de-DE', {
		dateStyle: 'short'
	});
</script>

<Table.Root>
	<Table.Header>
		<Table.Row>
			<Table.Head>Hergestellt am</Table.Head>
			<Table.Head>Läuft ab am</Table.Head>
			<Table.Head>Anzahl</Table.Head>
			<Table.Head>Notizen</Table.Head>
			<Table.Head class="w-[1%] text-nowrap"></Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#if batches.length === 0}
			<Table.Row>
				<Table.Cell colspan={5}>Keine Chargen vorhanden</Table.Cell>
			</Table.Row>
		{/if}
		{#each batches as batch}
			<Table.Row>
				<Table.Cell>{df.format(new Date(batch.manufactured))}</Table.Cell>
				<Table.Cell>{df.format(new Date(batch.expires))}</Table.Cell>
				<Table.Cell>{batch.quantity}</Table.Cell>
				<Table.Cell>{batch.notes}</Table.Cell>
				<Table.Cell>
					<BatchActions {batch} />
				</Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
