<script lang="ts">
	import { pb } from '$lib/stores/pb';
</script>

{#await $pb.collection('products').getFullList() then res}
	{#each res as product}
		<div class="flex flex-col items-center justify-center gap-2 p-4">
			<h1>{product.name}</h1>
			<div>{@html product.notes}</div>
		</div>
	{/each}
{/await}
