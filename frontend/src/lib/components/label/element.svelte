<script lang="ts">
	import { Checkbox, Input, Label } from '../ui';
	import Separator from '../ui/separator/separator.svelte';
	import ColorInput from './color-input.svelte';
	import type { Element } from './elements';
	import NumberInput from './number-input.svelte';
	import TextInput from './text-input.svelte';

	export let element: Element;
	let keepAspectRatio = false;

	let width = element.width;
	let height = element.height;
	let aspectRatio = element.width / element.height;

	$: {
		if (keepAspectRatio) {
			aspectRatio = element.width / element.height;
			if (width !== aspectRatio * height) {
				height = width / aspectRatio;
			} else if (height !== width / aspectRatio) {
				width = height * aspectRatio;
			}
		}
		element.width = width;
		element.height = height;
	}
</script>

<div class="bg-accent flex flex-col gap-2 rounded-md p-2">
	<span class="font-bold">Eigenschaften</span>

	{#if element.draggable}
		<div class="grid grid-cols-2 gap-2">
			<NumberInput label="X" bind:value={element.x} />
			<NumberInput label="Y" bind:value={element.y} />
			<NumberInput label="Z" bind:value={element.z} />
		</div>
		<Separator />
	{/if}
	{#if element.resize}
		<div class="flex gap-2">
			<Checkbox bind:checked={keepAspectRatio} />
			<span class="text-sm">Verhältnis halten</span>
		</div>
		<div class="flex gap-2">
			<NumberInput label="B" bind:value={width} />
			<NumberInput label="H" bind:value={height} />
		</div>
		<Separator />
	{/if}
	{#if element.attributes}
		<div class="flex flex-col gap-1">
			{#each Object.entries(element.attributes) as [key, value]}
				{#if !value.hidden}
					{#if value.type === 'text'}
						<TextInput label={key.slice(0, 1).toUpperCase()} bind:value={value.value} />
					{:else if value.type === 'color'}
						<ColorInput label={key.slice(0, 1).toUpperCase()} bind:value={value.value} />
					{/if}
				{/if}
			{/each}
		</div>
	{/if}
</div>
