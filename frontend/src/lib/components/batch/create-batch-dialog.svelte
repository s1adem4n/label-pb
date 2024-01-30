<script lang="ts">
	import { Button, Label, Dialog, Slider, DatePicker, Input } from '$lib/components/ui';
	import { auth } from '$lib/stores/auth';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';

	export let product: string;
	let batch = {
		product,
		quantity: 0,
		notes: '',
		manufactured: new Date(),
		expires: new Date(),
		owner: $auth.model.id
	};
	$: {
		batch.product = product;
	}

	let quantity: number[] = [2];
	$: batch.quantity = quantity[0];

	export let open = false;
	let loading = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Charge erstellen</Dialog.Title>
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
				<DatePicker bind:date={batch.manufactured} />
			</div>
			<div class="flex w-full flex-col gap-1">
				<Label>Läuft ab am</Label>
				<DatePicker bind:date={batch.expires} />
			</div>
		</div>
		<Dialog.Footer>
			<Button
				disabled={loading}
				on:click={() => {
					loading = true;
					$pb
						.collection('batches')
						.create(batch)
						.then(() => {
							loading = false;
							open = false;
						})
						.catch(() => {
							loading = false;
							toast.error('Charge konnte nicht erstellt werden.');
						});
				}}>Charge erstellen</Button
			>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
