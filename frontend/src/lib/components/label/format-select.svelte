<script lang="ts" context="module">
	interface Format {
		label: string;
		width: number;
		height: number;
	}
	const baseRes = 100;
	export const formats: Format[] = [
		{
			label: '4x3',
			width: 4 * baseRes,
			height: 3 * baseRes
		},
		{
			label: '3x4',
			width: 3 * baseRes,
			height: 4 * baseRes
		},
		{
			label: '3x3',
			width: 3 * baseRes,
			height: 3 * baseRes
		}
	];
</script>

<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import { createEventDispatcher } from 'svelte';
	const eventDispatcher = createEventDispatcher<{ input: Format }>();
</script>

<Select.Root
	onSelectedChange={(e) => {
		if (!e) return;
		const format = formats.find((f) => f.label === e.value);
		if (!format) return;
		eventDispatcher('input', format);
	}}
>
	<Select.Trigger class="w-[180px]">
		<Select.Value placeholder="Wähle ein Format" />
	</Select.Trigger>
	<Select.Content>
		<Select.Group>
			<Select.Label>Formate</Select.Label>
			{#each formats as format}
				<Select.Item value={format.label} label={`${format.width} x ${format.height}`}
					>{format.width} x {format.height}</Select.Item
				>
			{/each}
		</Select.Group>
	</Select.Content>
	<Select.Input name="favoriteFruit" />
</Select.Root>
