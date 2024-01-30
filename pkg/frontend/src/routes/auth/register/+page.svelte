<script lang="ts">
	import { Button, Label, Input } from '$lib/components/ui';

	let username = '';
	let email = '';
	let password = '';

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
	<Label for="password">Passwort</Label>
	<Input placeholder="●●●●●" type="password" bind:value={password} />
</div>

<Button
	class="w-full {loading && 'cursor-wait'}"
	disabled={!username || !email || !password || loading}
	on:click={() => {}}>Registrieren</Button
>

<p class="mt-4 text-center text-sm">
	Du hast bereits ein Konto? <a class="underline" href="/auth/login">Einloggen</a>
</p>
