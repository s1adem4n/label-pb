<script lang="ts">
	import CalendarIcon from 'virtual:icons/lucide/calendar';

	import {
		type DateValue,
		DateFormatter,
		getLocalTimeZone,
		fromDate
	} from '@internationalized/date';
	import { cn } from '$lib/utils';
	import { Button, Calendar, Popover } from '$lib/components/ui';

	const df = new DateFormatter('de-DE', {
		dateStyle: 'long'
	});
	export let date = new Date();
	$: date = value ? value.toDate(getLocalTimeZone()) : new Date();

	let value: DateValue = fromDate(date, getLocalTimeZone());
</script>

<Popover.Root>
	<Popover.Trigger asChild let:builder>
		<Button
			variant="outline"
			class={cn('justify-start text-left font-normal', !value && 'text-muted-foreground')}
			builders={[builder]}
		>
			<CalendarIcon class="mr-2 h-4 w-4" />
			{df.format(value.toDate(getLocalTimeZone()))}
		</Button>
	</Popover.Trigger>
	<Popover.Content class="w-auto p-0">
		<Calendar bind:value initialFocus />
	</Popover.Content>
</Popover.Root>
