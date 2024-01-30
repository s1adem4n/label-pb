import { writable } from 'svelte/store';

interface SettingsStore {
	dark: boolean;
}

const createSettingsStore = () => {
	const { subscribe, set, update } = writable<SettingsStore>({ dark: false });

	const settingsData = localStorage.getItem('settings');
	if (settingsData) {
		set(JSON.parse(settingsData));
	}

	subscribe((state) => {
		localStorage.setItem('settings', JSON.stringify(state));

		if (state.dark) {
			document.documentElement.classList.add('dark');
		} else {
			document.documentElement.classList.remove('dark');
		}
	});

	return {
		subscribe,
		set,
		update
	};
};

export const settings = createSettingsStore();
