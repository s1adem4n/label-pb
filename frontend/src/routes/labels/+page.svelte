<script lang="ts">
	import { onMount } from 'svelte';
	import {
		Rectangle,
		Container,
		Text,
		ImageElement,
		centerElement
	} from '$lib/components/label/elements';
	import FormatSelect, { formats } from '$lib/components/label/format-select.svelte';
	import Element from '$lib/components/label/element.svelte';

	const container = new Container();
	container.push(new Rectangle(0, 0, 100, 100, true));
	container.push(
		new Text(100, 100, {
			text: {
				type: 'text',
				label: 'Text',
				value: 'Hallo Welt'
			},
			font: {
				type: 'text',
				label: 'Schriftart',
				value: 'Arial'
			},
			size: {
				type: 'text',
				label: 'Schriftgröße',
				value: '28px'
			},
			color: {
				type: 'color',
				label: 'Farbe',
				value: '#000000'
			}
		})
	);
	container.push(
		new ImageElement(
			100,
			100,
			{
				src: {
					type: 'text',
					label: 'Bild',
					value: 'https://picsum.photos/200/300'
				}
			},
			100,
			100,
			true
		)
	);
	let canvas: HTMLCanvasElement;
	let ctx: CanvasRenderingContext2D;
	let dragging: boolean = false;
	let mousePos = { x: 0, y: 0 };
	let frame = 0;
	const focused = container.focused;

	$: dragging
		? (document.body.style.cursor = 'grabbing')
		: (document.body.style.cursor = 'default');

	const draw = () => {
		canvas.width = canvas.clientWidth * devicePixelRatio;
		canvas.height = canvas.clientHeight * devicePixelRatio;

		ctx.clearRect(0, 0, canvas.width, canvas.height);
		container.draw(ctx);
		frame = requestAnimationFrame(draw);
	};

	onMount(() => {
		ctx = canvas.getContext('2d')!;
		requestAnimationFrame(draw);

		return () => {
			cancelAnimationFrame(frame);
		};
	});
</script>

<svelte:window
	on:keydown={(e) => {
		if (e.key === '+') {
			container.resize(mousePos.x, mousePos.y, 1);
		}
		if (e.key === '-') {
			container.resize(mousePos.x, mousePos.y, -1);
		}
	}}
	on:resize={() => {}}
/>

<div class="flex h-full flex-col gap-2 p-2">
	<div class="flex items-center gap-2">
		<span>Format:</span>
		<FormatSelect
			on:input={(e) => {
				const format = e.detail;
				canvas.style.width = `${format.width}px`;
				canvas.style.height = `${format.height}px`;
			}}
		/>
	</div>
	<div class="flex h-full">
		<div class="grid-pattern flex h-full w-full items-center justify-center">
			<canvas
				on:click={(e) => {
					container.click(e.offsetX, e.offsetY);
				}}
				on:mousemove={(e) => {
					mousePos = { x: e.offsetX, y: e.offsetY };

					if (dragging) {
						container.drag(mousePos.x, mousePos.y, e.movementX, e.movementY);
					} else {
						container.hover(mousePos.x, mousePos.y);
					}
				}}
				on:mousedown={() => {
					dragging = true;
				}}
				on:mouseup={() => {
					dragging = false;
				}}
				on:mouseleave={() => {
					dragging = false;
				}}
				class="rounded-md border bg-white transition-[width,height]"
				style="width: 400px; height: 300px;"
				bind:this={canvas}
			/>
		</div>
		<div class="flex w-60 gap-2">
			{#if $focused}
				<Element element={$focused} />
			{:else}
				<div class="flex h-full w-full items-center justify-center rounded-md bg-gray-100">
					<span class="text-center text-gray-400">Kein Element ausgewählt</span>
				</div>
			{/if}
		</div>
	</div>
</div>

<style lang="postcss">
	.grid-pattern {
		background-image: radial-gradient(hsl(var(--secondary)) 1px, transparent 1px);
		background-size: 16px 16px;
	}
</style>
