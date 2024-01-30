<script lang="ts">
	import { auth } from '$lib/stores/auth';
</script>

<div class="flex flex-col gap-2 p-2">
	<h1 class="text-2xl">
		Hallo, {$auth.model.username}!
	</h1>
</div>
