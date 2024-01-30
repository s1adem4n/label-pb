import PocketBase, { type RecordService } from 'pocketbase';

interface Base {
	id: string;
	collectionId: string;
	created: string;
	updated: string;
}

export interface Batch extends Base {
	product: string;
	owner: string;
	quantity: number;
	notes: string;
	expires: string;
	manufactured: string;
}

export interface Product extends Base {
	owner: string;
	name: string;
	notes: string;
	image: string;
}

export interface Image extends Base {
	owner: string;
	name: string;
	file: string;
}

export interface TypedPocketBase extends PocketBase {
	collection(idOrName: string): RecordService;
	collection(idOrName: 'batches'): RecordService<Batch>;
	collection(idOrName: 'products'): RecordService<Product>;
	collection(idOrName: 'images'): RecordService<Image>;
}
