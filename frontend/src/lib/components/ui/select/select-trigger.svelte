<script lang="ts">
	import ChevronsUpDown from 'virtual:icons/lucide/chevrons-up-down';

	import { Select as SelectPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils';

	type $$Props = SelectPrimitive.TriggerProps;
	type $$Events = SelectPrimitive.TriggerEvents;

	let className: $$Props['class'] = undefined;
	export { className as class };
</script>

<SelectPrimitive.Trigger
	class={cn(
		'border-input ring-offset-background placeholder:text-muted-foreground focus:ring-ring flex h-9 w-full items-center justify-between rounded-md border bg-transparent px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring-1 disabled:cursor-not-allowed disabled:opacity-50',
		className
	)}
	{...$$restProps}
>
	<slot />
	<div>
		<ChevronsUpDown class="h-4 w-4 opacity-50" />
	</div>
</SelectPrimitive.Trigger>
