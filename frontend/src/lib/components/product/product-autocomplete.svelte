<script lang="ts">
	import Check from 'virtual:icons/lucide/check';
	import ChevronsUpDown from 'virtual:icons/lucide/chevrons-up-down';

	import { Command, Button, Popover } from '$lib/components/ui';
	import { cn } from '$lib/utils';
	import { onMount, tick } from 'svelte';
	import type { Product } from '$lib/api/types';
	import { pb } from '$lib/stores/pb';

	let products: Product[] = [];
	let open = false;
	export let value: Product | null = null;
	$: selectedValue = products.find((p) => p.id === value?.id)?.name ?? 'Wähle ein Produkt aus...';

	onMount(() => {
		$pb
			.collection('products')
			.getFullList()
			.then((res) => {
				products = res;
			});
	});

	function closeAndFocusTrigger(triggerId: string) {
		open = false;
		tick().then(() => {
			document.getElementById(triggerId)?.focus();
		});
	}
</script>

<Popover.Root bind:open let:ids>
	<Popover.Trigger asChild let:builder>
		<Button
			builders={[builder]}
			variant="outline"
			role="combobox"
			aria-expanded={open}
			class="min-w-44 justify-between"
		>
			{selectedValue}
			<ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
		</Button>
	</Popover.Trigger>
	<Popover.Content sameWidth class="p-0">
		<Command.Root>
			<Command.Input placeholder="Produkt suchen..." />
			<Command.Empty>Kein Produkt gefunden.</Command.Empty>
			<Command.Group>
				{#each products as product}
					<Command.Item
						value={product.id.toString()}
						onSelect={(currentValue) => {
							value = products.find((p) => p.id === currentValue) ?? null;
							closeAndFocusTrigger(ids.trigger);
						}}
					>
						<Check class={cn('mr-2 h-4 w-4', value?.id !== product.id && 'text-transparent')} />
						{product.name}
					</Command.Item>
				{/each}
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
