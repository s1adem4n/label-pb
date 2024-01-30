<script lang="ts">
	import Check from 'virtual:icons/lucide/check';
	import ChevronsUpDown from 'virtual:icons/lucide/chevrons-up-down';

	import { Command, Button, Popover } from '$lib/components/ui';
	import { cn } from '$lib/utils';
	import { onMount, tick } from 'svelte';
	import type { Image } from '$lib/api/types';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';

	let images: Image[] = [];
	let open = false;
	export let value: Image | null = null;
	$: selectedValue = images.find((p) => p.id === value?.id)?.name ?? 'Wähle ein Bild aus...';

	onMount(() => {
		$pb
			.collection('images')
			.getFullList()
			.then((res) => {
				images = res;
			})
			.catch(() => {
				toast.error('Bilder konnten nicht geladen werden.');
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
			<Command.Input placeholder="Bild suchen..." />
			<Command.Empty>Kein Bild gefunden.</Command.Empty>
			<Command.Group>
				{#each images as image}
					<Command.Item
						value={image.id.toString()}
						onSelect={(currentValue) => {
							value = images.find((p) => p.id === currentValue) ?? null;
							closeAndFocusTrigger(ids.trigger);
						}}
					>
						<Check class={cn('mr-2 h-4 w-4', value?.id !== image.id && 'text-transparent')} />
						{image.name}
					</Command.Item>
				{/each}
				<Command.Item
					onSelect={() => {
						value = null;
						closeAndFocusTrigger(ids.trigger);
					}}
				>
					<Check class={cn('mr-2 h-4 w-4', value !== null && 'text-transparent')} />
					Kein Bild
				</Command.Item>
			</Command.Group>
		</Command.Root>
	</Popover.Content>
</Popover.Root>
