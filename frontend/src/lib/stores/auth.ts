import { get, writable } from 'svelte/store';
import { pb } from './pb';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
type AuthModel = Record<string, any>;

interface Auth {
	model: AuthModel;
	token: string;
	isValid: boolean;
}

const createAuth = () => {
	const { subscribe, set, update } = writable<Auth>({
		model: {},
		token: '',
		isValid: false
	});

	pb.subscribe((pb) => {
		pb.authStore.onChange((token, model) => {
			set({
				model: model || {},
				token,
				isValid: pb.authStore.isValid
			});
		}, true);
	});

	const login = async (usernameOrEmail: string, password: string) => {
		const client = get(pb);
		await client.collection('users').authWithPassword(usernameOrEmail, password);
	};

	const logout = () => {
		const client = get(pb);
		client.authStore.clear();
		set({
			model: {},
			token: '',
			isValid: false
		});
	};

	return {
		subscribe,
		set,
		update,
		login,
		logout
	};
};

export const auth = createAuth();
