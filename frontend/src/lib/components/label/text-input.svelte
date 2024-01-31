<script lang="ts">
	export let label: string;
	export let value: string = '';
</script>

<div class="relative flex items-center">
	<span class="absolute left-2">
		{label}
	</span>
	<input
		class="focus-visible:border-primary border-input bg-background h-8 w-full rounded-md border p-1 pl-6 text-sm focus-visible:outline-none"
		type="text"
		bind:value
	/>
</div>
