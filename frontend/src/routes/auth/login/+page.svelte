<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { auth } from '$lib/stores/auth';

	let usernameOrEmail = '';
	let password = '';

	let loading = false;
</script>

<h1 class="text-3xl font-bold">Willkommen!</h1>
<p class="text-muted-foreground text-center text-sm">
	Bitte logge dich mit deinem Benutzernamen und Passwort ein.
</p>

<div class="w-full">
	<Label for="username">Benutzername / E-Mail</Label>
	<Input placeholder="max-mustermann" id="username" bind:value={usernameOrEmail} />
</div>
<div class="w-full">
	<Label for="password">Passwort</Label>
	<Input placeholder="●●●●●" type="password" bind:value={password} />
</div>

<Button
	class="w-full {loading && 'cursor-wait'}"
	disabled={!usernameOrEmail || !password || loading}
	on:click={() => {
		loading = true;
		auth.login(usernameOrEmail, password).then(() => {
			loading = false;
			goto(`${base}/`);
		});
	}}>Einloggen</Button
>

<p class="mt-4 text-center text-sm">
	Du hast noch kein Konto? <a class="underline" href="{base}/auth/register">Registrieren</a>
</p>
