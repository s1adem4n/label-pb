<script lang="ts">
	import type { Batch } from '$lib/api/types';
	import { Button, Label, Dialog, Slider, DatePicker, Input } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';

	export let batch: Batch;
	export let open = false;
	let loading = false;

	let expires = new Date(batch.expires);
	let manufactured = new Date(batch.manufactured);
	$: {
		expires = new Date(batch.expires);
		manufactured = new Date(batch.manufactured);
	}

	let quantity: number[] = [batch.quantity];
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Charge bearbeiten</Dialog.Title>
		</Dialog.Header>
		<div class="flex flex-col gap-4">
			<div class="flex w-full flex-col gap-1">
				<Label for="quantity">Anzahl</Label>
				<div class="flex items-center gap-2">
					<span class="flex w-4 justify-center text-sm tabular-nums">{batch.quantity}</span>
					<Slider bind:value={quantity} id="quantity" min={1} max={99} step={1} />
				</div>
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label for="notes">Notizen</Label>
				<Input placeholder="Notizen" id="notes" bind:value={batch.notes} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label>Hergestellt am</Label>
				<DatePicker bind:date={manufactured} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label>Läuft ab am</Label>
				<DatePicker bind:date={expires} />
			</div>
		</div>
		<Dialog.Footer>
			<Button
				disabled={loading}
				on:click={() => {
					loading = true;
					$pb
						.collection('batches')
						.update(batch.id, {
							...batch,
							quantity: quantity[0],
							manufactured,
							expires
						})
						.then(() => {
							loading = false;
							open = false;
						})
						.catch(() => {
							loading = false;
							toast.error('Charge konnte nicht bearbeitet werden.');
						});
				}}>Charge bearbeiten</Button
			>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
