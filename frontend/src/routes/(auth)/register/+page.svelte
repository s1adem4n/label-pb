<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { Button, Label, Input } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { toast } from 'svelte-sonner';

	let username = '';
	let email = '';
	let password = '';
	let passwordConfirm = '';

	let loading = false;
</script>

<h1 class="text-3xl font-bold">Willkommen!</h1>
<p class="text-muted-foreground text-center text-sm">
	Registriere dich mit deinem Benutzernamen, Passwort und Registrierungscode.
</p>

<div class="w-full">
	<Label for="username">Benutzername</Label>
	<Input placeholder="max-mustermann" id="username" bind:value={username} />
</div>
<div class="w-full">
	<Label for="username">E-Mail</Label>
	<Input placeholder="max@mustermann.de" id="username" bind:value={email} />
</div>
<div class="w-full">
	<Label for="password">Passwort (mind. 8 Zeichen)</Label>
	<Input placeholder="●●●●●" type="password" bind:value={password} />
</div>
<div class="w-full">
	<Label for="password">Passwort bestätigen</Label>
	<Input placeholder="●●●●●" type="password" bind:value={passwordConfirm} />
</div>

<Button
	class="w-full {loading && 'cursor-wait'}"
	disabled={!username ||
		!email ||
		!password ||
		password !== passwordConfirm ||
		password.length < 8 ||
		loading}
	on:click={() => {
		$pb
			.collection('users')
			.create({
				username,
				email,
				password,
				passwordConfirm
			})
			.then(() => {
				goto(`${base}/login`);
			})
			.catch(() => {
				toast.error('Registrierung fehlgeschlagen.');
			})
			.finally(() => {
				loading = false;
			});
	}}>Registrieren</Button
>

<p class="mt-4 text-center text-sm">
	Du hast bereits ein Konto? <a class="underline" href="{base}/login">Einloggen</a>
</p>
