import { BASE_URL } from '$lib/api';
import type { TypedPocketBase } from '$lib/api/types';
import PocketBase from 'pocketbase';
import { writable } from 'svelte/store';

export const pb = writable(new PocketBase(BASE_URL) as TypedPocketBase);
