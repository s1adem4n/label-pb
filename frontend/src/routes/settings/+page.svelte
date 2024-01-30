<script lang="ts">
	import { Button, Label, Input, AlertDialog, Separator } from '$lib/components/ui';
	import { pb } from '$lib/stores/pb';
	import { auth } from '$lib/stores/auth';

	let username = '';
	let oldPassword = '';
	let password = '';
	let passwordConfirm = '';

	let loading = false;
	let dialogOpen = false;
</script>

<AlertDialog.Root bind:open={dialogOpen}>
	<AlertDialog.Content>
		<AlertDialog.Header>
			<AlertDialog.Title>Bist du wirklich sicher?</AlertDialog.Title>
			<AlertDialog.Description>
				Diese Aktion kann nicht rückgängig gemacht werden. Hiermit werden alle deine Daten <bold
					>unwiderruflich</bold
				> gelöscht.
			</AlertDialog.Description>
		</AlertDialog.Header>
		<AlertDialog.Footer>
			<AlertDialog.Cancel>Abbrechen</AlertDialog.Cancel>
			<AlertDialog.Action asChild>
				<Button
					disabled={loading}
					on:click={() => {
						loading = true;
						$pb
							.collection('users')
							.delete($auth.model.id)
							.finally(() => {
								loading = false;
								dialogOpen = false;
								auth.logout();
							});
					}}
					variant="destructive">Konto löschen</Button
				>
			</AlertDialog.Action>
		</AlertDialog.Footer>
	</AlertDialog.Content>
</AlertDialog.Root>

<div class="flex flex-col gap-6 p-4">
	<div class="flex flex-col gap-2">
		<h3 class="text-xl font-bold leading-6">Benutzername ändern</h3>
		<div>
			<Label for="username">Neuer Benutzername</Label>
			<Input placeholder="max-mustermann" bind:value={username} id="username" />
		</div>

		<Button
			disabled={!username || loading}
			on:click={() => {
				loading = true;
				$pb
					.collection('users')
					.update($auth.model.id, { username })
					.then(() => {
						loading = false;
						$auth.model.name = username;
					})
					.catch(() => {
						loading = false;
					});
			}}>Benutzername ändern</Button
		>
	</div>

	<Separator />

	<div class="flex flex-col gap-2">
		<h3 class="text-xl font-bold leading-6">Passwort ändern</h3>

		<div>
			<Label for="oldPassword">Altes Passwort</Label>
			<Input type="password" placeholder="●●●●●" bind:value={oldPassword} id="oldPassword" />
		</div>

		<div>
			<Label for="password">Neues Passwort</Label>
			<Input type="password" placeholder="●●●●●" bind:value={password} id="password" />
		</div>

		<div>
			<Label for="passwordConfirm">Passwort bestätigen</Label>
			<Input
				type="password"
				placeholder="●●●●●"
				bind:value={passwordConfirm}
				id="passwordConfirm"
			/>
		</div>

		<Button
			disabled={!password || password !== passwordConfirm || loading}
			on:click={() => {
				loading = true;
				$pb
					.collection('users')
					.update($auth.model.id, {
						password,
						passwordConfirm,
						oldPassword
					})
					.finally(() => {
						loading = false;
						auth.logout();
					});
			}}>Passwort ändern</Button
		>
	</div>

	<Separator />

	<div class="flex flex-col gap-2">
		<h3 class="text-xl font-bold leading-6">Konto löschen</h3>
		<p class="text-destructive text-sm">Warnung: Dies kann nicht rückgängig gemacht werden!</p>
		<Button variant="destructive" on:click={() => (dialogOpen = true)}>Konto löschen</Button>
	</div>

	<Separator />

	{#if $auth.model.admin}
		<div class="flex flex-col gap-2">
			<h3 class="text-xl font-bold leading-6">Server-Aktionen</h3>
			<Button
				disabled={loading}
				on:click={() => {
					loading = true;
					$pb.send('/update', {}).finally(() => {
						loading = false;
					});
				}}
			>
				Update installieren
			</Button>
			<Button
				disabled={loading}
				on:click={() => {
					loading = true;
					$pb.send('/restart', {}).finally(() => {
						loading = false;
					});
				}}
			>
				Neustarten
			</Button>
		</div>
	{/if}
</div>
